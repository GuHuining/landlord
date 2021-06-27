/*
** @Title		tools/tools.go
** @Description	Save tool functions
** @Author		Huining
** @Created		2021-05-30 00:28:20
 */
package tools

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
)


// Md5EncodingPassword md5加密密码
func Md5EncodingPassword(password string) (string, string) {
	saltBytes := make([]byte, 32)
	rand.Read(saltBytes)                 // generate salt randomly
	salt := fmt.Sprintf("%x", saltBytes) // convert saltBytes from []byte to hex string
	passwordBytes := []byte(password)
	passwordBytes = append(passwordBytes, []byte(salt)...) // concat origin password and salt
	encodedPassword := md5.Sum(passwordBytes)              // use md5 to encode
	return fmt.Sprintf("%x", encodedPassword), salt
}

// ValidatePassword 判断加密后的密码是否相符
func ValidatePassword(encoded, password, salt string) bool {
	passwordBytes := append([]byte(password), []byte(salt)...)
	encodedPassword := fmt.Sprintf("%x", md5.Sum(passwordBytes))
	return encoded == encodedPassword
}
