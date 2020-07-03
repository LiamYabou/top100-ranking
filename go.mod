// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/LiamYabou/top100-ranking

go 1.14

require (
	github.com/LiamYabou/top100-pkg v0.0.0-20200622082036-12669ddb8700
	github.com/alexflint/go-filemutex v1.1.0 // indirect
	github.com/jackc/pgx/v4 v4.7.1
	github.com/khaiql/dbcleaner v2.3.0+incompatible // indirect
	github.com/lib/pq v1.7.0
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	github.com/panjf2000/ants/v2 v2.4.1
	github.com/romanyx/polluter v1.2.2
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.0
	gopkg.in/khaiql/dbcleaner.v2 v2.3.0
)
