package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Model(User{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取 ID
func (obj *_UserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取 名称
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithAccount Account获取 账号
func (obj *_UserMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["Account"] = account })
}

// WithPasswd Passwd获取 密码
func (obj *_UserMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["Passwd"] = passwd })
}

// WithGender Gender获取 性别：0=女，1=男
func (obj *_UserMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["Gender"] = gender })
}

// WithTag Tag获取 标签
func (obj *_UserMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["Tag"] = tag })
}

// WithLastLogIn LastLogIn获取 最近登录时间
func (obj *_UserMgr) WithLastLogIn(lastLogIn time.Time) Option {
	return optionFunc(func(o *options) { o.query["LastLogIn"] = lastLogIn })
}

// WithCreatUser CreatUser获取 创建者ID
func (obj *_UserMgr) WithCreatUser(creatUser int) Option {
	return optionFunc(func(o *options) { o.query["CreatUser"] = creatUser })
}

// WithCreatTime CreatTime获取 创建时间
func (obj *_UserMgr) WithCreatTime(creatTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatTime"] = creatTime })
}

// WithUpdTime UpdTime获取 更新者ID
func (obj *_UserMgr) WithUpdTime(updTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["UpdTime"] = updTime })
}

// WithUpdUser UpdUser获取 更新时间
func (obj *_UserMgr) WithUpdUser(updUser int) Option {
	return optionFunc(func(o *options) { o.query["UpdUser"] = updUser })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过ID获取内容 ID
func (obj *_UserMgr) GetFromID(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ID` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_UserMgr) GetBatchFromID(ids []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过Name获取内容 名称
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAccount 通过Account获取内容 账号
func (obj *_UserMgr) GetFromAccount(account string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 账号
func (obj *_UserMgr) GetBatchFromAccount(accounts []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPasswd 通过Passwd获取内容 密码
func (obj *_UserMgr) GetFromPasswd(passwd string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找 密码
func (obj *_UserMgr) GetBatchFromPasswd(passwds []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromGender 通过Gender获取内容 性别：0=女，1=男
func (obj *_UserMgr) GetFromGender(gender int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Gender` = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找 性别：0=女，1=男
func (obj *_UserMgr) GetBatchFromGender(genders []int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Gender` IN (?)", genders).Find(&results).Error

	return
}

// GetFromTag 通过Tag获取内容 标签
func (obj *_UserMgr) GetFromTag(tag string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 标签
func (obj *_UserMgr) GetBatchFromTag(tags []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromLastLogIn 通过LastLogIn获取内容 最近登录时间
func (obj *_UserMgr) GetFromLastLogIn(lastLogIn time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`LastLogIn` = ?", lastLogIn).Find(&results).Error

	return
}

// GetBatchFromLastLogIn 批量查找 最近登录时间
func (obj *_UserMgr) GetBatchFromLastLogIn(lastLogIns []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`LastLogIn` IN (?)", lastLogIns).Find(&results).Error

	return
}

// GetFromCreatUser 通过CreatUser获取内容 创建者ID
func (obj *_UserMgr) GetFromCreatUser(creatUser int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`CreatUser` = ?", creatUser).Find(&results).Error

	return
}

// GetBatchFromCreatUser 批量查找 创建者ID
func (obj *_UserMgr) GetBatchFromCreatUser(creatUsers []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`CreatUser` IN (?)", creatUsers).Find(&results).Error

	return
}

// GetFromCreatTime 通过CreatTime获取内容 创建时间
func (obj *_UserMgr) GetFromCreatTime(creatTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`CreatTime` = ?", creatTime).Find(&results).Error

	return
}

// GetBatchFromCreatTime 批量查找 创建时间
func (obj *_UserMgr) GetBatchFromCreatTime(creatTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`CreatTime` IN (?)", creatTimes).Find(&results).Error

	return
}

// GetFromUpdTime 通过UpdTime获取内容 更新者ID
func (obj *_UserMgr) GetFromUpdTime(updTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`UpdTime` = ?", updTime).Find(&results).Error

	return
}

// GetBatchFromUpdTime 批量查找 更新者ID
func (obj *_UserMgr) GetBatchFromUpdTime(updTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`UpdTime` IN (?)", updTimes).Find(&results).Error

	return
}

// GetFromUpdUser 通过UpdUser获取内容 更新时间
func (obj *_UserMgr) GetFromUpdUser(updUser int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`UpdUser` = ?", updUser).Find(&results).Error

	return
}

// GetBatchFromUpdUser 批量查找 更新时间
func (obj *_UserMgr) GetBatchFromUpdUser(updUsers []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`UpdUser` IN (?)", updUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`ID` = ?", id).Find(&result).Error

	return
}
