package db

import (
	"encoding/binary"
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

func CreateTask(task string) (int, error) {
	// Stores the task in BoltDB bucket

	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return 0, nil
}

func DeleteTask(key int) error {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})

	if err != nil {
		return err
	}
	return nil
}

// AllTasks Returns every items in the bucket
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
