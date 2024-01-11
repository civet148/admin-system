package utils

import (
	"fmt"
	"github.com/fatih/color"
	"regexp"
	"strings"
)

func ColorBool(ok bool) string {
	if ok {
		return color.GreenString("true")
	}
	return "false"
}

func ColorWarn(strFmt string, args ...interface{}) string {
	if strFmt != "" {
		return color.YellowString(strFmt, args...)
	}
	return ""
}

func ColorError(strFmt string, args ...interface{}) string {
	if strFmt != "" {
		return color.RedString(strFmt, args...)
	}
	return ""
}

func BlueString(v interface{}) string {
	return color.BlueString("%v", v)
}

func GreenString(v interface{}) string {
	return color.GreenString("%v", v)
}

func YellowString(v interface{}) string {
	return color.YellowString("%v", v)
}

func RedString(v interface{}) string {
	return color.RedString("%v", v)
}

func RedStringNotOK(v interface{}) string {
	str := fmt.Sprintf("%v", v)
	str = strings.ToLower(str)
	if str == "ok" {
		return str
	}
	return color.RedString("%v", v)
}

func CyanString(v interface{}) string {
	return color.CyanString("%v", v)
}

func RedInt(v int) string {
	if v == 0 {
		return fmt.Sprintf("%v", v)
	}
	return color.RedString("%v", v)
}

// phoneNumber
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,1,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
