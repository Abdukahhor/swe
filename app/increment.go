package app

import (
	"log"

	"github.com/abdukahhor/swe/models"
)

//Increment - Increases the current number by the size of the increment.
func (c Core) Increment(id string) models.Reply {
	r := models.Success
	if id == "" {
		return models.ErrID
	}
	err := c.db.Increment(id)
	if err != nil {
		log.Println(err)
		if c.db.IsNotFound(err) {
			return models.NotFound
		}
		return models.DBError
	}
	r.ID = id
	return r
}
