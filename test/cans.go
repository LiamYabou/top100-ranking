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
	CannedRankings = `{"status":"success","data":{"categories":[{"id":3,"name":"Amazon Device Accessories"},{"id":4,"name":"Amazon Devices"}],"products":[{"name":"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019","rank":1,"image_path":"images/B0v6F98VAjwrgJ8U.jpg"},{"name":"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal","rank":2,"image_path":null},{"name":"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release","rank":3,"image_path":null},{"name":"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone","rank":4,"image_path":null},{"name":"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal","rank":5,"image_path":null}],"root_category":{"id":1,"name":"Any Department"},"selected_category_name":"Amazon Devices \u0026 Accessories"}}`
)
