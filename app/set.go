package app

import (
	"github.com/abdukahhor/swe/models"
)

//Set - increment settings
//if id is empty sets new increment settings, else updates
func (c Core) Set(s models.Settings) models.Reply {
	var (
		r   = models.Success
		err error
	)

	switch {
	case s.Size == 0:
		return models.ErrSize
	case s.Max == 0:
		return models.ErrMax
	}

	// set new increment settings
	if s.ID == "" {
		r.ID, err = c.db.Setting(s)
		if err != nil {
			return models.DBError
		}
		return r
	}

	// update increment settings
	err = c.db.UpdateSetting(s)
	if err != nil {
		if c.db.IsNotFound(err) {
			return models.NotFound
		}
		return models.DBError
	}
	r.ID = s.ID
	return r
}
