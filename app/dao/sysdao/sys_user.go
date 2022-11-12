package sysdao

import (
	"context"
	"go-vea/app/common/e"
	"go-vea/app/common/page"
	"go-vea/app/model/system"
	"go-vea/app/model/system/request"
	"go-vea/configs"
	"gorm.io/gorm"
)

type SysUserDao struct {
	*gorm.DB
}

func NewSysUserDao(ctx context.Context) *SysUserDao {
	return &SysUserDao{configs.GetDB(ctx)}
}

func NewSysUserDaoByDB(db *gorm.DB) *SysUserDao {
	return &SysUserDao{db}
}

func (dao *SysUserDao) SelectList(sysUser *request.SysUser) (p *page.Pagination, err error) {
	var userLIst []*system.SysUser
	p = new(page.Pagination)

	dao.DB = dao.DB.Table("sys_user u").
		Select("u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber," +
			" u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark," +
			" d.dept_name, d.leader").
		Joins("left join sys_dept d on u.dept_id = d.dept_id").
		Where("u.del_flag = '0'")
	if sysUser.UserName != "" {
		dao.DB = dao.DB.Where("u.user_name = ?", sysUser.UserName)
	}
	if sysUser.Status != "" {
		dao.DB = dao.DB.Where("u.status = ?", sysUser.Status)
	}
	if sysUser.Phonenumber != "" {
		dao.DB = dao.DB.Where("u.phonenumber = ?", sysUser.Phonenumber)
	}
	if sysUser.DeptID != 0 {
		dao.DB = dao.DB.Where("u.dept_id = ?", sysUser.DeptID).
			Or("u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(?, ancestors) )", sysUser.DeptID)
	}
	if sysUser.DataScope != "" {
		// todo 数据范围过滤
	}

	if sysUser.OpenPage {
		p.PageNum = sysUser.PageNum
		p.PageSize = sysUser.PageSize
		err = dao.DB.Scopes(page.SelectPage(userLIst, p, dao.DB)).Find(&userLIst).Error
	} else {
		err = dao.DB.Find(&userLIst).Error
	}
	p.Rows = userLIst
	if err != nil {
		p.Code = e.ERROR
		p.Msg = err.Error()
		return p, err
	}
	return p, err
}

func (dao *SysUserDao) SelectAll(sysUser *request.SysUser) (list []system.SysUser, err error) {
	if sysUser.UserID != 0 {
		dao.DB = dao.DB.Where("user_id = ?", sysUser.UserID)
	}
	if sysUser.UserName != "" {
		dao.DB = dao.DB.Where("user_name = ?", sysUser.UserName)
	}
	if sysUser.Status != "" {
		dao.DB = dao.DB.Where("status = ?", sysUser.Status)
	}
	if sysUser.DataScope != "" {
		// todo
	}

	err = dao.DB.Where("del_flag = '0'").Find(&list).Error
	return
}

func (dao *SysUserDao) SelectUserByUserName(username string) (sysUser *system.SysUser, err error) {
	err = dao.DB.Model(&system.SysUser{}).Where("user_name=?", username).First(&sysUser).Error
	return
}

func (dao *SysUserDao) SelectById(id int64) (sysUser *system.SysUser, err error) {
	err = dao.DB.Model(&sysUser).Where("user_id=?", id).Find(&sysUser).Error
	return
}

func (dao *SysUserDao) Insert(sysUser *system.SysUser) error {
	return dao.DB.Model(&system.SysUser{}).Create(sysUser).Error
}

func (dao *SysUserDao) UpdateById(sysUser *system.SysUser) error {
	// Updates 根据 `struct` 更新属性，只会更新非零值的字段
	return dao.DB.Updates(sysUser).Error
}

func (dao *SysUserDao) DeleteById(id int64) error {
	return dao.DB.Where("user_id = ?", id).Delete(&system.SysUser{}).Error
}

func (dao *SysUserDao) DeleteByIds(ids []int64) error {
	return dao.DB.Where("user_id in (?)", ids).Delete(&system.SysUser{}).Error
}

func (dao *SysUserDao) CheckUserNameUnique(username string) (sysUser *system.SysUser, err error) {
	// select user_id, user_name from sys_user where user_name = #{userName} and del_flag = '0' limit 1
	err = dao.DB.Select("user_id, user_name").Where("user_name = ? and del_flag = '0'", username).First(&sysUser).Error
	return sysUser, err
}

func (dao *SysUserDao) CheckPhoneUnique(phoneNumber string) (sysUser *system.SysUser, err error) {
	// select user_id, email from sys_user where email = #{email} and del_flag = '0' limit 1
	err = dao.DB.Select("user_id, phonenumber").Where("phonenumber = ? and del_flag = '0'", phoneNumber).First(&sysUser).Error
	return sysUser, err
}

func (dao *SysUserDao) CheckEmailUnique(email string) (sysUser *system.SysUser, err error) {
	// select user_id, email from sys_user where email = #{email} and del_flag = '0' limit 1
	err = dao.DB.Select("user_id, email").Where("email = ? and del_flag = '0'", email).First(&sysUser).Error
	return sysUser, err
}
