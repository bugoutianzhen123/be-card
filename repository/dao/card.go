package dao

import (
	"context"
	"fmt"
	"github.com/asynccnu/be-card/domain"
	"time"
)

const (
	Card    = domain.Card
	Virtual = domain.Virtual
)

// 记录虚拟卡消费记录
func (dao *CardDao) NoteRecords(ctx context.Context, info []domain.RecordsInRepository, kind string, grade string) error {
	tableNameVirtual := fmt.Sprintf("%s_records_%s", kind, grade)
	return dao.db.WithContext(ctx).Table(tableNameVirtual).Create(&info).Error
}

//// 记录实体卡消费记录
//func (dao *CardDao) NoteCardRecords(info []domain.RecordsInRepository, grade string) error {
//	tableNameCard := fmt.Sprintf("card_records_%s", grade)
//	return dao.db.Table(tableNameCard).Create(&info).Error
//}

// 获取虚拟卡消费记录
func (dao *CardDao) GetVirtualRecordsByStudentId(ctx context.Context, StudentId string, grade string, startTime time.Time) ([]domain.Records, error) {
	var info []domain.Records
	tableName := fmt.Sprintf("virtual_records_%s", grade)
	endTime := startTime.AddDate(0, 1, 0) // 下个月的第一天
	err := dao.db.WithContext(ctx).Table(tableName).Where("student_id = ? And deal_time >= ? AND deal_time < ?", StudentId, startTime, endTime).Find(&info).Error
	return info, err
}

// 获取实体卡消费记录
func (dao *CardDao) GetCardRecordsByStudentId(ctx context.Context, StudentId string, grade string, startTime time.Time) ([]domain.Records, error) {
	var info []domain.Records
	tableName := fmt.Sprintf("card_records_%s", grade)
	endTime := startTime.AddDate(0, 1, 0) // 下个月的第一天
	err := dao.db.WithContext(ctx).Table(tableName).Where("student_id = ? And deal_time >= ? AND deal_time < ?", StudentId, startTime, endTime).Find(&info).Error
	return info, err
}

// 更新信息
func (dao *CardDao) UpdateRecordsAndUserInfo(ctx context.Context, info []domain.RecordsInRepository, msg domain.ServiceMsg, user domain.User) error {
	month := int(msg.Stime.Year()-msg.GradeTime.Year())*12 + int(msg.Stime.Month()-9)
	tx := dao.db.Begin()

	err := dao.NoteRecords(ctx, info, msg.Type, msg.Grade)
	if err != nil {
		tx.Rollback()
		return err
	}

	switch msg.Type {
	case Card:
		err := dao.UpdateStateByStudentId(ctx, user.StudentId, user.CardState|(1<<month), msg.Type)
		if err != nil {
			tx.Rollback()
			return err
		}
	case Virtual:
		err := dao.UpdateStateByStudentId(ctx, user.StudentId, user.VirtualState|(1<<month), msg.Type)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
