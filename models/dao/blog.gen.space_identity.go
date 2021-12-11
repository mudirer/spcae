package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SpaceIDentityMgr struct {
	*_BaseMgr
}

// SpaceIDentityMgr open func
func SpaceIDentityMgr(db *gorm.DB) *_SpaceIDentityMgr {
	if db == nil {
		panic(fmt.Errorf("SpaceIDentityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpaceIDentityMgr{_BaseMgr: &_BaseMgr{DB: db.Model(SpaceIDentity{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SpaceIDentityMgr) GetTableName() string {
	return "space_identity"
}

// Get 获取
func (obj *_SpaceIDentityMgr) Get() (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SpaceIDentityMgr) Gets() (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SpaceIDentityMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 标识码
func (obj *_SpaceIDentityMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithSpaceAvatar space_avatar获取 头像
func (obj *_SpaceIDentityMgr) WithSpaceAvatar(spaceAvatar string) Option {
	return optionFunc(func(o *options) { o.query["space_avatar"] = spaceAvatar })
}

// WithSpaceClothing space_clothing获取 服装
func (obj *_SpaceIDentityMgr) WithSpaceClothing(spaceClothing string) Option {
	return optionFunc(func(o *options) { o.query["space_clothing"] = spaceClothing })
}

// WithSex sex获取 性别（1男 2女）
func (obj *_SpaceIDentityMgr) WithSex(sex bool) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithSort sort获取 排序
func (obj *_SpaceIDentityMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithModifier modifier获取 更新人
func (obj *_SpaceIDentityMgr) WithModifier(modifier int64) Option {
	return optionFunc(func(o *options) { o.query["modifier"] = modifier })
}

// WithCreator creator获取 创建人
func (obj *_SpaceIDentityMgr) WithCreator(creator int64) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithGmtModified gmt_modified获取 更新时间
func (obj *_SpaceIDentityMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithGmtCreated gmt_created获取 创建时间
func (obj *_SpaceIDentityMgr) WithGmtCreated(gmtCreated time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_created"] = gmtCreated })
}

// GetByOption 功能选项模式获取
func (obj *_SpaceIDentityMgr) GetByOption(opts ...Option) (result SpaceIDentity, err error) {
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
func (obj *_SpaceIDentityMgr) GetByOptions(opts ...Option) (results []*SpaceIDentity, err error) {
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

// GetFromID 通过id获取内容
func (obj *_SpaceIDentityMgr) GetFromID(id int64) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SpaceIDentityMgr) GetBatchFromID(ids []int64) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 标识码
func (obj *_SpaceIDentityMgr) GetFromCode(code string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error

	return
}

// GetBatchFromCode 批量查找 标识码
func (obj *_SpaceIDentityMgr) GetBatchFromCode(codes []string) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromSpaceAvatar 通过space_avatar获取内容 头像
func (obj *_SpaceIDentityMgr) GetFromSpaceAvatar(spaceAvatar string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_avatar` = ?", spaceAvatar).Find(&result).Error

	return
}

// GetBatchFromSpaceAvatar 批量查找 头像
func (obj *_SpaceIDentityMgr) GetBatchFromSpaceAvatar(spaceAvatars []string) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_avatar` IN (?)", spaceAvatars).Find(&results).Error

	return
}

// GetFromSpaceClothing 通过space_clothing获取内容 服装
func (obj *_SpaceIDentityMgr) GetFromSpaceClothing(spaceClothing string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_clothing` = ?", spaceClothing).Find(&result).Error

	return
}

// GetBatchFromSpaceClothing 批量查找 服装
func (obj *_SpaceIDentityMgr) GetBatchFromSpaceClothing(spaceClothings []string) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_clothing` IN (?)", spaceClothings).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别（1男 2女）
func (obj *_SpaceIDentityMgr) GetFromSex(sex bool) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 性别（1男 2女）
func (obj *_SpaceIDentityMgr) GetBatchFromSex(sexs []bool) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SpaceIDentityMgr) GetFromSort(sort int) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SpaceIDentityMgr) GetBatchFromSort(sorts []int) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromModifier 通过modifier获取内容 更新人
func (obj *_SpaceIDentityMgr) GetFromModifier(modifier int64) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modifier` = ?", modifier).Find(&results).Error

	return
}

// GetBatchFromModifier 批量查找 更新人
func (obj *_SpaceIDentityMgr) GetBatchFromModifier(modifiers []int64) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modifier` IN (?)", modifiers).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 创建人
func (obj *_SpaceIDentityMgr) GetFromCreator(creator int64) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 创建人
func (obj *_SpaceIDentityMgr) GetBatchFromCreator(creators []int64) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容 更新时间
func (obj *_SpaceIDentityMgr) GetFromGmtModified(gmtModified time.Time) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找 更新时间
func (obj *_SpaceIDentityMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromGmtCreated 通过gmt_created获取内容 创建时间
func (obj *_SpaceIDentityMgr) GetFromGmtCreated(gmtCreated time.Time) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gmt_created` = ?", gmtCreated).Find(&results).Error

	return
}

// GetBatchFromGmtCreated 批量查找 创建时间
func (obj *_SpaceIDentityMgr) GetBatchFromGmtCreated(gmtCreateds []time.Time) (results []*SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`gmt_created` IN (?)", gmtCreateds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SpaceIDentityMgr) FetchByPrimaryKey(id int64) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_SpaceIDentityMgr) FetchUniqueByCode(code string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&result).Error

	return
}

// FetchUniqueBySpaceAvatar primary or index 获取唯一内容
func (obj *_SpaceIDentityMgr) FetchUniqueBySpaceAvatar(spaceAvatar string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_avatar` = ?", spaceAvatar).Find(&result).Error

	return
}

// FetchUniqueBySpaceClothing primary or index 获取唯一内容
func (obj *_SpaceIDentityMgr) FetchUniqueBySpaceClothing(spaceClothing string) (result SpaceIDentity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`space_clothing` = ?", spaceClothing).Find(&result).Error

	return
}
