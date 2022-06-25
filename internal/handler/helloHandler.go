package handler

type HelloHandler

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world!")
}
