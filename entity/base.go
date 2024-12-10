package entity

type TemuPageItems[T any] struct {
	PageItems []T
	Total     int
}
