package config

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	init()
	fmt.Println(AppConfig)
	err := errors.New("123")
	test01(err)
}

func init() {

}

func test01(v any) {
	fmt.Println(v)
}
