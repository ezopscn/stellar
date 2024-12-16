package service

import (
	"stellar/common"
	"stellar/model"
)

func GetSystemJobPositionListService() (jobPositions []model.SystemJobPosition, err error) {
	err = common.MySQLDB.Find(&jobPositions).Error
	return
}
