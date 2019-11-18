package storage

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var RepositoryBucket = []byte("repository")
var Storage = initStorage()

type Store struct {
	db *bolt.DB
}

func initStorage() *Store {
	db, err := bolt.Open("miya.db", 0755, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatalf("DB connection can not opened %v", err)
	}

	err = initBuckets(db)
	if err != nil {
		log.Fatalf("Bucket can not created %v", err)
	}

	return &Store{db: db}
}

func initBuckets(db *bolt.DB) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, _ = tx.CreateBucketIfNotExists(RepositoryBucket)
		return nil
	}); err != nil {
		return err
	}

	return nil
}
