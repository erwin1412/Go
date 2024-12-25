package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}

	token, err := user.GenerateToken()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func fetchService(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid reference ID")
	}
	return nil
}

func Register(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		Password    string `json:"password" validate:"required"`
		PhoneNumber string `json:"phone_number"`
		DivisionID  uint   `json:"division_id" validate:"required"`
		PositionID  uint   `json:"position_id" validate:"required"`
		JoinDate    string `json:"join_date"` // Tambahkan field join_date (opsional)
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Validasi DivisionID dengan memanggil Division Service
	divisionURL := "http://localhost:8083/api/divisions/" + string(rune(input.DivisionID))
	if err := fetchService(divisionURL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Division ID",
		})
	}

	// Validasi PositionID dengan memanggil Position Service
	positionURL := "http://localhost:8082/api/positions/" + string(rune(input.PositionID))
	if err := fetchService(positionURL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Position ID",
		})
	}

	// Validasi password
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to hash password",
		})
	}

	// Atur nilai join_date
	var joinDate *time.Time
	if input.JoinDate != "" {
		parsedDate, err := time.Parse("2006-01-02", input.JoinDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid join_date format. Use YYYY-MM-DD.",
			})
		}
		joinDate = &parsedDate
	}

	// Buat objek user
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		PhoneNumber: input.PhoneNumber,
		DivisionID:  input.DivisionID,
		PositionID:  input.PositionID,
		JoinDate:    joinDate.Format("2006-01-02"), // Tetapkan nilai join_date
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, user)
}
