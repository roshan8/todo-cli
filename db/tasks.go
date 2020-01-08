package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("task")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Init : won't run automatically
func Init(dbPath string) error {

	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second}) // Can't use :=  because scope will change to local instead of package leve

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}
