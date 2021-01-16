package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"text/template"
	"time"

	"github.com/TechDepa/c_tool/domain/model"
	"github.com/TechDepa/c_tool/infrastructures"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Setenv("C_TOOL_ENVIRONMENT", "test")

	router := setupRouter()

	// テスト用に現在時刻を固定 2021-11-02 09:57:19
	model.OverrapNowTime(2021, time.Month(11), 2, 19, 57, 19, 100)

	// // goldenを構成
	// g := goldie.New(
	// 	t,
	// 	goldie.WithFixtureDir("testdata/golden_fixtures"),
	// 	goldie.WithNameSuffix(".golden.json"),
	// 	goldie.WithTestNameForDir(true),
	// 	goldie.WithSubTestNameForDir(true),
	// )

	// t.Run("GET adminusers returns 1 user when there is one user", func(t *testing.T) {
	// 	// フィクスチャーデータ
	// 	infrastructures.WithDatabaseAndTransaction(
	// 		func(db infrastructures.Dababase, tx infrastructures.Transaction) error {
	// 			r := gateways.NewAdminUsersRepository(db, tx)
	// 			r.Create(model.AdminUser{BaseUser: model.BaseUser{
	// 				Email: "katiesanchez@henry.info",
	// 				Name:  "宇野 太郎",
	// 			}})
	// 			return nil
	// 		},
	// 	)
	// 	t.Cleanup(func() { truncateTables("admin_users", "base_users") })

	// 	// 実行
	// 	req, _ := http.NewRequest(http.MethodGet, "/v1/admin/users", nil)
	// 	router.ServeHTTP(w, req)

	// 	// チェック
	// 	assert.Equal(t, 200, w.Code)
	// 	g.Assert(t, "get_users", w.Body.Bytes())
	// })

	t.Run("POST /v1/admin/users", func(t *testing.T) {
		cases := []struct {
			TestName        string
			Email           string
			Name            string
			Password        string
			PasswordConfirm string
			ExpectedStatus  int
			ExpectedUsers   int
		}{
			{"正常ケースで200応答", "jima@hotmail.com", "廣川 舞", "a1a2a3a4", "a1a2a3a4", 200, 1},
			{"パスワード不一致で400応答", "jima@hotmail.com", "廣川 舞", "a1a2a3a4", "a1a2a3a4a5", 400, 0},
			{"パスワード長0文字で400応答", "jima@hotmail.com", "廣川 舞", "", "", 400, 0},
			{"パスワード長5文字で400応答", "jima@hotmail.com", "廣川 舞", "a1a2a", "a1a2a", 400, 0},
			{"パスワード長6文字で200応答", "jima@hotmail.com", "廣川 舞", "a1a2a3", "a1a2a3", 200, 1},
			{"パスワード長24文字で200応答", "jima@hotmail.com", "廣川 舞", "a1a2a3a1a2a3a1a2a3a1a2a3", "a1a2a3a1a2a3a1a2a3a1a2a3", 200, 1},
			{"パスワード長25文字で200応答", "jima@hotmail.com", "廣川 舞", "a1a2a3a1a2a3a1a2a3a1a2a3a", "a1a2a3a1a2a3a1a2a3a1a2a3a", 400, 0},
		}
		for _, c := range cases {
			t.Run(c.TestName, func(t *testing.T) {
				// t.Cleanup(func() { truncateTables("base_users", "admin_users") })

				// 前提条件
				tmpl := template.Must(template.ParseFiles("testdata/post_users.json.template"))
				b := &bytes.Buffer{}
				err := tmpl.Execute(b, struct {
					Email           string
					Name            string
					Password        string
					PasswordConfirm string
				}{c.Email, c.Name, c.Password, c.PasswordConfirm})
				if err != nil {
					log.Fatal(err)
				}

				w := httptest.NewRecorder()

				// 実行
				req, _ := http.NewRequest(http.MethodPost, "/v1/admin/users", b)
				router.ServeHTTP(w, req)

				// チェック
				assert.Equal(t, c.ExpectedStatus, w.Code)

				// db := infrastructures.NewDatabasePointerInstance()
				// db.BeginConnection()

				// _, err = gateways.NewAdminUsersRepository(db).FindAll()
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// assert.Equal(t, c.ExpectedUsers, len(u))
			})
		}
	})

	// t.Run("POST /login", func(t *testing.T) {
	// 	t.Cleanup(func() { truncateTables("base_users", "admin_users") })

	// 	// 前提条件
	// 	b, err := os.Open("testdata/post_login.json")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer b.Close()

	// 	j := strings.NewReader(`{"password":"admin","username":"admin"}`)

	// 	// 実行
	// 	req, _ := http.NewRequest(http.MethodPost, "/login", j)
	// 	router.ServeHTTP(w, req)

	// 	// チェック
	// 	assert.Equal(t, 200, w.Code)
	// 	fmt.Println(w.Body.String())

	// })
}

// truncateTables dbMap.TruncateTablesが外部キー制約非対応のため作成
func truncateTables(tableNames ...string) {
	dbMap := infrastructures.NewDbMap()
	defer dbMap.Db.Close()

	for _, tn := range tableNames {
		_, err := dbMap.Exec("truncate table " + tn + " restart identity cascade")
		if err != nil {
			log.Fatalf("テーブル消去失敗: %s", err.Error())
		}
	}
}
