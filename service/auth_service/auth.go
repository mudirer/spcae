package auth_service

import (
	"space/models"
	model "space/models/dao"
)


type SpaceIDentityColumns struct {
	ID            int64
	Code          string
	SpaceAvatar   string
	SpaceClothing string
	Sex           int
	Sort          int
	Modifier      int64
	Creator       int64
}


func (t *SpaceIDentityColumns) Add() (*model.SpaceIDentity,error) {
	return models.Register(t.SpaceAvatar,t.SpaceClothing,t.Sex,t.Sort)
}

func (t *SpaceIDentityColumns) UpdAva() (*model.SpaceIDentity,error) {
	return models.UpdAvatar(t.ID,t.SpaceAvatar)
}
func (t *SpaceIDentityColumns) DelSpa() error {
	return models.DelSpace(t.ID)
}
func (t *SpaceIDentityColumns) GetSpa() (*model.SpaceIDentity,error) {
	return models.GetSpace(t.ID)
}
func (t *SpaceIDentityColumns) ExistByID() (bool,error) {
	return models.ExistByID(t.ID)
}
func (t *SpaceIDentityColumns) GetList() ([]*model.SpaceIDentity,error) {
	return models.GetSpaceList(t.Sex)
}