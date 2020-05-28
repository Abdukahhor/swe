package app

//Reply - response structure for Core
type Reply struct {
	Code    int32
	Message string
	Num     uint64
	ID      string
}

// Replies list
var (
	Success  = Reply{Code: 1, Message: "Successfully"}
	NotFound = Reply{Code: 2, Message: "ID not found"}
	DBError  = Reply{Code: 3, Message: "Database Error"}
	ErrSize  = Reply{Code: 4, Message: "Increment size must be positive."}
	ErrMax   = Reply{Code: 5, Message: "The max number of the increment must be positive"}
	ErrID    = Reply{Code: 6, Message: "Invalid ID"}
)
