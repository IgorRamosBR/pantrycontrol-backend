package configuration

import (
	"github.com/go-pg/pg/v10"
)

func CreateDatabase(url string) *pg.DB {
	opt, err := pg.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return pg.Connect(opt)
}