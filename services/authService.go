package services

import (
	"fmt"
	"learn-go-auth/config"
	"learn-go-auth/models"
	"learn-go-auth/repository"
	"learn-go-auth/utils"
)

var userRepo repository.UserRepository

// RegisterUser registers a new user, including email verification and OTP
func RegisterUser(username, email, password, role string) (*models.User, error) {

	if userRepo == nil {
		userRepo = repository.NewUserRepository(config.DB)
	}
	
	// Validate user data
	userData := models.User{Username: username, Email: email, PasswordHash: password, Role: role}
	if validationErr := utils.ValidateStruct(userData, "Username", "Email", "PasswordHash"); validationErr != "" {
		return nil, fmt.Errorf(validationErr)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Generate OTP for email verification
	otp := utils.GenerateOTP()

	// Create user object
	user := &models.User{
		Username:        username,
		Email:           email,
		PasswordHash:    hashedPassword,
		Role:            role,
		OTP:             otp,
		IsEmailVerified: false,
	}

	// Save user to the database
	if err := userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Send OTP email for verification
	if err := SendOTPEmail(email, otp); err != nil {
		return nil, fmt.Errorf("failed to send OTP email: %v", err)
	}

	return user, nil
}

// LoginUser authenticates the user and returns a JWT token
func LoginUser(email, password string) (map[string]interface{}, error) {
	// Find user by email
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Validate password
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating JWT: %v", err)
	}

	return map[string]interface{}{
		"token": token,
		"role":  user.Role,
	}, nil
}

// ResendOTP resends the OTP for email verification
func ResendOTP(email string) error {
	// Find user by email
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Generate new OTP
	otp := utils.GenerateOTP()
	user.OTP = otp

	// Save OTP to the database
	if err := userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update OTP: %v", err)
	}

	// Send OTP email
	if err := SendOTPEmail(user.Email, otp); err != nil {
		return fmt.Errorf("failed to send OTP email: %v", err)
	}

	return nil
}

// SendVerificationOTP generates and sends an OTP for email verification
func SendVerificationOTP(email string) error {
	// Find user by email
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Generate OTP
	otp := utils.GenerateOTP()
	user.OTP = otp

	// Save OTP to DB
	if err := userRepo.Update(user); err != nil {
		return fmt.Errorf("error saving OTP: %v", err)
	}

	// Send OTP to user email
	if err := SendOTPEmail(user.Email, otp); err != nil {
		return fmt.Errorf("error sending OTP: %v", err)
	}

	return nil
}

// VerifyEmail verifies the OTP entered by the user
func VerifyEmail(email, otp string) (string, error) {
	// Find user by email
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	// Check if the OTP matches
	if user.OTP != otp {
		return "", fmt.Errorf("invalid OTP")
	}

	// Mark email as verified
	user.IsEmailVerified = true
	user.OTP = "" // Clear OTP

	// Update user in the database
	if err := userRepo.Update(user); err != nil {
		return "", fmt.Errorf("error updating email verification status: %v", err)
	}

	return "Email verified successfully", nil
}
