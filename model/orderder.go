package model

import (
	"github.com/hxx258456/pyramidel-chain-baas/pkg/utils/localtime"
	"gorm.io/gorm"
)

type Orderer struct {
	Domain         string               `json:"Domain" gorm:"column:Domain;unique"`    // 域名 ${uscc}.example.com
	DueTime        *localtime.LocalTime `json:"dueTime" gorm:"column:dueTime"`         // 到期时间
	RestartTime    *localtime.LocalTime `json:"restartTime" gorm:"column:restartTime"` // 重启时间
	NodeCore       uint                 `json:"nodeCore" gorm:"column:nodeCore"`
	NodeMemory     uint                 `json:"nodeMemory" gorm:"column:nodeMemory"`
	NodeBandwidth  uint                 `json:"nodeBandwidth" gorm:"column:nodeBandwidth"`
	NodeDisk       uint                 `json:"nodeDisk" gorm:"column:nodeDisk"`
	HostId         uint                 `json:"hostId" gorm:"column:hostId"` // 所在主机
	Host           Host                 `json:"-" gorm:"foreignKey:HostId"`
	SerialNumber   uint                 `json:"serialNumber" gorm:"column:serialNumber"`     // 序列号
	Port           uint                 `json:"port" gorm:"column:port"`                     // 占用端口号
	Name           string               `json:"name" gorm:"column:name"`                     // 节点名ex: orderer1
	OrganizationId uint                 `json:"organizationId" gorm:"column:organizationId"` // 所属组织
	Organization   Organization         `json:"-" gorm:"foreignKey:OrganizationId" `
	OrgPackageId   uint64               `json:"orgPackageId" gorm:"column:orgPackageId"` // 订单id
	Status         int                  `json:"status" gorm:"column:status"`             // 状态
	Error          string               `json:"error" gorm:"column:_"`                   // 节点当前错误
	NodeType       int                  `json:"nodeType" gorm:"column:_"`                // 1代表orderer

	Base
}

func (Orderer) TableName() string {
	return "baas_orderer"
}

func (o *Orderer) Create(tx *gorm.DB) error {
	if tx == nil {
		tx = db.Session(&gorm.Session{
			SkipDefaultTransaction: true,
		})
	}
	if err := tx.Create(o).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (o *Orderer) GetMaxSerial(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db.Session(&gorm.Session{
			SkipDefaultTransaction: true,
		})
	}
	if err := tx.Where("organizationId = ?", id).Order("serialNumber DESC").First(&o).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			o.ID = 0
			return nil
		} else {
			return err
		}
	}
	return nil
}
