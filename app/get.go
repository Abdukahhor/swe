package app

//Get -
func (c Core) Get(id string) Reply {
	r := Success
	num, err := c.db.Get(id)
	if err != nil {
		if c.db.IsNotFound(err) {
			return NotFound
		}
		return DBError
	}
	r.Num = num
	return r
}
