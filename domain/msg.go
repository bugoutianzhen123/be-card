package domain

import (
	"fmt"
	"time"
)

const (
	Year      = "2006"
	YearMonth = "2006 01"
)

// 服务需要的信息
type ServiceMsg struct {
	StudentId string //初始
	Key       string //初始
	StartTime string //初始
	Stime     time.Time
	Type      string //初始 获取即判断type格式是否正确
	Grade     string
	GradeTime time.Time
}

func (m *ServiceMsg) GetMsg() error {
	m.Grade = GetGrade(m.StudentId)
	var err error
	m.GradeTime, err = GetGradeTime(m.Grade)
	if err != nil {
		return err
	}
	m.Stime, err = GetStartTime(m.StartTime)
	if err != nil {
		return err
	}
	return nil
}

// 将start_time从string转换成
func GetStartTime(start_time string) (time.Time, error) {
	t, err := time.Parse(YearMonth, start_time)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return t, err
	}
	return t, nil
}

// 从学号 获取年级
func GetGrade(StudentId string) string {
	grade := StudentId[:4]
	return grade
}

// 将年级转换成年份
func GetGradeTime(grade string) (time.Time, error) {
	t, err := time.Parse(Year, grade)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return t, err
	}
	return t, nil
}
