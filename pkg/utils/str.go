package utils

import "strings"

func TrimSpace(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func StrMap2Uint(s string, max uint64) (idx uint) {
	var m uint64
	for _, v := range s {
		m += uint64(v)
	}
	return uint(m % max)
}

func UrlSuffix(strUri string) (strRouter, strSuffix string) {
	idx := strings.LastIndex(strUri, "/")
	if idx == 0 {
		strRouter = strUri[0:1]
		strSuffix = strUri[idx+1:]
	} else if idx > 0 {
		strRouter = strUri[:idx]
		strSuffix = strUri[idx+1:]
	}
	return
}
