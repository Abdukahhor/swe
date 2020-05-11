package storage

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	var size, max uint64 = 3, 20
	con, err := Connect("/tmp/testdb")
	if err != nil {
		t.Error(err)
	}
	defer con.Close()
	id, err := con.Setting(size, max)
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
		exp += size
		num, err = con.Get(id)
		if err != nil {
			t.Error(err)
		}

		if exp >= max {
			exp = 0
		}
		fmt.Println(num, exp)
		if num != exp {
			t.Errorf("Returned = %d, Expected = %d", num, exp)
		}
	}
}
