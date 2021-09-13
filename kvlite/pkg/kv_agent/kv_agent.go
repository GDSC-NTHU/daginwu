package kv_agent

import "io"

type KVAgent interface {
	Init() error
	Shutdown() error
	Get(key string) ([]byte, io.Closer, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
}
