package wild

type (
	Middleware func(next Handler) Handler
	Handler    func(c Context) error
)

func Compose(middlewares []Middleware) Middleware {
	return func(next Handler) Handler {
		c := next
		for _, middleware := range middlewares {
			c = middleware(c)
		}
		return c
	}
}
