package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	srv Service
}

func New(srv Service) *Endpoint {
	return &Endpoint{
		srv: srv,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	days := e.srv.DaysLeft()

	res := fmt.Sprintf("Количество дней: %d", days)

	return ctx.String(http.StatusOK, res)
}
