## 使用技術
- アプリケーション
  - プログラミング言語: Go 1.11
  - Webフレームワーク: [Gin](https://github.com/gin-gonic/gin)
  - ORマッパー: [gorp](https://github.com/go-gorp/gorp)
  - データベース: PostgreSQL
  - マイグレーション: [sql-migrate](https://github.com/rubenv/sql-migrate)

go get github.com/rubenv/sql-migrate/...


### ローカルサーバー起動

```
export GO111MODULE=on
docker-compose up
dev_appserver.py app.yml
```

### インテグレーションテスト

```
docker-compose -f docker-compose.yml -f docker-compose-test.yml up -d
sql-migrate up -env=test
go test
```

#### テストの追加

#### リクエストの作成

既存テスト(`serverside/main_test.go`)を参考に作成する。
#### レスポンス期待値の作成
[Golden](https://github.com/sebdah/goldie)を使用する。
まずはいったん既存テストのようにgoldieのアサートを記述し、下記を実行する。

```
go test -update
```

対象サブテストに対応するディレクトリ(`testdata/golden_fixtures`)にフィクスチャーデータが作られるので、内容を確認する。

次回以降`go test`を実行するとフィクスチャーデータと一致するかチェックするようになる。


### マイグレーション実行
```
sql-migrate up
```

#### 状態確認
```
sql-migrate status
```
