package app

import (
	"log"
	"testing"

	"github.com/abdukahhor/swe/models"
	"github.com/abdukahhor/swe/storage"
)

func TestSet(t *testing.T) {
	db, err := storage.Connect("/tmp/testdb")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	c := New(db)

	reply := c.Set(models.Settings{ID: "", Size: 3, Max: 0})
	if reply != models.ErrSize {
		t.Error(reply, models.ErrSize)
	}

	reply = c.Set(models.Settings{ID: "", Size: 3, Max: 0})
	if reply != models.ErrMax {
		t.Error(reply, models.ErrMax)
	}

	reply = c.Set(models.Settings{ID: "", Size: 4, Max: 500})
	if reply.Code != 1 {
		t.Error(reply, "code is not success code")
	}
	if reply.ID == "" {
		t.Error(reply, "id is empty")
	}

	c.Increment(reply.ID)

	reply = c.Get(reply.ID)
	if reply.Code != 1 {
		t.Error(reply.Code, "reply is not success")
	}

	if reply.Num != 4 {
		t.Error(reply.Num, "number is not 4")
	}

	reply = c.Set(models.Settings{ID: reply.ID, Size: 3, Max: 100})
	if reply.Code != 1 {
		t.Error(reply, "code is not success code")
	}

	reply = c.Increment(reply.ID)
	if reply.Code != 1 {
		t.Error(reply.Code, "reply is not success")
	}

	reply = c.Get(reply.ID)
	if reply.Code != 1 {
		t.Error(reply.Code, "reply is not success")
	}

	if reply.Num != 7 {
		t.Error(reply.Num, "number is not 7")
	}
}
