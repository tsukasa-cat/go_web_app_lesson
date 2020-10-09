package Infrastructure

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    // _ "github.com/go-sql-driver/mysql" ← こっちでも動く
)

type DataBase struct {
    Host string
    Username string
    Password string
    DBName string
    Connection *gorm.DB
}

func NewDB() *DataBase {
    c := NewConfig()
    return newDB(&DataBase{
        Host: c.DataBase.Development.Host,
        Username: c.DataBase.Development.Username,
        Password: c.DataBase.Development.Password,
        DBName: c.DataBase.Development.DBName,
    })
}

/**
 * func Open(dialect string, args ...interface{}) (db *DB, err error)
 * https://godoc.org/github.com/jinzhu/gorm#Open
 */
func newDB(d *DataBase) *DataBase {
    //
    // ex) MySQL
    // https://github.com/go-sql-driver/mysql#examples
    //
    // ex) MySQL Parameters
    // https://github.com/go-sql-driver/mysql#parameters
    db, err := gorm.Open("mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
    }
    d.Connection = db
    return d
}

// Begin begins a transaction
func (db *DataBase) Begin() *gorm.DB {
    return db.Connection.Begin()
}

func (db *DataBase) Connect() *gorm.DB {
    return db.Connection
}