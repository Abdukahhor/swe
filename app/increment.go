package app

//Increment -
func (c Core) Increment(id string) Reply {
	r := Success
	err := c.db.Increment(id)
	if err != nil {
		if c.db.IsNotFound(err) {
			return NotFound
		}
		return DBError
	}
	return r
}
