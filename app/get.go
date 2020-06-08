package app

import (
	"log"

	"github.com/abdukahhor/swe/models"
)

//Get - returns number
func (c Core) Get(id string) models.Reply {
	r := models.Success
	if id == "" {
		return models.ErrID
	}
	num, err := c.db.Get(id)
	if err != nil {
		log.Println(err)
		if c.db.IsNotFound(err) {
			return models.NotFound
		}
		return models.DBError
	}
	r.Num = num
	r.ID = id
	return r
}
