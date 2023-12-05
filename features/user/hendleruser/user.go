package hendleruser

import (
	"BE-REPO/features/user"
	"BE-REPO/helper/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	srv user.Service
}

func New(s user.Service) user.Handler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(UserRequest)

		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input tidak sesuai",
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var inputProses = new(user.User)
		inputProses.Email = input.Email
		inputProses.PhoneNumber = input.PhoneNumber
		inputProses.FullName = input.FullName
		inputProses.UserName = input.UserName
		inputProses.Password = string(hashedPassword)
		// inputProses.Password = input.Password

		result, err := uc.srv.Register(*inputProses)
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "dobel input nama",
				})
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input tidak sesuai",
			})
		}

		strToken, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(UserResponse)
		response.ID = result.ID
		response.Email = result.Email
		response.PhoneNumber = result.PhoneNumber
		response.FullName = result.FullName
		response.UserName = result.UserName
		response.Token = strToken

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success",
			"data":    response,
		})
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		userData, err := uc.srv.GetUserByEmail(input.Email)
		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "pengguna tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "terjadi kesalahan",
			})
		}
		err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input.Password))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "password salah",
			})
		}

		strToken, err := jwt.GenerateJWT(userData.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.ID = userData.ID
		response.Email = userData.Email
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    response,
		})
	}
}
