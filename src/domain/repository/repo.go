package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrExists     = errors.New("already exists")
	ErrEmptyKey   = errors.New("empty key")
	ErrNilKeyFunc = errors.New("nil key func")
)

type repoImpl[T any, K comparable] struct {
	path    string
	keyFunc func(T) K

	mu sync.RWMutex
}

func NewFPRepo[T any, K comparable](
	path string,
	keyFunc func(T) K,
) (*repoImpl[T, K], error) {
	if keyFunc == nil {
		return nil, ErrNilKeyFunc
	}

	r := &repoImpl[T, K]{
		path:    path,
		keyFunc: keyFunc,
	}

	if dir := filepath.Dir(path); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err := r.writeJSONAtomic([]T{}); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return r, nil
}

func (this *repoImpl[T, K]) Create(entity T) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	items, err := this.readSlice()

	if err != nil {
		return err
	}

	key := this.keyFunc(entity)

	if this.isZeroComparable(key) {
		return ErrEmptyKey
	}

	for _, it := range items {
		if this.keyFunc(it) == key {
			return ErrExists
		}
	}

	items = append(items, entity)

	return this.writeJSONAtomic(items)
}

func (this *repoImpl[T, K]) Get(key K) (T, error) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	var zero T

	items, err := this.readSlice()

	if err != nil {
		return zero, err
	}

	for _, it := range items {
		if this.keyFunc(it) == key {
			return it, nil
		}
	}

	return zero, ErrNotFound
}

func (this *repoImpl[T, K]) List() ([]T, error) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	items, err := this.readSlice()

	if err != nil {
		return nil, err
	}

	out := make([]T, len(items))
	copy(out, items)

	return out, nil
}

func (this *repoImpl[T, K]) Update(entity T) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	items, err := this.readSlice()

	if err != nil {
		return err
	}

	key := this.keyFunc(entity)

	if this.isZeroComparable(key) {
		return ErrEmptyKey
	}

	for i, it := range items {
		if this.keyFunc(it) == key {
			items[i] = entity
			return this.writeJSONAtomic(items)
		}
	}

	return ErrNotFound
}

func (this *repoImpl[T, K]) Delete(key K) (T, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	var zero T
	items, err := this.readSlice()

	if err != nil {
		return zero, err
	}

	for i, it := range items {
		if this.keyFunc(it) != key {
			continue
		}

		items = append(items[:i], items[i+1:]...)

		if err := this.writeJSONAtomic(items); err != nil {
			return zero, err
		}

		return it, nil
	}

	return zero, ErrNotFound
}

func (this *repoImpl[T, K]) readSlice() ([]T, error) {
	bytes, err := os.ReadFile(this.path)

	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return []T{}, nil
	}

	var items []T

	if err := json.Unmarshal(bytes, &items); err != nil {
		return nil, err
	}

	if items == nil {
		items = []T{}
	}

	return items, nil
}

func (this *repoImpl[T, K]) writeJSONAtomic(items []T) error {
	bytes, err := json.MarshalIndent(items, "", "  ")

	if err != nil {
		return err
	}

	dir := filepath.Dir(this.path)
	tmp, err := os.CreateTemp(dir, ".tmp-*.json")

	if err != nil {
		return err
	}

	tmpName := tmp.Name()
	ok := false

	defer func() {
		tmp.Close()

		if !ok {
			os.Remove(tmp.Name())
		}
	}()

	if _, err := tmp.Write(bytes); err != nil {
		return err
	}

	if err := tmp.Sync(); err != nil {
		return err
	}

	if err := tmp.Close(); err != nil {
		return err
	}

	if err := os.Rename(tmpName, this.path); err != nil {
		return err
	}

	ok = true

	return nil
}

func (this *repoImpl[T, K]) isZeroComparable(k K) bool {
	var zero K

	return k == zero
}
