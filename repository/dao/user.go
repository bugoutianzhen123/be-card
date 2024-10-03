package dao

import (
	"context"
	"fmt"
	"github.com/asynccnu/be-card/domain"
)

// 记录学号及key
func (dao *CardDao) NoteKey(ctx context.Context, key string, StudentId string) error {
	user := domain.User{
		Key:          key,
		StudentId:    StudentId,
		CardState:    0,
		VirtualState: 0,
	}
	return dao.db.WithContext(ctx).Create(&user).Error
}

// 通过学号查询key
func (dao *CardDao) GetKeyByStudentId(ctx context.Context, StudentId string) (string, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).First(&user).Error
	return user.Key, err
}

// 通过学号获取CardState
func (dao *CardDao) GetCardStateByStudentId(ctx context.Context, StudentId string) (uint64, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).First(&user).Error
	return user.CardState, err
}

// 通过学号获取VirtualState
func (dao *CardDao) GetVirtualStateByStudentId(ctx context.Context, StudentId string) (uint64, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).First(&user).Error
	return user.VirtualState, err
}

// 通过学号获取用户信息
func (dao *CardDao) GetUserInfoByStudentId(ctx context.Context, StudentId string) (domain.User, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).Find(&user).Error
	return user, err
}

// 更新key
func (dao *CardDao) UpdateKeyByStudentId(ctx context.Context, StudentId string, newKey string) error {
	return dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).Update("key", newKey).Error
}

// 更新state
func (dao *CardDao) UpdateStateByStudentId(ctx context.Context, StudentId string, newState uint64, kind string) error {
	field := fmt.Sprintf("%s_state", kind)
	return dao.db.WithContext(ctx).Model(&domain.User{}).Where("student_id = ?", StudentId).Update(field, newState).Error
}
