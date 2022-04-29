package wild

type (
	Middleware func(next Handler) Handler
	Handler    func(c Context) error
)

func Compose(middlewares []Middleware) Middleware {
	return func(next Handler) Handler {
		c := next
		for i := len(middlewares); i >= 0; i -= 1 {
			c = middlewares[i](c)
		}
		return c
	}
}
