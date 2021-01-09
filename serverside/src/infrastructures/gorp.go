package infrastructures

import (
	"database/sql"
	"log"

	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/domain/model"
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
	db := NewDbMap()
	defer db.Db.Close()

	return f(Dababase{db})
}

// WithDatabaseAndTransaction 引数の処理を前後にデータベースコネションの確立と開放、トランザクションの開始と完了を行って実行する
func WithDatabaseAndTransaction(
	f func(Dababase, Transaction) error,
) error {
	db := NewDbMap()
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

// NewDbMap DB接続を開始
func NewDbMap() *gorp.DbMap {
	db, err := sql.Open("postgres", DBDataSource())
	if err != nil {
		log.Panicf("SQL接続失敗 %s", errors.WithStack(err))
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(model.BaseUser{}, "base_users").SetKeys(true, "Id")
	dbmap.AddTableWithName(gateways.AdminUserRecord{}, "admin_users").SetKeys(true, "Id")
	dbmap.AddTableWithName(model.AdminUser{}, "admin_users").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Companie{}, "companies").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Office{}, "offices").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.PartnerCompany{}, "partner_companies").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Personnel{}, "personnel").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Term{}, "terms").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ExceptTerm{}, "except_terms").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ImageList{}, "image_lists").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Image{}, "images").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ConstructionSiteSubscription{}, "construction_site_subscriptions").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ConstructionSite{}, "construction_sites").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.UserConstructionSite{}, "user_construction_sites").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ConstructionSiteSubscription{}, "construction_site_subscriptions").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.ConstructionSiteSubscription{}, "construction_site_subscriptions").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.Construction{}, "constructions").SetKeys(true, "Id")
	// dbmap.AddTableWithName(model.SubConstruction{}, "sub_constructions").SetKeys(true, "Id")

	return dbmap
}
