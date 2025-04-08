package intermediate_level

import "fmt"

type Database interface {
    Query(sql string) string
}

type MockDB struct {}
type PostgreDB struct {}

func (mdb *MockDB) Query(sql string) string{
    return fmt.Sprintf("mocking '%s'", sql)
}

func (pdb *PostgreDB) Query(sql string) string{
    return fmt.Sprintf("sending to Postgreee '%s'", sql)
}

func FetchData(db Database){
    fmt.Println(db.Query("Select * from db"))
}
