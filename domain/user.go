package domain

const (
	FirstYearSeptember = 1 << iota
	FirstYearOctober
	FirstYearNovember
	FirstYearDecember
	FirstYearJanuary
	FirstYearFebruary
	FirstYearMarch
	FirstYearApril
	FirstYearMay
	FirstYearJune
	FirstYearJuly
	FirstYearAugust
	SecondYearSeptember
	SecondYearOctober
	SecondYearNovember
	SecondYearDecember
	SecondYearJanuary
	SecondYearFebruary
	SecondYearMarch
	SecondYearApril
	SecondYearMay
	SecondYearJune
	SecondYearJuly
	SecondYearAugust
	ThirdYearSeptember
	ThirdYearOctober
	ThirdYearNovember
	ThirdYearDecember
	ThirdYearJanuary
	ThirdYearFebruary
	ThirdYearMarch
	ThirdYearApril
	ThirdYearMay
	ThirdYearJune
	ThirdYearJuly
	ThirdYearAugust
	FourthYearSeptember
	FourthYearOctober
	FourthYearNovember
	FourthYearDecember
	FourthYearJanuary
	FourthYearFebruary
	FourthYearMarch
	FourthYearApril
	FourthYearMay
	FourthYearJune
	FourthYearJuly
	FourthYearAugust
)

type User struct {
	//学号
	StudentId string
	//用于查询的key
	Key string
	//利用位运算确定消费记录存储情况
	CardState    uint64
	VirtualState uint64
}
