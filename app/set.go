package app

//Set - increment settings
//if id is empty sets new increment settings, else updates
func (c Core) Set(id string, size, max uint64) Reply {
	var (
		r   = Success
		err error
	)

	switch {
	case size == 0:
		return ErrSize
	case max == 0:
		return ErrMax
	}

	// set new increment settings
	if id == "" {
		r.ID, err = c.db.Setting(size, max)
		if err != nil {
			return DBError
		}
		return r
	}

	// update increment settings
	err = c.db.UpdateSetting(id, size, max)
	if err != nil {
		if c.db.IsNotFound(err) {
			return NotFound
		}
		return DBError
	}
	r.ID = id
	return r
}
