package dal

import (
	"github.com/JackTJC/backend/dal/cache"
	"github.com/JackTJC/backend/dal/db"
)

func InitDal() {
	db.Init()
	cache.Init()
}
