package middle

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/myrachanto/microservice/tag/src/pasetos"
)

const (
	authorisationHeaderKey = "Authorization"
	authorisationType      = "Bearer"
)

func PasetoAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get(authorisationHeaderKey)
		if len(authorizationHeader) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header not provided")
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization format provided")
		}
		authtype := fields[0]
		if authtype != authorisationType {
			return echo.NewHTTPError(http.StatusUnauthorized, "That type of Authorization is not allowed here!")
		}
		accessToken := fields[1]
		maker, _ := pasetos.NewPasetoMaker()
		_, err := maker.VerifyToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "That token is invalid!")
		}
		return next(c)
	}
}
