package app

import "log"

//Increment - Increases the current number by the size of the increment.
func (c Core) Increment(id string) Reply {
	r := Success
	if id == "" {
		return ErrID
	}
	err := c.db.Increment(id)
	if err != nil {
		log.Println(err)
		if c.db.IsNotFound(err) {
			return NotFound
		}
		return DBError
	}
	r.ID = id
	return r
}
