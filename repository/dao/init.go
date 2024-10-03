package dao

import (
	"context"
	"fmt"
	"github.com/asynccnu/be-card/domain"
	"github.com/asynccnu/be-card/pkg/logger"
	"gorm.io/gorm"
	"time"
)

type Dao interface {
	NoteKey(ctx context.Context, key string, StudentId string) error
	GetKeyByStudentId(ctx context.Context, StudentId string) (string, error)
	UpdateKeyByStudentId(ctx context.Context, StudentId string, newKey string) error
	UpdateRecordsAndUserInfo(ctx context.Context, info []domain.RecordsInRepository, msg domain.ServiceMsg, user domain.User) error
	//UpdateStateByStudentId(StudentId string, newState uint64, kind string) error
	//NoteRecords(info []domain.RecordsInRepository, kind string, grade string) error
	//NoteCardRecords(info []domain.RecordsInRepository, grade string) error
	GetCardStateByStudentId(ctx context.Context, StudentId string) (uint64, error)
	GetVirtualStateByStudentId(ctx context.Context, StudentId string) (uint64, error)
	GetUserInfoByStudentId(ctx context.Context, StudentId string) (domain.User, error)
	GetVirtualRecordsByStudentId(ctx context.Context, StudentId string, grade string, startTime time.Time) ([]domain.Records, error)
	GetCardRecordsByStudentId(ctx context.Context, StudentId string, grade string, startTime time.Time) ([]domain.Records, error)
}

type CardDao struct {
	db *gorm.DB
	l  logger.Logger
}

func NewCardDao(db *gorm.DB) Dao {
	return &CardDao{db: db}
}

// 根据年级创建表
func InitTables(db *gorm.DB) error {
	db.AutoMigrate(&domain.User{})
	// 获取当前时间
	now := time.Now()
	// 定义2021年9月的时间
	//start := time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC)
	// 计算时间差
	//duration := now.Sub(start)
	// 计算经过的年数
	//year := int(duration.Hours() / (365.25 * 24)) // 考虑闰年
	year := now.Year()
	var years []int
	i := 0
	for i = 0; i <= 4; i++ {
		years = append(years, year-i)
	}

	for _, y := range years {
		tableNameCard := fmt.Sprintf("card_records_%d", y)
		tableNameVirtual := fmt.Sprintf("virtual_records_%d", y)
		if err := db.Table(tableNameCard).AutoMigrate(&domain.RecordsInRepository{}); err != nil {
			return err
		}
		if err := db.Table(tableNameVirtual).AutoMigrate(&domain.RecordsInRepository{}); err != nil {
			return err
		}
	}

	return db.Error
}
