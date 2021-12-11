package model

import (
	"time"
)

// SpaceIDentity [...]
type SpaceIDentity struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"-"`
	Code          string    `gorm:"unique;column:code;type:varchar(20);not null" json:"code"`                                 // 标识码
	SpaceAvatar   string    `gorm:"unique;column:space_avatar;type:varchar(70)" json:"spaceAvatar"`                           // 头像
	SpaceClothing string    `gorm:"unique;column:space_clothing;type:varchar(70)" json:"spaceClothing"`                       // 服装
	Sex           int      `gorm:"column:sex;type:tinyint(1);not null" json:"sex"`                                           // 性别（1男 2女）
	Sort          int       `gorm:"column:sort;type:int(11);not null;default:1" json:"sort"`                                  // 排序
	Modifier      int64     `gorm:"column:modifier;type:bigint(20);not null" json:"modifier"`                                 // 更新人
	Creator       int64     `gorm:"column:creator;type:bigint(20);not null" json:"creator"`                                   // 创建人
}

// TableName get sql table name.获取数据库表名
func (m *SpaceIDentity) TableName() string {
	return "space_identity"
}

// SpaceIDentityColumns get sql column name.获取数据库列名
var SpaceIDentityColumns = struct {
	ID            string
	Code          string
	SpaceAvatar   string
	SpaceClothing string
	Sex           string
	Sort          string
	Modifier      string
	Creator       string
	GmtModified   string
	GmtCreated    string
}{
	ID:            "id",
	Code:          "code",
	SpaceAvatar:   "space_avatar",
	SpaceClothing: "space_clothing",
	Sex:           "sex",
	Sort:          "sort",
	Modifier:      "modifier",
	Creator:       "creator",
	GmtModified:   "gmt_modified",
	GmtCreated:    "gmt_created",
}

// User [...]
type User struct {
	ID        int       `gorm:"primaryKey;column:ID;type:int(11);not null" json:"-"`                                // ID
	Name      string    `gorm:"column:Name;type:varchar(64)" json:"name"`                                           // 名称
	Account   string    `gorm:"column:Account;type:varchar(64);not null;default:''" json:"account"`                 // 账号
	Passwd    string    `gorm:"column:Passwd;type:varchar(64);default:''" json:"passwd"`                            // 密码
	Gender    int8      `gorm:"column:Gender;type:tinyint(4);default:0" json:"gender"`                              // 性别：0=女，1=男
	Tag       string    `gorm:"column:Tag;type:varchar(128);default:0" json:"tag"`                                  // 标签
	LastLogIn time.Time `gorm:"column:LastLogIn;type:datetime;default:CURRENT_TIMESTAMP" json:"lastLogIn"`          // 最近登录时间
	CreatUser int       `gorm:"column:CreatUser;type:int(11)" json:"creatUser"`                                     // 创建者ID
	CreatTime time.Time `gorm:"column:CreatTime;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creatTime"` // 创建时间
	UpdTime   time.Time `gorm:"column:UpdTime;type:datetime;default:CURRENT_TIMESTAMP" json:"updTime"`              // 更新者ID
	UpdUser   int       `gorm:"column:UpdUser;type:int(11)" json:"updUser"`                                         // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	Name      string
	Account   string
	Passwd    string
	Gender    string
	Tag       string
	LastLogIn string
	CreatUser string
	CreatTime string
	UpdTime   string
	UpdUser   string
}{
	ID:        "ID",
	Name:      "Name",
	Account:   "Account",
	Passwd:    "Passwd",
	Gender:    "Gender",
	Tag:       "Tag",
	LastLogIn: "LastLogIn",
	CreatUser: "CreatUser",
	CreatTime: "CreatTime",
	UpdTime:   "UpdTime",
	UpdUser:   "UpdUser",
}
