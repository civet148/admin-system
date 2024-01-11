package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	UNIT_EiB     = "EiB"
	UNIT_PIB     = "PiB"
	UNIT_DOLLAR  = "$"
	UNIT_BILLION = "billion"
)

const (
	UNIT_ALL = UNIT_EiB + "|" + UNIT_PIB + "|" + UNIT_DOLLAR + "|" + UNIT_BILLION
)

func MakeTimestampSuffix(strKey string) string {
	now := time.Now().Unix()
	return fmt.Sprintf("%s_%d", strKey, now)
}

func ConvertFloat64NonUnit(strNumber string) float64 {
	units := strings.Split(UNIT_ALL, "|")
	for _, unit := range units {
		strNumber = strings.Replace(strNumber, unit, "", -1)
	}
	strNumber = strings.TrimSpace(strNumber)
	number, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		fmt.Printf("parse float (%s) error [%s]", strNumber, err)
		return 0
	}
	return number
}

func CheckDurationMinutes(strDuration string) (ok bool, minutes int64, err error)	{
	if strings.Contains(strDuration, "m") {
		ok =true
		strMinutes := strings.TrimSuffix(strDuration, "m")
		minutes, err = strconv.ParseInt(strMinutes, 10, 32)
	}
	return
}

func CheckDurationHours(strDuration string) (ok bool, hours int64, err error)	{
	if strings.Contains(strDuration, "h") {
		ok =true
		strHours := strings.TrimSuffix(strDuration, "h")
		hours, err = strconv.ParseInt(strHours, 10, 32)
	}
	return
}