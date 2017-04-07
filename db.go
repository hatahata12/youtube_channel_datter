package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
	"log"
)

type MyDB struct {
	con *genmai.DB
}

type Videos struct {
	Id           string `db:"pk"`
	Title        string
	Published_at string
	Thumbnails   string
	genmai.TimeStamp
}

var (
	db MyDB
)

func (this *MyDB) Connect() {
	var err error
	this.con, err = genmai.New(&genmai.MySQLDialect{}, "user:pass@(xxxxxx.com:3306)/db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
}

func (this *MyDB) Close() {
	var err error
	err = this.con.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(video Videos) {

	var results []Videos
	if err := db.con.Select(&results, db.con.Where("id", "=", video.Id)); err != nil {
		return
	}

	if len(results) != 0 {
		fmt.Printf("すでに登録されています %s \n", video.Id)
	} else {
		obj := &video
		_, err := db.con.Insert(obj)
		if err != nil {
			panic(err)
		}
		fmt.Printf("登録しました %s \n", video.Id)
	}

}
