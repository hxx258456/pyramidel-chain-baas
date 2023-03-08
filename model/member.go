package model

import "gorm.io/gorm"

type Member struct {
	Base
	StoreType      int          `json:"storeType" gorm:"column:storeType"`                                // 0：密钥托管，1：自行管理
	Name           string       `json:"name" gorm:"column:name;uniqueIndex:org_name" binding:"required"`  // ca用户名
	PassWord       string       `json:"passWord" gorm:"column:passWord" binding:"required"`               // ca密码
	UserType       string       `json:"userType" gorm:"column:userType" binding:"required"`               // client,admin
	OrganizationId uint         `json:"organizationId" gorm:"column:organizationId;uniqueIndex:org_name"` // 所属组织
	organization   Organization `json:"-" gorm:"foreignKey:OrganizationId"`
	IsFrozen       bool         `json:"IsFrozen" gorm:"column:IsFrozen"`     // 是否冻结默认为false
	Uscc           string       `json:"orgUscc" binding:"required" gorm:"-"` // 组织唯一标识
}

func (Member) TableName() string {
	return "baas_member"
}

func (m *Member) Create() error {
	tx := db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	if err := tx.Create(m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}