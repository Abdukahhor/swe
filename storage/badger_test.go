package storage

import (
	"testing"

	"github.com/abdukahhor/swe/model"
)

func TestDB(t *testing.T) {
	var s = model.Settings{Size: 3, Max: 20}
	con, err := Connect("/tmp/testdb")
	if err != nil {
		t.Error(err)
	}
	defer con.Close()
	id, err := con.Setting(s)
	if err != nil {
		t.Error(err)
	}

	num, err := con.Get(id)
	if err != nil {
		t.Error(err)
	}

	if num != 0 {
		t.Errorf("Returned = %d, Expected = 0", num)
	}
	var exp uint64
	for i := 1; i <= 40; i++ {
		err = con.Increment(id)
		if err != nil {
			t.Error(err)
		}
		exp += s.Size
		num, err = con.Get(id)
		if err != nil {
			t.Error(err)
		}

		if exp >= s.Max {
			exp = 0
		}
		if num != exp {
			t.Errorf("Returned = %d, Expected = %d", num, exp)
		}
	}
}
