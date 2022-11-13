// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package monitor

import (
	"time"
)

const TableNameSysLogininfor = "sys_logininfor"

// SysLogininfor mapped from table <sys_logininfor>
type SysLogininfor struct {
	InfoID        int64      `gorm:"column:info_id;type:bigint;primaryKey;autoIncrement:true" json:"infoId"` // 访问ID
	Username      *string    `gorm:"column:username;type:varchar(50)" json:"username"`                       // 用户账号
	Ipaddr        *string    `gorm:"column:ipaddr;type:varchar(128)" json:"ipaddr"`                          // 登录IP地址
	LoginLocation *string    `gorm:"column:login_location;type:varchar(255)" json:"loginLocation"`           // 登录地点
	Browser       *string    `gorm:"column:browser;type:varchar(50)" json:"browser"`                         // 浏览器类型
	Os            *string    `gorm:"column:os;type:varchar(50)" json:"os"`                                   // 操作系统
	Status        *string    `gorm:"column:status;type:char(1);default:0" json:"status"`                     // 登录状态（0成功 1失败）
	Msg           *string    `gorm:"column:msg;type:varchar(255)" json:"msg"`                                // 提示消息
	LoginTime     *time.Time `gorm:"column:login_time;type:datetime" json:"loginTime"`                       // 访问时间
}

// TableName SysLogininfor's table name
func (*SysLogininfor) TableName() string {
	return TableNameSysLogininfor
}
