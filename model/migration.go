package model

//执行初始化

func migration() {
	// GORM 的AutoMigrate函数，仅支持建表，不支持修改字段和删除字段，避免意外导致丢失数据。
	// 根据User结构体，自动创建表结构.
	_ = DB.AutoMigrate(&User{}, &Host{})
}
