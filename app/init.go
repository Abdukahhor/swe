package app

import "github.com/abdukahhor/swe/storage"

//Core - main layer
type Core struct {
	db storage.DB
}

//New instance of Core
func New(s storage.DB) *Core {
	return &Core{db: s}
}

//Reply -
type Reply struct {
	Code    int
	Message string
	Num     uint64
	ID      string
}

// Replies list
var (
	Success  = Reply{Code: 1, Message: "Успешно"}
	NotFound = Reply{Code: 2, Message: "Инкремент не найден"}
	DBError  = Reply{Code: 3, Message: "Ошибка базы данных"}
	ErrSize  = Reply{Code: 4, Message: "Размер инкремента должен быть положительным"}
	ErrMax   = Reply{Code: 5, Message: "Верхней границы инкремента должен быть положительным"}
)
