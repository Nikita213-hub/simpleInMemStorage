package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	store map[string]interface{}
	mx    sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		store: make(map[string]interface{}),
		mx:    sync.Mutex{},
	}
}

func (s *Storage) Put(key string, value interface{}) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.store[key] = value
	return nil
}

func (s *Storage) Get(key string) interface{} {
	s.mx.Lock()
	defer s.mx.Unlock()
	v, ok := s.store[key]
	if !ok {
		return nil
	}
	return v
}

func (s *Storage) Delete(key string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	_, ok := s.store[key]
	if !ok {
		return errors.New("no value with such key")
	}
	delete(s.store, key)
	return nil
}
