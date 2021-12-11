package models

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"space/models/dao"
	"space/pkg/crypto"
	"space/pkg/logutil"
	"space/pkg/suid"
)


const Salt="LiHS"

func Register(spaceAvatar,spaceClothing string ,sex,sort int)  (*model.SpaceIDentity,error) {

	uuid,err:=suid.New()
	if err!=nil{
		logutil.Log.Errorf("suid error is %v",err)
		return nil,err
	}
	user := model.SpaceIDentity{
		Code :uuid,
		SpaceAvatar:spaceAvatar,
		SpaceClothing :spaceClothing,
		Sex   :sex,
		Sort  :sort,
	}
	result := db.Create(&user)
	if result.Error!=nil{
		return nil,result.Error
	}
	res,err:=UpdModCre(&user)
	if err!=nil{
		return nil,err
	}

	return res,nil
}


func UpdAvatar(id int64,spaceAvatar string )  (*model.SpaceIDentity,error) {

	user := model.SpaceIDentity{
		ID: id,
		SpaceAvatar:spaceAvatar,
	}
	result:=db.Model(&user).Updates(model.SpaceIDentity{SpaceAvatar:user.SpaceAvatar})
	if result.Error!=nil{
		logutil.Log.Errorf("updmodify  error is %v",result.Error)
		return nil,result.Error
	}
	return &user,nil
}

func UpdModCre(user *model.SpaceIDentity) (*model.SpaceIDentity,error) {
	result:=db.Model(&user).Updates(map[string]interface{}{"modifier": user.ID, "creator":user.ID})
	if result.Error!=nil{
		logutil.Log.Errorf("updmodify  error is %v",result.Error)
		return nil,result.Error
	}
	return user,nil
}

func DelSpace(id int64 )  error {

	user := model.SpaceIDentity{
		ID: id,
	}
	result:=db.Delete(&user)
	if result.Error!=nil{
		logutil.Log.Errorf("delspace  error is %v",result.Error)
		return result.Error
	}
	return nil
}
func ExistByID(id int64 )  (bool,error) {

	var user model.SpaceIDentity
	err := db.Select("id").Where("id = ?  ", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}


func GetSpace(id int64 )  (*model.SpaceIDentity,error) {

	user := model.SpaceIDentity{
		ID: id,
	}
	result:=db.First(&user)
	if result.Error!=nil&&result.Error != gorm.ErrRecordNotFound{
		logutil.Log.Errorf("getspace  error is %v",result.Error)
		return nil,result.Error
	}
	return &user,nil
}

func GetSpaceList(sex int )  ([]*model.SpaceIDentity,error) {

	var user = []*model.SpaceIDentity{}
	result:=db.Where("sex = ?", sex).Find(&user)
	if result.Error!=nil {
		logutil.Log.Errorf("getspace  error is %v",result.Error)
		return nil,result.Error
	}
	return user,nil
}





func encryptpwd(pwd string) string {
	var pushbuf bytes.Buffer
	pushbuf.WriteString(Salt)
	pushbuf.WriteString(pwd)
	encryptPassword := crypto.Sha256Hex(pushbuf.String())
	return encryptPassword
}