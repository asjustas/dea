package main

import "github.com/dgraph-io/badger"

type Storage struct {
	kv *badger.KV
}

func NewStorage() *Storage {
	return new(Storage)
}

func (storage *Storage) Open() error {
	opts := badger.DefaultOptions
	opts.Dir = "./data"
	opts.ValueDir = "./data"

	kv, err := badger.NewKV(&opts)
	storage.kv = kv

	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) Add(domain string) error {
	return storage.kv.Set([]byte(domain), []byte("1"), 0x00)
}

func (storage *Storage) Exists(domain string) (bool, error) {
	return storage.kv.Exists([]byte(domain))
}

func (storage *Storage) Close() {
	storage.kv.Close()
}
