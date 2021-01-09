package infrastructures

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"gopkg.in/gorp.v1"
)

// Dababase DB接続情報を保持
type Dababase struct {
	*gorp.DbMap
}

// Transaction トランサクション情報
type Transaction struct {
	*gorp.Transaction
}

// WithDatabase 引数の処理を前後にデータベースコネションの確立と開放を行って実行する
func WithDatabase(
	f func(Dababase) error,
) error {
	db := newDbMap()
	defer db.Db.Close()

	return f(Dababase{db})
}

// WithDatabaseAndTransaction 引数の処理を前後にデータベースコネションの確立と開放、トランザクションの開始と完了を行って実行する
func WithDatabaseAndTransaction(
	f func(Dababase, Transaction) error,
) error {
	db := newDbMap()
	defer db.Db.Close()

	tx, err := db.Begin()
	if err != nil {
		return errors.WithMessagef(err, "トランザクション開始失敗")
	}

	if err := f(Dababase{db}, Transaction{tx}); err != nil {
		e := tx.Rollback()
		if err != nil {
			log.Printf("データベースロールバック失敗 %s %s", err, e)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.WithMessagef(err, "トランザクションコミット失敗")
	}
	return err
}

// newDbMap DB接続を開始
func newDbMap() *gorp.DbMap {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Panicf("SQL接続失敗 %s", errors.WithStack(err))
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
}
