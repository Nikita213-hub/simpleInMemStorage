package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPut(t *testing.T) {
// 	tests := []struct {
// 		key string
// 		value string

// 	} {
// 		key: "",

// 	}
// }

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		wanted interface{}
	}{{
		name:   "Test empty key",
		key:    "",
		wanted: nil,
	}, {
		name:   "Test noraml key value",
		key:    "test",
		wanted: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := NewStorage()
			strg.Put("test", true)
			v := strg.Get(tt.key)
			assert.Equal(t, tt.wanted, v)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		isErr bool
	}{{
		name:  "Test existing key",
		key:   "test",
		isErr: false,
	}, {
		name:  "Test non-existing key",
		key:   "non-exist",
		isErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := NewStorage()
			strg.Put("test", true)
			err := strg.Delete(tt.key)
			if tt.isErr {
				assert.Error(t, err)
			}
		})
	}
}
