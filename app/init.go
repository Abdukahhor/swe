package app

import "github.com/abdukahhor/swe/storage"

//Core - main layer
type Core struct {
	db storage.DB
	//todo logging service
}

//New - инициализировать новый экземпляр Core
func New(s storage.DB) *Core {
	return &Core{db: s}
}
