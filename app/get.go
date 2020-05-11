package app

import "log"

//Get -
func (c Core) Get(id string) Reply {
	r := Success
	if id == "" {
		return ErrID
	}
	num, err := c.db.Get(id)
	if err != nil {
		log.Println(err)
		if c.db.IsNotFound(err) {
			return NotFound
		}
		return DBError
	}
	r.Num = num
	r.ID = id
	return r
}
