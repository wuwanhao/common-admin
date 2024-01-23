// 时间工具类
package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var (
	formatTime = "2006-01-02 15:04:05"
)

// 为HTime结构体定义一个MarshalJSON的方法，该方法实现了Json.Marshaler接口，用于自定义HTime结构体如何被转换为JSON格式
func (t HTime)MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	// 将其转换为字节切片并返回
	return []byte(formatted), nil
}

func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	*t = HTime{Time: now}
	return
}

func (t HTime)Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
		
	return t.Time, nil
}

func (t *HTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil }
	return fmt.Errorf("can not convert %v to timestamp", v)
}