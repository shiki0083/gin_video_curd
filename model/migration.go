package model

//执行数据迁移 迁移到SQL数据库

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
}
