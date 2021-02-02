module main

go 1.15

replace (
	github.com/vpenando/piggy/piggy => ./piggy
	localization => ./localization
	routing => ./routing
)

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/gorilla/mux v1.8.0
	github.com/vpenando/piggy/piggy v0.0.0-20201124232141-06292e00dada
	gopkg.in/ini.v1 v1.62.0
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.7
	localization v0.0.0-00010101000000-000000000000
	routing v0.0.0-00010101000000-000000000000
)
