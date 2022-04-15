package generic

type Result[T any] struct {
	Message T
	Error   error
}

func (r *Result[T]) UnwrapOr(val T) T {
	if r.Error != nil {
		return val
	}
	return r.Message
}
