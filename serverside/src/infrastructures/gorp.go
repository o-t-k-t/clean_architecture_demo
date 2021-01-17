package infrastructures

import (
	"database/sql"
	"log"

	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/domain/model"
	_ "github.com/lib/pq"

	"github.com/pkg/errors"
	"gopkg.in/gorp.v1"
)

// Database DB接続情報を保持
type Database struct {
	dbMap       *gorp.DbMap
	transaction *gorp.Transaction
}

var database = Database{}

func NewDatabasePointer() *Database {
	return &database
}

// Transaction トランサクション情報
type Transaction struct {
	*gorp.Transaction
}

// BeginConnection
func (db Database) BeginConnection() {
	db.dbMap = NewDbMap()
}

// BeginConnectionAndTransaction 引数の処理を前後にデータベースコネションの確立と開放、トランザクションの開始と完了を行って実行する
func (db *Database) BeginConnectionAndTransaction() error {
	db.dbMap = NewDbMap()

	tx, err := db.dbMap.Begin()
	if err != nil {
		return errors.WithMessagef(err, "トランザクション開始失敗")
	}
	db.transaction = tx

	return nil
}

// Select
func (db Database) Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return db.dbMap.Select(i, query, args...)
}

// Insert
func (db Database) Insert(list ...interface{}) error {
	if tx := db.transaction; tx != nil {
		return tx.Insert(list...)
	}

	return db.dbMap.Insert(list...)
}

// CloseAndCommitOrRollback
func (db *Database) CommitAndClose() {
	db.CommitOrRollbackAndClose(true)
}

// CommitAndRollback
func (db *Database) CommitAndRollback() {
	db.CommitOrRollbackAndClose(false)
}

// CloseAndCommitOrRollback
func (db *Database) CommitOrRollbackAndClose(commit bool) {
	if tx := db.transaction; tx != nil {
		if commit {
			if err := tx.Commit(); err != nil {
				log.Fatalf("データベースコミット失敗 %s", err.Error())
			}
		} else {
			if err := tx.Rollback(); err != nil {
				log.Fatalf("データベースロールバック失敗 %s", err.Error())
			}
		}
	}

	if db.dbMap != nil {
		db.dbMap.Db.Close()
	}
}

// Close
func (db *Database) Close() {
	if db.dbMap != nil {
		db.dbMap.Db.Close()
	}
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
