package model

// Host 用户模型
type Host struct {
	// gorm.Model
	HostId      int64  `json:"host_id" gorm:"primary_key"`
	SN          string `json:"sn" gorm:"size:128"`
	HostName    string `json:"host_name" gorm:"size:128"`
	Region      string `json:"region" gorm:"size:128"`
	Data_Center string `json:"data_center" gorm:"size:128"`
}

// GetHost 用ID获取用户
func GetHost(ID interface{}) (Host, error) {
	var host Host
	result := DB.First(&host, ID)
	return host, result.Error
}

// 创建host
func (host *Host) AddHost() error {
	// 创建用户
	if err := DB.Create(&host).Error; err != nil {
		return err
	}
	return nil
}
