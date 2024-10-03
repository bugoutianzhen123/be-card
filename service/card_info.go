package service

import (
	"fmt"
	v1 "github.com/asynccnu/be-api/gen/proto/card/v1"
	"github.com/asynccnu/be-card/domain"
	"github.com/asynccnu/be-card/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

const (
	Year      = "2006"
	YearMonth = "2006 01"
	SetTime   = "2006-01-02 15:04:05"
)

// 判断是否是本月
func CheckMonth(start_time string) bool {
	t, err := time.Parse(YearMonth, start_time)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	now := time.Now()
	if t.Year() == now.Year() && t.Month() == now.Month() {
		return true
	} else {
		return false
	}
}

// 查看目标月份是否在数据库
func CheckState(user domain.User, msg domain.ServiceMsg) bool {
	month := int(msg.Stime.Year()-msg.GradeTime.Year())*12 + int(msg.Stime.Month()-9)
	switch msg.Type {
	case repository.Card:
		return user.CardState&(1<<month) != 0
	case repository.Virtual:
		return user.VirtualState&(1<<month) != 0
	}
	return false
}

// 将爬虫结果转为返回结果
func RecordsOfConsumptionToResponse(info []domain.ResponseRecordsOfConsumption) ([]*v1.RecordOfConsumption, error) {
	var results []*v1.RecordOfConsumption
	for _, record := range info {
		times, _ := strconv.Atoi(record.SMT_TIMES)
		dealTime, _ := time.Parse(SetTime, record.SMT_DEALDATETIME)
		money, _ := strconv.ParseFloat(record.Money, 32)
		afterMoney, _ := strconv.ParseFloat(record.AfterMoney, 32)
		results = append(results, &v1.RecordOfConsumption{
			SMT_TIMES:        uint32(times),
			SMT_DEALDATETIME: timestamppb.New(dealTime),
			SMT_ORG_NAME:     record.SMT_ORG_NAME,
			SMT_DEALNAME:     record.SMT_DEALNAME,
			AfterMoney:       float32(afterMoney),
			Money:            float32(money),
		})
	}
	return results, nil
}

// 将数据库查询结果转化为返回结果
func RecordsToResponse(info []domain.Records) ([]*v1.RecordOfConsumption, error) {
	var results []*v1.RecordOfConsumption
	for _, record := range info {
		results = append(results, &v1.RecordOfConsumption{
			SMT_TIMES:        uint32(record.Times),
			SMT_DEALDATETIME: timestamppb.New(record.DealTime),
			SMT_ORG_NAME:     record.DealWindow,
			SMT_DEALNAME:     record.DealWay,
			AfterMoney:       record.AfterMoney,
			Money:            record.Money,
		})
	}
	return results, nil
}

// 为数据添加StudentId用于存储
func GetCardRecordsForRepository(info []domain.ResponseRecordsOfConsumption, StudentId string) ([]domain.RecordsInRepository, error) {
	var results []domain.RecordsInRepository
	for _, record := range info {
		times, _ := strconv.Atoi(record.SMT_TIMES)
		dealTime, _ := time.Parse("2006-01-02 15:04:05", record.SMT_DEALDATETIME)
		money, _ := strconv.ParseFloat(record.Money, 32)
		afterMoney, _ := strconv.ParseFloat(record.AfterMoney, 32)
		results = append(results, domain.RecordsInRepository{
			StudentId:  StudentId,
			Times:      uint16(times),
			DealTime:   dealTime,
			DealWindow: record.SMT_ORG_NAME,
			DealWay:    record.SMT_DEALNAME,
			AfterMoney: float32(afterMoney),
			Money:      float32(money),
		})
	}
	return results, nil
}

//// 数据转化,获取记录
//func GetCardRecords(info []domain.ResponseRecordsOfConsumption) ([]domain.Records, error) {
//	var results []domain.Records
//	for _, record := range info {
//		times, _ := strconv.Atoi(record.SMT_TIMES)
//		dealTime, _ := time.Parse(record.SMT_DEALDATETIME, record.SMT_DEALDATETIME)
//		money, _ := strconv.ParseFloat(record.Money, 32)
//		afterMoney, _ := strconv.ParseFloat(record.AfterMoney, 32)
//		results = append(results, domain.Records{
//			Times:      uint16(times),
//			DealTime:   dealTime,
//			DealWindow: record.SMT_ORG_NAME,
//			DealWay:    record.SMT_DEALNAME,
//			AfterMoney: float32(afterMoney),
//			Money:      float32(money),
//		})
//	}
//	return results, nil
//}
