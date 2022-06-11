package config

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	Init()
	fmt.Println(AppConfig)
	err := errors.New("123")
	test01(err)
}

func Init() {

}

func test01(v interface{}) {
	fmt.Println(v)
}
