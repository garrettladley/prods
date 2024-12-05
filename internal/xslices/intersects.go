package xslices

func Intersects[T comparable](a []T, b []T) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	baseMap := make(map[T]struct{}, len(a))
	for _, v := range a {
		baseMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := baseMap[v]; exists {
			return true
		}
	}

	return false
}
