package models

import ( 
	"github.com/go-xorm/xorm"
	"strings"     
	"fmt"     
	_ "github.com/go-sql-driver/mysql"
)
var Eng *xorm.Engine
const (
TimeFormat = "2006-01-02 15:04:05"
)
func RegisterDB() {
	DBtype := "mysql"
	url := "tcp(127.0.0.1:3306)/gocms?charset=utf8"
	name := "root"
	password := "root"
	dbUrl := strings.Join([]string{strings.Join([]string{name, password}, ":"), url}, "@")
	var err error
    Eng, err = xorm.NewEngine(DBtype, dbUrl)
    if err != nil{
       fmt.Printf(err.Error())
    }
	Eng.ShowSQL = true
  
}

 
