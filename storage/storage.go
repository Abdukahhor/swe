package storage

// DB is interface for database
type DB interface {
	//Gets number from storage by id
	Get(id string) (num uint64, err error)
	//increments number by id
	Increment(id string) (err error)
	//inserts new increment settings,
	//size of increment,
	//max number of increment
	//return id of increment, id is unique and random generated
	Setting(size uint64, max uint64) (id string, err error)
	//updates increment settings
	UpdateSetting(id string, size uint64, max uint64) (err error)
	//
	IsNotFound(err error) bool
	//Close database
	Close() error
}
