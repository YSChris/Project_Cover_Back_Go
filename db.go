package dataFromDB

import (
	"database/sql"
    "fmt"
	"log"
	"time"
    _ "github.com/denisenkom/go-mssqldb"
)

type dbConn struct{
	server string
    user string
    password string
    database string
    encrypt string
}

func (connection dbConn) connect() (*sql.DB){
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", 
		connection.server, connection.database, connection.user, connection.password, connection.encrypt)
	
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	return conn
}

func (connection dbConn) Sql(sqlStr string){
	
    conn := connection.connect()
    
    defer conn.Close()

    stmt, err := conn.Prepare(sqlStr)
    if err != nil {
        log.Fatal("Prepare failed:", err.Error())
    }
    defer stmt.Close()

    rows, err := stmt.Query()
    if err != nil {
        log.Fatal("Query failed:", err.Error())
    }

    cols, err := rows.Columns()
    var colsdata = make([]interface{}, len(cols))
    for i := 0; i < len(cols); i++ {
        colsdata[i] = new(interface{})
        fmt.Print(cols[i])
        fmt.Print("\t")
    }
    fmt.Println()

    for rows.Next() {
        rows.Scan(colsdata...)
        fmt.Print(colsdata)
    }
    defer rows.Close()
}

func PrintRow(colsdata []interface{}) {
    for _, val := range colsdata {
        switch v := (*(val.(*interface{}))).(type) {
        case nil:
            fmt.Print("NULL")
        case bool:
            if v {
                fmt.Print("True")
            } else {
                fmt.Print("False")
            }
        case []byte:
            fmt.Print(string(v))
        case time.Time:
            fmt.Print(v.Format("2016-01-02 15:05:05.999"))
        default:
            fmt.Print(v)
        }
        fmt.Print("\t")
    }
    fmt.Println()
}