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

### マイグレーション実行
```

```

#### 状態確認
```
sql-migrate status
```
