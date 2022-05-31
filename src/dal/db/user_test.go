package db

import (
	"fmt"
	"testing"
)

func Test_db_user(t *testing.T)  {
	Init()
	fmt.Println(QueryUser(nil, "changlu"))
}