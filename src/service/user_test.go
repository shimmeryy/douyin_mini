package service

import (
	"fmt"
	"testing"
	"tiktok/src/dal/db"
)

func Test_user_checkUserById(t *testing.T)  {
	db.Init()
	//存在的id
	fmt.Println(UserServiceInstance().checkUserById(nil, 1))
	//不存在的id
	fmt.Println(UserServiceInstance().checkUserById(nil, 2))
}