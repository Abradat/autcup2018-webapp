package main

import (
	_ "autcup/routers"
	models "autcup/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)


func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "models/aut_db.db")
	orm.RegisterModel(new(models.Team), new(models.Challenge), new(models.Score), new(models.Draw))
}


func scoreWithRatio(score float32, ratio float32)(finalScore float32){
	finalScore = ratio * score
	return
}


func main() {
	beego.AddFuncMap("scoreWithRatio",scoreWithRatio)
	beego.Run()
}

