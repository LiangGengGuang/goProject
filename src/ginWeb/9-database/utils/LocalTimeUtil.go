package utils

import (
	"fmt"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2023/6/23

//MyTime 自定义时间
type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
