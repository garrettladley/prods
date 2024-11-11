package rand

import "golang.org/x/exp/rand"

func ChooseOne[T any](seed uint64, from ...T) T {
	return from[rand.Intn(len(from))]
}

func ChooseN[T any](seed uint64, n uint, from ...T) []T {
	if n > uint(len(from)) {
		return from
	}
	rand.Seed(seed)
	indices := rand.Perm(len(from))
	chosen := make([]T, n)
	for i := uint(0); i < n; i++ {
		chosen[i] = from[indices[i]]
	}
	return chosen
}
