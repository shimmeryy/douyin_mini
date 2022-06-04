package db

import (
	"fmt"
	"testing"
)

func Test_db_user(t *testing.T) {
	Init()
	user, _ := QueryUser(nil, "changlu")
	fmt.Println(user)
}
