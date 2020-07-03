// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/LiamYabou/top100-ranking

go 1.14

require (
	github.com/LiamYabou/top100-pkg v0.0.1
	github.com/alexflint/go-filemutex v1.1.0 // indirect
	github.com/getsentry/sentry-go v0.6.1
	github.com/jackc/pgx/v4 v4.7.1
	github.com/khaiql/dbcleaner v2.3.0+incompatible // indirect
	github.com/lib/pq v1.7.0
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	github.com/panjf2000/ants/v2 v2.4.1
	github.com/romanyx/polluter v1.2.2
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/khaiql/dbcleaner.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
