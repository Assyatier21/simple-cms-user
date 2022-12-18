package database

import (
	"cms/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init() {
	conn := "user=" + config.User + " password= " + config.Password + " host=" + config.Host + " port= " + config.Port + " dbname=" + config.Database + " sslmode=" + config.Sslmode
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("[Database] can't connect to database: ", err.Error())
		return
	}

	log.Println("[Database] successfully connected")
}
