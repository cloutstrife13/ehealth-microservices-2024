package auth

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"gorm.io/gorm"
)

type AuthService struct {
	Db *gorm.DB
}

func generateActivationCode() *models.VerificationStatus {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength := 8

	b := make([]byte, codeLength)
	for i := range b {
		// Pick a random character from the charset securely
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}

	code := string(b)
	expiresAt := time.Now().AddDate(0, 0, 30)

	return &models.VerificationStatus{
		VerificationCode: code,
		Status: 	      "UNVERIFIED",
		ExpiresAt:        expiresAt,
	}
}

func (s *AuthService) Register(newUser *models.User) error {
	newUser.VerificationStatus = generateActivationCode()

	// Publish VerificationLetterEvent to PubSub topic verification-print-jobs.

	return s.Db.Create(newUser).Error
}
