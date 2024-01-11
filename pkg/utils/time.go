package utils

import (
	"fmt"
	"github.com/civet148/log"
	"strconv"
	"strings"
	"time"
)

const (
	HEIGHTS_ON_DAY          = 2880
	SECONDS_ONE_DAY         = 24 * 60 * 60
	TIME_FORMAT_DATE        = "2006-01-02"
	TIME_FORMAT_DATETIME    = "2006-01-02 15:04:05"
	FILECOIN_GENESIS_TIME   = 1598306400 //height is 0
	FILECOIN_BLOCK_DURATION = 30         //30s create a new block
)

func Now() string {
	return time.Now().Format(TIME_FORMAT_DATETIME)
}

func NowDate() string {
	return time.Now().Format(TIME_FORMAT_DATE)
}

func NowRandom() string {
	return time.Now().Format("20060102150405.000000000")
}

func NowHeight() int64 {
	now64 := time.Now().Unix()
	return UnixTimeToHeight(now64)
}

func Unix2DateTime(t64 uint64) string {
	if t64 == 0 {
		return ""
	}
	t := time.Unix(int64(t64), 0)
	return t.Format(fmt.Sprintf("%s", TIME_FORMAT_DATETIME))
}

func Unix2Date(t64 uint64) string {
	if t64 == 0 {
		return ""
	}
	t := time.Unix(int64(t64), 0)
	return t.Format(fmt.Sprintf("%s", TIME_FORMAT_DATE))
}

func Minute2Hour(minutes int) string {
	if minutes <= 0 {
		return "0h0m"
	}
	if minutes >= 60 {
		return fmt.Sprintf("%dh%dm", minutes/60, minutes%60)
	}
	return fmt.Sprintf("0h%dm", minutes)
}

func Now64() int64 {
	return time.Now().Unix()
}

func TodayZero() int64 {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return t.Unix()
}

func YesterDay(cur64 int64) int64 {
	t := time.Unix(cur64, 0)
	d, err := time.ParseDuration("-24h")
	if err != nil {
		log.Errorf("parse duration error [%s]", err)
		return 0
	}
	return t.Add(d).Unix()
}

func ISODateByUnix64(v int64) (t time.Time) {
	return time.Unix(v, 0)
}

func DateTimeStr2Unix(strDateTime string, timeFmt ...interface{}) (unixTime int64) {

	var t time.Time
	var bNormal = true

	if len(timeFmt) > 0 {

		if timeFmt[0].(string) == "/" {
			bNormal = false
		}
	} else {
		if strings.Contains(strDateTime, "/") {
			bNormal = false
		}
	}

	if len(strDateTime) != 19 {

		nIndex := strings.Index(strDateTime, " ")
		if nIndex == -1 {
			fmt.Println("error: DateTimeStr2Unix invalid datetime format")
			return 0
		}

		sdt := strings.Split(strDateTime, " ")
		if len(sdt) == 1 {
			fmt.Println("error: DateTimeStr2Unix invalid datetime format")
			return 0
		}

		ymd := sdt[0]
		hms := sdt[1]

		var s1, s2 []string
		if bNormal {
			s1 = strings.Split(ymd, "-")

		} else {
			s1 = strings.Split(ymd, "/")
		}

		s2 = strings.Split(hms, ":")

		if len(s1) != 3 || len(s2) != 3 {
			fmt.Println("error: DateTimeStr2Unix invalid datetime format, not match 'YYYY-MM-DD hh:mm:ss' or 'YYYY/MM/DD hh:mm:ss'")
			return 0
		}
		year := s1[0]
		month := s1[1]
		day := s1[2]
		hour := s2[0]
		min := s2[1]
		sec := s2[2]
		if len(year) != 4 {
			fmt.Println("error: DateTimeStr2Unix invalid year format, not match YYYY")
			return 0
		}
		if len(month) == 1 {
			month = "0" + month
		}
		if len(day) == 1 {
			day = "0" + day
		}
		if len(hour) == 1 {
			hour = "0" + hour
		}
		if len(min) == 1 {
			min = "0" + min
		}
		if len(sec) == 1 {
			sec = "0" + sec
		}

		if bNormal {
			strDateTime = fmt.Sprintf("%v-%v-%v %v:%v:%v", year, month, day, hour, min, sec)

		} else {
			strDateTime = fmt.Sprintf("%v/%v/%v %v:%v:%v", year, month, day, hour, min, sec)
		}
	}

	if strDateTime != "" {

		loc, _ := time.LoadLocation("Local")

		if bNormal {
			t, _ = time.ParseInLocation("2006-01-02 15:04:05", strDateTime, loc)
		} else {
			t, _ = time.ParseInLocation("2006/01/02 15:04:05", strDateTime, loc)
		}

		unixTime = t.Unix()
	}

	return
}

