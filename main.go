package main

import(
	_ "github.com/denisenkom/go-mssqldb"
	"dataFromDB"
)

//Server=120.196.136.235;database=mclCoverSystem_Web_BAGX;User Id=sa;Password=maxt8899MAX
func main(){
    dbConn := Database.dbConn{
        server: "120.196.136.235",
        user: "sa",
        password: "maxt8899MAXT",
        database: "mclCoverSystem_Web_BAGX",
        encrypt: "disable",
    }
 
    dbConn.Sql("select F_Id, 序列号, IMSI码 from VW_BreakDownBill")
 }