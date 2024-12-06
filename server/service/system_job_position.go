package service

import (
	"stellar/common"
	"stellar/model"
)

func GetJobPositionListService() (jobPositions []model.SystemJobPosition, err error) {
	err = common.MySQLDB.Find(&jobPositions).Error
	return
}
