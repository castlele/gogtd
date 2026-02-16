package repository

type Repo[T any, K comparable] interface {
	Create(entity T) error
	Get(key K) (T, error)
	List() ([]T, error)
	Update(entity T) error
	Delete(key K) (T, error)
}
