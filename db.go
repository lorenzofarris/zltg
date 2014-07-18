package main

import (
	"os"
	"fmt"
	"log"
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Cedict struct {
	Id            int64  
	Traditional   string 
	Simplified    string 
	Pinyin        string 
	English       string
	Line          string
}

func cedict_map(dbmap *gorp.DbMap) *gorp.TableMap {
	tm := dbmap.AddTableWithName(Cedict{}, "cedict")
	tm.ColMap("Id").Rename("id")
	tm.ColMap("Traditional").Rename("traditional").SetMaxSize(255)
	tm.ColMap("Simplified").Rename("simplified").SetMaxSize(255)
	tm.ColMap("Pinyin").Rename("pinyin").SetMaxSize(255)
	tm.ColMap("English").Rename("english").SetMaxSize(1024)
	tm.ColMap("Line").Rename("line").SetMaxSize(2048)
	tm.SetKeys(true, "Id")
	return tm
}

type Card struct{
	Id          int64    `db:"id"`
	Traditional string   `db:"traditional"`
	Simplified  string   `db:"simplified"`
	Pinyin      string   `db:"pinyin"`
	English     string   `db:"english"`
}

func card_map(dbmap *gorp.DbMap) *gorp.TableMap {
	tm := dbmap.AddTableWithName(Card{},"cards");
	tm.ColMap("Pinyin").SetMaxSize(1024)
	tm.ColMap("English").SetMaxSize(2048)
	return tm
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDb(db_name string) {
	
	/****************************
   * set up and test database
   ****************************/
  //db_name, _ := config["db"]
	db, err := sql.Open("sqlite3", db_name)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "zltg", log.Lmicroseconds))
	cedict_map(dbmap)
	tmap := card_map(dbmap)
	log.Printf("schema string is %s", tmap.SchemaName)
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")
	var entries []Cedict
  _, err = dbmap.Select(&entries, "select * from cedict limit 10")
	CheckErr(err, "failed select")
	for _, record := range entries {
		fmt.Printf("%s / %s / %s \n", record.Simplified, record.Pinyin, record.English)
	}
}
