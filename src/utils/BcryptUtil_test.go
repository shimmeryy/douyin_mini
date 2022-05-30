package utils

import (
	"fmt"
	"testing"
)

func Test_BcryptUtil(t *testing.T) {
	salt, _ := HashAndSalt("123456")
	fmt.Println(salt) //$2a$04$i6RIyd4B5McvCx4GZtBNxudrpT8Z3oeuVbMNNgS2hyg26Y0R8TNDO
	fmt.Println(ComparePasswords(salt, "123456"))
}
