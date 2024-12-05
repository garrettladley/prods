package types

type Pair[T any] struct {
	First  T `json:"first"`
	Second T `json:"second"`
}
