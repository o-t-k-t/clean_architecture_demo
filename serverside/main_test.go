package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/TechDepa/c_tool/infrastructures"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Setenv("C_TOOL_ENVIRONMENT", "test")

	router := setupRouter()
	w := httptest.NewRecorder()

	// テスト用に現在時刻を固定 2021-11-02 09:57:19
	model.OverrapNowTime(2021, time.Month(11), 2, 19, 57, 19, 100)

	// goldenを構成
	g := goldie.New(
		t,
		goldie.WithFixtureDir("testdata/golden_fixtures"),
		goldie.WithNameSuffix(".golden.json"),
		goldie.WithTestNameForDir(true),
		goldie.WithSubTestNameForDir(true),
	)

	t.Run("GET admin users returns 1 user when there is one user", func(t *testing.T) {
		// フィクスチャーデータ
		infrastructures.WithDatabaseAndTransaction(
			func(db infrastructures.Dababase, tx infrastructures.Transaction) error {
				r := gateways.NewAdminUsersRepository(db, tx)
				r.Create(model.AdminUser{BaseUser: model.BaseUser{
					Email: "katiesanchez@henry.info",
					Name:  "宇野 太郎",
				}})
				return nil
			},
		)
		t.Cleanup(func() { truncateTables("admin_users", "base_users") })

		// 実行
		req, _ := http.NewRequest(http.MethodGet, "/v1/admin/users", nil)
		router.ServeHTTP(w, req)

		// チェック
		assert.Equal(t, 200, w.Code)
		g.Assert(t, "get_users", w.Body.Bytes())
	})

	t.Run("POST /v1/admin/users", func(t *testing.T) {
		t.Cleanup(func() { truncateTables("base_users", "admin_users") })

		// 前提条件
		b, err := os.Open("testdata/post_users.json")
		if err != nil {
			log.Fatal(err)
		}
		defer b.Close()

		// 実行
		req, _ := http.NewRequest(http.MethodPost, "/v1/admin/users", b)
		router.ServeHTTP(w, req)

		// チェック
		assert.Equal(t, 200, w.Code)
	})
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
