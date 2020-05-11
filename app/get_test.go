package app

import (
	"log"
	"testing"

	"github.com/abdukahhor/swe/storage"
)

func TestGet(t *testing.T) {
	db, err := storage.Connect("/tmp/testdb")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	c := New(db)

	reply := c.Get("")
	if reply != ErrID {
		t.Error(reply, ErrID)
	}

	reply = c.Get("234234234")
	if reply != NotFound {
		t.Error(reply, NotFound)
	}

	reply = c.Set("", 4, 500)
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
}
