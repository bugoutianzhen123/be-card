package service

import (
	"context"
	v1 "github.com/asynccnu/be-api/gen/proto/card/v1"
	"github.com/asynccnu/be-card/domain"
	"github.com/asynccnu/be-card/repository"
	"log"
)

type Service interface {
	CreateUser(ctx context.Context, msg domain.ServiceMsg) error
	UpdateUserKey(ctx context.Context, msg domain.ServiceMsg) error
	GetRecordOfConsumption(ctx context.Context, msg domain.ServiceMsg) ([]*v1.RecordOfConsumption, error)
}

type cardService struct {
	repo repository.CardRepository
}

func NewCardService(repo repository.CardRepository) Service {
	return &cardService{repo: repo}
}

// 记录用户信息
func (ser *cardService) CreateUser(ctx context.Context, msg domain.ServiceMsg) error {
	return ser.repo.NoteKey(ctx, msg.Key, msg.StudentId)
}

// 更新用户key
func (ser *cardService) UpdateUserKey(ctx context.Context, msg domain.ServiceMsg) error {
	return ser.repo.UpdateKeyByStudentId(ctx, msg.StudentId, msg.Key)
}

// 获取消费记录
func (ser *cardService) GetRecordOfConsumption(ctx context.Context, msg domain.ServiceMsg) ([]*v1.RecordOfConsumption, error) {
	if err := msg.GetMsg(); err != nil {
		return nil, err
	}

	user, err := ser.repo.GetUserInfoByStudentId(ctx, msg.StudentId)
	if err != nil {
		return nil, err
	}

	//检查key
	if msg.Key == "" {
		//key 为空，从数据库获取
		k, err := ser.repo.GetKeyByStudentId(ctx, msg.StudentId)
		if err != nil {
			return nil, err
		}
		msg.Key = k
	}
	//判断查询月份是否是本月
	if CheckMonth(msg.StartTime) {
		//是本月，调用爬虫
		response, err := GetRecordOfConsumptionByVCard(msg)
		if err != nil {
			return nil, err
		}
		return RecordsOfConsumptionToResponse(response)
	} else {
		//不是本月，查看目标月份是否在数据库
		if CheckState(user, msg) {
			//在，提取数据返回
			response, err := ser.repo.GetRecordsByStudentId(ctx, msg.StudentId, msg.Type, msg.Grade, msg.Stime)
			if err != nil {
				return nil, err
			}
			return RecordsToResponse(response)
		} else {
			//不在，调用爬虫，并异步存入库
			response, err := GetRecordOfConsumptionByVCard(msg)
			if err != nil {
				return nil, err
			}
			//异步插入数据库
			go func() {
				var records []domain.RecordsInRepository
				//如果出错则重新尝试
				//for i := 0; i < 5; i++ {
				//数据转换
				records, err = GetCardRecordsForRepository(response, msg.StudentId)
				if err != nil {
					log.Println(err)
					//continue
				}
				//}
				//for i := 0; i < 5; i++ {
				//记录数据并更新用户数据
				err = ser.repo.UpdateRecordsAndUserInfo(ctx, records, msg, user)
				if err != nil {
					log.Println(err)
					//continue
				}
				//break
				//}

			}()
			return RecordsOfConsumptionToResponse(response)
		}
	}
}
