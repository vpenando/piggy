module main

go 1.15

replace (
	github.com/vpenando/piggy/localization => ./localization
	github.com/vpenando/piggy/piggy => ./piggy
	github.com/vpenando/piggy/routing => ./routing
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/vpenando/piggy/localization v0.0.0-00010101000000-000000000000
	github.com/vpenando/piggy/piggy v0.0.0-20201124232141-06292e00dada // indirect
	github.com/vpenando/piggy/routing v0.0.0-00010101000000-000000000000
	gopkg.in/ini.v1 v1.62.0
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.7
)
