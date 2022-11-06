package DB

import "database/sql"

var DD *sql.DB
var DBpassword string = env("password")
var DBname string = env("dbname")
var DBuser string = env("user")
var DBhost string = env("host")
var DBport int = envint("port")
