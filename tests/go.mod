module github.com/sanhuanshisanshao/gen/tests

go 1.16

require (
	golang.org/x/mod v0.6.0 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/sqlite v1.4.3
	github.com/sanhuanshisanshao/gen v0.3.16
	gorm.io/gorm v1.24.0
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0
)

replace github.com/sanhuanshisanshao/gen => ../
