package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/civet148/log"
)

func GenerateSalt() string {

	return ""
}

func PasswordMD5(strPassword, strSalt string) (strMD5 string) {
	m := md5.New()
	strEnc := fmt.Sprintf("%s%s", strPassword, strSalt)
	if _, err := m.Write([]byte(strEnc)); err != nil {
		log.Errorf("md5 write error [%s]", err.Error())
		return ""
	}
	return fmt.Sprintf("%x", m.Sum(nil))
}
