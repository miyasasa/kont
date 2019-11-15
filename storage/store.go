package storage

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

func (s *Store) Ping() error {
	return s.db.View(func(tx *bolt.Tx) error {
		return nil
	})
}

func (s *Store) PUT(key string, i interface{}) error {
	return s.db.Update(func(tx *bolt.Tx) error {

		content, err := json.Marshal(i)
		if err != nil {
			return err
		}

		bucket := tx.Bucket(BitBucket)
		err = bucket.Put([]byte(key), content)

		if err != nil {
			return err
		}

		return nil
	})
}

func (s *Store) GET(key string, i interface{}) error {
	return s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BitBucket)

		encoded := bucket.Get([]byte(key))
		if encoded == nil {
			return fmt.Errorf("Record not available")
		}

		err := json.Unmarshal(encoded, i)
		if err != nil {
			return fmt.Errorf("Record can not mapped to given interface")
		}

		return nil
	})
}

func (s *Store) ForEach(fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BitBucket)

		err := bucket.ForEach(fn)
		if err != nil {
			return err
		}

		return nil
	})
}
