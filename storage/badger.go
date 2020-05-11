package storage

import (
	"encoding/binary"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/nats-io/nuid"
)

type db struct {
	conn *badger.DB
}

//Close closes db
func (p *db) Close() error {
	return p.conn.Close()
}

//Connect to embedded database,
//BadgerDB is persistent and fast key-value (KV) database
func Connect(path string) (DB, error) {

	bg, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &db{conn: bg}, nil
}

//Get gets number from storage by id
func (p db) Get(id string) (num uint64, err error) {
	var b []byte
	b, err = p.get([]byte(id))
	if err != nil {
		return
	}
	num = bytesToUint64(b[:8])
	return
}

func (p db) get(key []byte) ([]byte, error) {
	var rtn []byte
	err := p.conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			rtn = append([]byte{}, val...)
			return nil
		})
		return err
	})
	return rtn, err
}

//Increment increments number by id
func (p db) Increment(id string) (err error) {

	err = p.conn.Update(func(txn *badger.Txn) error {
		var (
			val []byte
			key = []byte(id)
		)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(v []byte) error {
			val = append([]byte{}, v...)
			return nil
		})
		if err != nil {
			return err
		}
		num := bytesToUint64(val[:8])
		size := bytesToUint64(val[8:16])
		max := bytesToUint64(val[16:24])
		num += size
		if num >= max {
			num = 0
		}
		copy(val[:8], uint64ToBytes(num))
		return txn.SetEntry(badger.NewEntry(key, val))
	})
	return
}

//inserts new increment settings,
//size of increment,
//max number of increment
//return id of increment, id is unique and random generated
func (p db) Setting(size uint64, max uint64) (id string, err error) {
	err = p.conn.Update(func(txn *badger.Txn) error {
		id = nuid.Next()
		val := make([]byte, 24)
		copy(val[:8], uint64ToBytes(0))
		copy(val[8:16], uint64ToBytes(size))
		copy(val[16:24], uint64ToBytes(max))
		return txn.SetEntry(badger.NewEntry([]byte(id), val))
	})
	return
}

//updates increment settings
func (p db) UpdateSetting(id string, size uint64, max uint64) (err error) {
	err = p.conn.Update(func(txn *badger.Txn) error {
		var (
			val []byte
			key = []byte(id)
		)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(v []byte) error {
			val = append([]byte{}, v...)
			return nil
		})
		if err != nil {
			return err
		}
		num := bytesToUint64(val[:8])
		max := bytesToUint64(val[16:24])
		if num >= max {
			copy(val[:8], uint64ToBytes(0))
		}
		copy(val[8:16], uint64ToBytes(size))
		copy(val[16:24], uint64ToBytes(max))
		return txn.SetEntry(badger.NewEntry(key, val))
	})
	return
}

func uint64ToBytes(i uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], i)
	return buf[:]
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

//IsNotFound -
func (p db) IsNotFound(err error) bool {
	return err == badger.ErrKeyNotFound
}
