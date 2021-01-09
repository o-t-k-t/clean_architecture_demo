package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TechDepa/c_tool/infrastructures"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Setenv("C_TOOL_ENVIRONMENT", "test")

	router := setupRouter()
	w := httptest.NewRecorder()

	t.Run("GET /users", func(t *testing.T) {
		// 実行
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		router.ServeHTTP(w, req)

		// チェック
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "[]", w.Body.String())
	})

	t.Run("POST /users", func(t *testing.T) {
		t.Cleanup(func() { truncateTables("base_users", "admin_users") })

		// 前提条件
		b, err := os.Open("testdata/post_users.json")
		if err != nil {
			log.Fatal(err)
		}
		defer b.Close()

		// 実行
		req, _ := http.NewRequest(http.MethodPost, "/users", b)
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
		_, err := dbMap.Exec("truncate table " + tn + " cascade")
		if err != nil {
			log.Fatalf("テーブル消去失敗: %s", err.Error())
		}
	}
}
