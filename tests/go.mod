module github.com/sanhuanshisanshao/gen/tests

go 1.16

require (
	github.com/sanhuanshisanshao/gen v0.0.0-20230321034533-beaf66af161f
	golang.org/x/mod v0.6.0 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.0
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0
)

replace github.com/sanhuanshisanshao/gen => ../
