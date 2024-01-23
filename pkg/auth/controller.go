package auth

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/bio426/vugopo/datasource"
	"github.com/bio426/vugopo/pkg/core"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type AuthCtl core.Controller

const (
	JwtSecret          = "mysecret"
	CookieName         = "vugopo_jwt"
	TokenDurationHours = 1
)

func (ctl *AuthCtl) Register(c echo.Context) error {
	body := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// service
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return err
	}
	_, err = datasource.Postgres.ExecContext(
		c.Request().Context(),
		"insert into users(username,password,role) values ($1,$2,$3)",
		body.Username, hashedPassword, "user",
	)
	if err != nil {
		return err
	}
	// ~service

	return c.NoContent(http.StatusOK)
}

func (ctl *AuthCtl) Login(c echo.Context) error {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	// service
	var (
		id       int32
		username string
		password string
		role     string
	)
	row := datasource.Postgres.QueryRowContext(
		c.Request().Context(),
		"select id, username, password, role from users where username = $1",
		body.Username,
	)
	if err := row.Scan(&id, &username, &password, &role); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// expiryLimit := time.Now().Add(time.Hour * TokenDurationHours)
	expiryLimit := time.Now().Add(time.Minute * 1)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: jwt.NewNumericDate(expiryLimit),
	})
	token, err := claims.SignedString([]byte(JwtSecret))
	if err != nil {
		return err
	}
	// ~service

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Expires:  expiryLimit,
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	res := struct {
		Username   string    `json:"username"`
		Role       string    `json:"role"`
		ExpiryDate time.Time `json:"expiryDate"`
	}{
		Username:   username,
		Role:       role,
		ExpiryDate: expiryLimit,
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Logout(c echo.Context) error {

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Expires:  time.Now().Add(time.Second * -1),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

var Controller = &AuthCtl{}
