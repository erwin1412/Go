package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// Define custom claims
type jwtCustomClaims struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	DivisionID uint   `json:"division_id"`
	PositionID uint   `json:"position_id"`
	jwt.RegisteredClaims
}

// JWTMiddleware - Middleware untuk validasi JWT
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari header Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Ambil token setelah "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		// Parse dan validasi token
		claims := &jwtCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Secret key untuk validasi token
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// Menyimpan klaim ke dalam context untuk digunakan di handler
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("divisionID", claims.DivisionID)
		c.Set("positionID", claims.PositionID)

		return next(c)
	}
}
