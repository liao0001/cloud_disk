package utils

import "time"

func CurrTime() (t int64) {
	return time.Now().UnixNano() / 1e6
}

func NowMillion() (t int64) {
	return time.Now().UnixNano() / 1e6
}

/*
	仅支持1970年1月1日后的日期
*/

//将标准毫秒数转为时间
func ToDatetime(t int64) time.Time {
	return time.Unix(t/1e3, (t%1e3)*1e3).Local()
}

const (
	DateTimeForDate          = "2006-01-02"
	DatetimeForDateTime      = "2006-01-02 15:04:05"
	DatetimeForShortDateTime = "2006-01-02 15:04"
)

//将时间转为固定格式
func DateFormat(t int64, format ...string) string {
	if t == 0 {
		return ""
	}
	if len(format) == 0 {
		return ToDatetime(t).Format(DatetimeForDateTime)
	}
	return ToDatetime(t).Format(format[0])
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
