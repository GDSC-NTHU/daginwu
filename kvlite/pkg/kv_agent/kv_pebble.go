package kv_agent

import (
	"io"
	"log"

	"github.com/cockroachdb/pebble"
)

type KVPebble struct {
	db *pebble.DB
}

func (kv *KVPebble) Init() error {

	db, err := pebble.Open(
		"kvlite",
		&pebble.Options{},
	)

	if err != nil {
		return err
	}

	kv.db = db

	return nil
}

func (kv *KVPebble) Shutdown() error {

	return kv.Shutdown()
}

func (kv *KVPebble) Get(key string) ([]byte, io.Closer, error) {

	value, closer, err := kv.db.Get([]byte(key))
	if err != nil {
		return value, closer, err
	}

	return value, closer, nil
}

func (kv *KVPebble) Set(key []byte, value []byte) error {

	err := kv.db.Set(
		key,
		value,
		pebble.Sync,
	)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (kv *KVPebble) Delete(key []byte) error {

	err := kv.db.Delete(
		key,
		pebble.Sync,
	)

	if err != nil {
		log.Println(err)
	}

	return err
}
