package domain

import "time"

const (
	Card    = "card"
	Virtual = "virtual"
)

// 爬虫获取的回复以及直接通过爬虫获取数据的返回类型
type ResponseRecordsOfConsumption struct {
	//StudentId string
	SMT_TIMES string `json:"SMT_TIMES"` //消费次数
	//SMT_INMONEY  string `json:"SMT_INMONEY"`
	//SMT_OUTMONEY string `json:"SMT_OUTMONEY"`
	//MONEY            string `json:"MONEY"`
	SMT_DEALDATETIME string `json:"SMT_DEALDATETIME"` //消费时间
	SMT_ORG_NAME     string `json:"SMT_ORG_NAME"`     //窗口
	SMT_DEALNAME     string `json:"SMT_DEALNAME"`     //消费方式
	//BeforeMoney      string `json:"before_money"`
	AfterMoney string `json:"after_money"` //余额
	Money      string `json:"money"`       //本次消费金额
}

// 作为在数据库获取类型以及返回给前端类型
type Records struct {
	Times uint16 //消费次数
	//SMT_INMONEY  string `json:"SMT_INMONEY"`
	//SMT_OUTMONEY string `json:"SMT_OUTMONEY"`
	//MONEY            string `json:"MONEY"`
	DealTime   time.Time //消费时间
	DealWindow string    //窗口
	DealWay    string    //消费方式
	//BeforeMoney      string `json:"before_money"`
	AfterMoney float32
	Money      float32
}

// 用于将数据存进数据库
type RecordsInRepository struct {
	User      User `gorm:"foreignKey:StudentId;references:StudentId"`
	StudentId string
	Times     uint16 //消费次数
	//SMT_INMONEY  string `json:"SMT_INMONEY"`
	//SMT_OUTMONEY string `json:"SMT_OUTMONEY"`
	//MONEY            string `json:"MONEY"`
	DealTime   time.Time //消费时间
	DealWindow string    //窗口
	DealWay    string    //消费方式
	//BeforeMoney      string `json:"before_money"`
	AfterMoney float32
	Money      float32
}
