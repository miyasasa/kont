package storage

import "github.com/boltdb/bolt"

func (s *Store) Ping() error {
	return s.db.View(func(tx *bolt.Tx) error {
		return nil
	})
}