func GetDateBeginTime(strDate string) string {
	return strDate + " 00:00:00"
}

func GetDateEndTime(strDate string) string {
	return strDate + " 23:59:30"
}

func DateToEpochBeginTime64(strDate string) int64 {
	strDateTime := GetDateBeginTime(strDate)
	return DateTimeStr2Unix(strDateTime)
}

func DateToEpochEndTime64(strDate string) int64 {
	strDateTime := GetDateEndTime(strDate)
	return DateTimeStr2Unix(strDateTime)
}

func DateBeginHeight(strDate string) int64 {
	strDateTime := GetDateBeginTime(strDate)
	return DateTimeToHeight(strDateTime)
}

func DateEndHeight(strDate string) int64 {
	strDateTime := GetDateEndTime(strDate)
	return DateTimeToHeight(strDateTime)
}

func HeightToTimeUnix64(e int64) int64 {
	return HeightToTime(e).Unix()
}

func HeightToTime(e int64) (t time.Time) {
	unix64 := FILECOIN_GENESIS_TIME + (e * FILECOIN_BLOCK_DURATION)
	return time.Unix(unix64, 0)
}

func HeightToDateTime(e int64) string {
	return HeightToTime(e).Format(TIME_FORMAT_DATETIME)
}

func HeightToDate(e int64) string {
	return HeightToTime(e).Format(TIME_FORMAT_DATE)
}

func UnixTimeToHeight(time64 int64) int64 {

	d := time64 - FILECOIN_GENESIS_TIME
	if d < 0 {
		return 0
	}
	return d / FILECOIN_BLOCK_DURATION
}

func DateTimeToHeight(strDateTime string) int64 {
	time64 := DateTimeStr2Unix(strDateTime)
	return UnixTimeToHeight(time64)
}

func GetLatestHeight() int64 {
	return UnixTimeToHeight(time.Now().Unix())
}

func TimestampToDate(t time.Time) string {
	return t.Format(TIME_FORMAT_DATE)
}

func DateIsValid(strDate string) bool {

	if strDate == "" || len(strDate) < 6 {
		return false
	}
	if strDate[0] == '0' {
		return false
	}
	strDate = strings.Replace(strDate, "-", "", -1)
	strDate = strings.Replace(strDate, "/", "", -1)
	if date, err := strconv.Atoi(strDate); err != nil {
		return false
	} else {
		if len(strDate) == 6 {
			if date >= 197011 {
				return true
			}
		} else if len(strDate) == 8 {
			if date >= 19700101 {
				return true
			}
		}
	}
	return false
}

func DateStr2Unix(strDate string) int64 {

	t, err := time.Parse(TIME_FORMAT_DATE, strDate)
	if err != nil {
		log.Errorf("date string [%s] parse error [%s]", strDate, err.Error())
		return 0
	}
	return t.Unix()
}

func DateLessThan(strDate1, strDate2 string) bool {
	d1 := DateStr2Unix(strDate1)
	d2 := DateStr2Unix(strDate2)
	return d1 < d2
}

func DateLessThanNow(strDate string) bool {
	strNowDate := NowDate()
	return DateLessThan(strDate, strNowDate)
}
