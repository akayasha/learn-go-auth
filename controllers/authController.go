package controllers

import (
	"github.com/gin-gonic/gin"
	"learn-go-auth/services"
	"learn-go-auth/utils"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var registerData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"` // Accept role input
	}

	if err := c.BindJSON(&registerData); err != nil {
		utils.RespondError(c, 400, "Invalid request data")
		return
	}

	user, err := services.RegisterUser(registerData.Username, registerData.Email, registerData.Password, registerData.Role)
	if err != nil {
		utils.RespondError(c, 400, err.Error())
		return
	}

	utils.Respond(c, 201, "User registered successfully", user)
}

func ResendOTP(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		utils.RespondError(c, 400, "Email is required")
		return
	}

	err := services.ResendOTP(email)
	if err != nil {
		utils.RespondError(c, 400, err.Error())
		return
	}

	utils.Respond(c, 200, "OTP resent successfully", nil)
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		utils.RespondError(c, 400, "Invalid request data")
		return
	}

	response, err := services.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		utils.RespondError(c, 401, err.Error())
		return
	}

	utils.Respond(c, 200, "Login successful", response)
}

// VerifyEmail handles email verification using OTP
func VerifyEmail(c *gin.Context) {
	var verifyData struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := c.BindJSON(&verifyData); err != nil {
		utils.RespondError(c, 400, "Invalid request data")
		return
	}

	message, err := services.VerifyEmail(verifyData.Email, verifyData.OTP)
	if err != nil {
		utils.RespondError(c, 400, err.Error())
		return
	}

	utils.Respond(c, 200, message, nil)
}
