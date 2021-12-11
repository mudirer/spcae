package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"space/pkg/app"
	"space/pkg/code"
	"space/pkg/logutil"
	"space/service/auth_service"
	"strconv"
	"time"
)


type SpaceIDentity struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id" binding:"max=20"`
	Code          string    `gorm:"unique;column:code;type:varchar(20);not null" json:"code" `                                 // 标识码
	SpaceAvatar   string    `gorm:"unique;column:space_avatar;type:varchar(70)" json:"spaceAvatar" binding:"required,max=70"`                           // 头像
	SpaceClothing string    `gorm:"unique;column:space_clothing;type:varchar(70)" json:"spaceClothing" binding:"required,max=70"`                       // 服装
	Sex           int      `gorm:"column:sex;type:tinyint(1);not null" json:"sex" binding:"required,min=1,max=2"`                                            // 性别（1男 2女）
	Sort          int       `gorm:"column:sort;type:int(11);not null;default:1" json:"sort" binding:"required"`                                  // 排序
	Modifier      int64     `gorm:"column:modifier;type:bigint(20);not null" json:"modifier"`                                 // 更新人
	Creator       int64     `gorm:"column:creator;type:bigint(20);not null" json:"creator"`                                   // 创建人
	GmtModified   time.Time `gorm:"column:gmt_modified;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"gmtModified"` // 更新时间
	GmtCreated    time.Time `gorm:"column:gmt_created;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"gmtCreated"`   // 创建时间
}




// @注册
// @Description Register
// @Accept  json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok","data": "data"}"
// @Failure 500 {string} json "{"code":1006, "msg": "添加账户失败", "data": "data"}"
// @Router /auth [post]
func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form SpaceIDentity
	)
	err:=c.ShouldBind(&form)
	if err!=nil{
		logutil.Log.Errorf("register shoubing %v err is %v",form,err)
		appG.Response(http.StatusMisdirectedRequest, code.PARAMETER_ERROR, nil)
		return
	}

	authSerivce := auth_service.SpaceIDentityColumns{
		Code: form.Code,
		SpaceAvatar: form.SpaceAvatar,
		SpaceClothing: form.SpaceClothing,
		Sex :form.Sex,
		Sort  :form.Sort,


	}
	result,err := authSerivce.Add()
	if err != nil {
		fmt.Println(err)
		logutil.Log.Errorf("register space error is %v",err)
		appG.Response(http.StatusInternalServerError, code.ERROR_POST_FAIL, nil)
		return
	}


	appG.Response(http.StatusOK, code.OK, result)
}

type UpdSpace struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id" binding:"max=20"`
	SpaceAvatar   string    `gorm:"unique;column:space_avatar;type:varchar(70)" json:"spaceAvatar" binding:"required,max=70"`                           // 头像
}

// @更新头像
// @Description UpdAvatar
// @Accept  json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok","data": "data"}"
// @Failure 500 {string} json "{"code":1006, "msg": "更新失败", "data": "data"}"
// @Router /updavatar [put]
func UpdAvatar(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UpdSpace
	)
	err:=c.ShouldBind(&form)
	if err!=nil{
		logutil.Log.Errorf("updavatar shoubing %v err is %v",form,err)
		appG.Response(http.StatusMisdirectedRequest, code.PARAMETER_ERROR, nil)
		return
	}

	updSerivce := auth_service.SpaceIDentityColumns{
		ID:form.ID,
		SpaceAvatar: form.SpaceAvatar,
	}
	res,err := updSerivce.UpdAva()
	if err != nil {
		fmt.Println(err)
		logutil.Log.Errorf("register space error is %v",err)
		appG.Response(http.StatusInternalServerError, code.ERROR_PUT_FAIL, nil)
		return
	}


	appG.Response(http.StatusOK, code.OK, res.ID)
}

