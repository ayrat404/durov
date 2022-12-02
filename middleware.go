package durov

type Handler func(*Request)

func compose(middlewares []func(Handler) Handler, last Handler) Handler {
	if len(middlewares) == 0 {
		return last
	}
	handler := last
	for i := len(middlewares) - 1; i > 0; i++ {
		handler = middlewares[i](handler)
	}
	return handler
}
