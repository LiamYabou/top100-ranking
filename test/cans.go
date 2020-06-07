package test

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/khaiql/dbcleaner.v2"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	DBpool *pgxpool.Pool
	PQconn *sql.DB
	Cleaner dbcleaner.DbCleaner
	CannedRankings = `{"status":"success","data":{"categories":[{"name":"Amazon Device Accessories","url":"https://www.amazon.com/Best-Sellers-Amazon-Device-Accessories/zgbs/amazon-devices/370783011/ref=zg_bs_nav_1_amazon-devices/145-7972861-4524441"},{"name":"Amazon Devices","url":"https://www.amazon.com/Best-Sellers-Amazon-Devices/zgbs/amazon-devices/2102313011/ref=zg_bs_nav_1_amazon-devices/145-7972861-4524441"}],"products":[{"name":"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019","rank":1},{"name":"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal","rank":2},{"name":"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release","rank":3},{"name":"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone","rank":4},{"name":"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal","rank":5}],"selected_category_entry":"Amazon Devices \u0026 Accessories"}}`
)