type SpaceID struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id" binding:"max=20"`
}
// @删除
// @Description DeleteSpace
// @Accept  json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok","data": "data"}"
// @Failure 500 {string} json "{"code":1006, "msg": "删除失败", "data": "data"}"
// @Router /delete [delete]
func DeleteSpace(c *gin.Context) {
	var appG = app.Gin{C: c}
	id ,err:= strconv.ParseInt(c.Param("id"), 10, 64)
	if err!=nil{
		logutil.Log.Errorf("deletespcae  %v err is %v",c.Param("id"),err)
		appG.Response(http.StatusNotFound, code.PARAMETER_ERROR, nil)
		return
	}
	validate := validator.New()
	form := SpaceID{ID: id}
	err = validate.Struct(form)
	if err!=nil{
		logutil.Log.Errorf("getSpace shoubing %v err is %v",form,err)
		appG.Response(http.StatusMisdirectedRequest, code.PARAMETER_ERROR, nil)
		return
	}

	delSerivce := auth_service.SpaceIDentityColumns{
		ID:form.ID,
	}

	exists, err := delSerivce.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR_ID_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, code.ERROR_NOT_EXIST, nil)
		return
	}

	err = delSerivce.DelSpa()
	if err != nil {
		fmt.Println(err)
		logutil.Log.Errorf("delspacce space error is %v",err)
		appG.Response(http.StatusInternalServerError, code.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, code.OK, delSerivce.ID)
}


// @获取
// @Description GetSpace
// @Accept  json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok","data": "data"}"
// @Failure 500 {string} json "{"code":1006, "msg": "获取失败", "data": "data"}"
// @Router /get [get]
func GetSpace(c *gin.Context) {
	var appG = app.Gin{C: c}
	id ,err:= strconv.ParseInt(c.Param("id"), 10, 64)
	if err!=nil{
		logutil.Log.Errorf("getSpace  %v err is %v",c.Param("id"),err)
		appG.Response(http.StatusNotFound, code.PARAMETER_ERROR, nil)
		return
	}
	validate := validator.New()
	form := SpaceID{ID: id}
	err = validate.Struct(form)
	if err!=nil{
		logutil.Log.Errorf("getSpace shoubing %v err is %v",form,err)
		appG.Response(http.StatusMisdirectedRequest, code.PARAMETER_ERROR, nil)
		return
	}

	getSerivce := auth_service.SpaceIDentityColumns{
		ID:form.ID,
	}
	exists, err := getSerivce.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR_CHECK_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, code.ERROR_NOT_EXIST, nil)
		return
	}

	space,err := getSerivce.GetSpa()
	if err != nil {
		fmt.Println(err)
		logutil.Log.Errorf("delspacce space error is %v",err)
		appG.Response(http.StatusInternalServerError, code.ERROR_GET_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, code.OK, space)
}


type SpaceSex struct {
	Sex           int      `gorm:"column:sex;type:tinyint(1);not null" json:"sex" binding:"required,min=1,max=2"`                                            // 性别（1男 2女）
}
// @获取
// @Description GetList
// @Accept  json
// @Produce json
// @Success 200 {string} json "{"code":200,"msg":"ok","data": "data"}"
// @Failure 500 {string} json "{"code":1006, "msg": "获取失败", "data": "data"}"
// @Router /getlist [get]
func GetList(c *gin.Context) {
	var appG = app.Gin{C: c}
	id ,err:= strconv.Atoi(c.Param("sex"))
	if err!=nil{
		logutil.Log.Errorf("getSpacelist  %v err is %v",c.Param("sex"),err)
		appG.Response(http.StatusNotFound, code.PARAMETER_ERROR, nil)
		return
	}
	validate := validator.New()
	form := SpaceSex{Sex: id}
	err = validate.Struct(form)
	if err!=nil{
		logutil.Log.Errorf("getSpace shoubing %v err is %v",form,err)
		appG.Response(http.StatusMisdirectedRequest, code.PARAMETER_ERROR, nil)
		return
	}
	getSerivce := auth_service.SpaceIDentityColumns{
		Sex:form.Sex,
	}

	space,err := getSerivce.GetList()
	if err != nil {
		fmt.Println(err)
		logutil.Log.Errorf("get space error is %v",err)
		appG.Response(http.StatusInternalServerError, code.ERROR_GET_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, code.OK, space)
}


// @测试接口
// @Produce json
// @Success 200 {string} json "{ "name": "1","password":"2"}"
// @Router /test [get]
func Get1(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "1",
		"password":"2",

	})

}

