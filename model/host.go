package model

import (
	"github.com/hxx258456/pyramidel-chain-baas/pkg/loadBalancer"
	"github.com/hxx258456/pyramidel-chain-baas/pkg/psutil/check"
	"gorm.io/gorm"
)

type Host struct {
	Base
	IntranetIp string `gorm:"column:intranet_ip;uniqueIndex:in_pub_ip;" json:"intranetIp" binding:"required"` // 内网ip地址
	PublicIp   string `gorm:"column:public_ip;uniqueIndex:in_pub_ip" json:"publicIp" binding:"required"`      // 公网ip地址
	Pw         string `gorm:"column:pw" json:"pw" binding:"required"`                                         // root用户密码
	Username   string `gorm:"column:username" json:"username" binding:"required"`                             // 用户名
	SSHPort    uint   `gorm:"column:sshPort" json:"sshPort" binding:"required"`                               // ssh port 为空时默认使用22端口
	Status     uint   `gorm:"column:status" json:"status"`                                                    // 是否开放使用
	UseIp      string `gorm:"column:use_ip" json:"useIp"`                                                     // 使用的Ip
	InstanceId string `gorm:"column:instanceId;unique" json:"instanceId" binding:"required"`                  // 实例id

	check.HostInfo
}

func (h *Host) Create() error {
	tx := db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	if err := tx.Create(h).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (h *Host) Update(val Host) error {
	tx := db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	tx.Model(h).Updates(val)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

func (h *Host) QueryAll(result interface{}) error {
	if db.Find(result).Error != nil {
		return db.Error
	}
	return nil
}

func (h *Host) QueryById(id uint, result interface{}) error {
	if err := db.Where("id = ?", id).Find(result).Error; err != nil {
		return err
	}
	return nil
}

func (Host) TableName() string {
	return "baas_host"
}

func (h *Host) InitHostLB() (*loadBalancer.LoadBalancer, error) {
	var result []Host
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	lb := &loadBalancer.LoadBalancer{}
	var services []loadBalancer.Service
	for _, v := range result {
		services = append(services, loadBalancer.Service{
			ID:             v.ID,
			Weight:         2,
			CurrentConns:   2,
			MaxConnections: 10,
		})
	}
	lb.Services = services
	return lb, nil
}
