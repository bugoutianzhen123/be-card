package repository

import (
	"context"
	"github.com/asynccnu/be-card/domain"
	"github.com/asynccnu/be-card/pkg/logger"
	"github.com/asynccnu/be-card/repository/cache"
	"github.com/asynccnu/be-card/repository/dao"
	"github.com/pingcap/errors"
	"time"
)

const (
	Card    = domain.Card
	Virtual = domain.Virtual
)

type CardRepository interface {
	NoteKey(ctx context.Context, key string, StudentId string) error
	GetKeyByStudentId(ctx context.Context, StudentId string) (string, error)
	GetStateByStudentId(ctx context.Context, StudentId string, kind string) (uint64, error)
	GetUserInfoByStudentId(ctx context.Context, StudentId string) (domain.User, error)
	UpdateKeyByStudentId(ctx context.Context, StudentId string, newKey string) error
	UpdateRecordsAndUserInfo(ctx context.Context, info []domain.RecordsInRepository, msg domain.ServiceMsg, user domain.User) error

	//UpdateStateByStudentId(StudentId string, newState uint64, kind string) error

	//NoteRecords(info []domain.RecordsInRepository, grade string, kind string) error
	//NoteCardRecords(info []domain.Records, grade string) error
	GetRecordsByStudentId(ctx context.Context, StudentId string, kind string, grade string, time time.Time) ([]domain.Records, error)
	//GetCardRecordsByStudentId(StudentId string) ([]domain.Records, error)
}

type CachedRepository struct {
	dao   dao.Dao
	cache cache.Cache
	l     logger.Logger
}

func NewCardRepository(dao dao.Dao, cache cache.Cache) CardRepository {
	return &CachedRepository{dao: dao, cache: cache}
}

func (repo *CachedRepository) NoteKey(ctx context.Context, key string, StudentId string) error {
	return repo.dao.NoteKey(ctx, key, StudentId)
}

func (repo *CachedRepository) GetKeyByStudentId(ctx context.Context, StudentId string) (string, error) {
	return repo.dao.GetKeyByStudentId(ctx, StudentId)
}

func (repo *CachedRepository) GetStateByStudentId(ctx context.Context, StudentId string, kind string) (uint64, error) {
	switch kind {
	case Card:
		return repo.dao.GetCardStateByStudentId(ctx, StudentId)
	case Virtual:
		return repo.dao.GetVirtualStateByStudentId(ctx, StudentId)
	default:
		return 0, errors.Errorf("invalid kind %s", kind)
	}

}

func (repo *CachedRepository) GetUserInfoByStudentId(ctx context.Context, StudentId string) (domain.User, error) {
	return repo.dao.GetUserInfoByStudentId(ctx, StudentId)
}

func (repo *CachedRepository) UpdateKeyByStudentId(ctx context.Context, StudentId string, newKey string) error {
	return repo.dao.UpdateKeyByStudentId(ctx, StudentId, newKey)
}

//func (repo *CachedRepository) UpdateStateByStudentId(StudentId string, newState uint64, kind string) error {
//	return repo.dao.UpdateStateByStudentId(StudentId, newState, kind)
//}

//func (repo *CachedRepository) NoteRecords(info []domain.RecordsInRepository, grade string, kind string) error {
//	switch kind {
//	case Card:
//		return repo.dao.NoteCardRecords(info, grade)
//	case Virtual:
//		return repo.dao.NoteVirtualRecords(info, grade)
//	default:
//		return errors.New("非法类型")
//	}
//}

func (repo *CachedRepository) GetRecordsByStudentId(ctx context.Context, StudentId string, kind string, grade string, time time.Time) ([]domain.Records, error) {
	switch kind {
	case Card:
		return repo.dao.GetCardRecordsByStudentId(ctx, StudentId, grade, time)
	case Virtual:
		return repo.dao.GetVirtualRecordsByStudentId(ctx, StudentId, grade, time)
	default:
		return nil, errors.New("非法类型")
	}
}

// 记录新数据并跟新用户状态
func (repo *CachedRepository) UpdateRecordsAndUserInfo(ctx context.Context, info []domain.RecordsInRepository, msg domain.ServiceMsg, user domain.User) error {
	return repo.dao.UpdateRecordsAndUserInfo(ctx, info, msg, user)
}
