package db

import (
	"fmt"
	"testing"
)

func TestCreatFollow(t *testing.T) {
	Init()
	err := CreatFollow(nil, 1, 2)
	if err != nil {
		return
	}
}

func TestDeleteFollow(t *testing.T) {
	Init()
	err := DeleteFollow(nil, 1, 2)
	if err != nil {
		return
	}
}

func TestQueryFansListById(t *testing.T) {
	Init()
	fans, err := QueryFansListById(nil, 2)
	if err != nil {
		return
	}
	for i := 0; i < len(fans); i++ {
		fmt.Println(fans[i])
	}
}

func TestQueryFollowerListById(t *testing.T) {
	Init()
	followers, err := QueryFollowerListById(nil, 1)
	if err != nil {
		return
	}
	for i := 0; i < len(followers); i++ {
		fmt.Println(i, followers[i])
	}
}

func Test_isFollow(t *testing.T) {
	Init()
	follow, err := IsFollow(nil, 1, 2)
	if err != nil {
		return
	}
	fmt.Println(follow)
}
