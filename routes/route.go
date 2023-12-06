package routes

import (
	"BE-REPO/features/posting"
	"BE-REPO/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.Handler, pu posting.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
	RoutePosting(e, pu)
}

func RouteUser(e *echo.Echo, uc user.Handler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}

func RoutePosting(e *echo.Echo, pu posting.Handler) {
	e.POST("/posting", pu.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
