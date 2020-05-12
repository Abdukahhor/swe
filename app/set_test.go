package app

import (
	"log"
	"testing"

	"github.com/abdukahhor/swe/storage"
)

func TestSet(t *testing.T) {
	db, err := storage.Connect("/tmp/testdb")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	c := New(db)

	reply := c.Set("", 0, 300)
	if reply != ErrSize {
		t.Error(reply, ErrSize)
	}

	reply = c.Set("", 3, 0)
	if reply != ErrMax {
		t.Error(reply, ErrMax)
	}

	reply = c.Set("", 4, 500)
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

	reply = c.Set(reply.ID, 3, 100)
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
