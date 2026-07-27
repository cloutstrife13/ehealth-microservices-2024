package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type AuthService struct {
	Db           *gorm.DB
	FirebaseAuth *auth.Client
}

func NewAuthService(db *gorm.DB) *AuthService {
	ctx := context.Background()

	if os.Getenv("FIREBASE_AUTH_EMULATOR_HOST") == "" {
		os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "firebase-auth:9099")
	}

	config := &firebase.Config{
		ProjectID: "ehealth-local-project",
	}

	app, err := firebase.NewApp(ctx, config, option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("Error initializing Firebase App: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v\n", err)
	}

	return &AuthService{
		Db:           db,
		FirebaseAuth: client,
	}
}

func generateActivationCode() *models.VerificationStatus {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength := 8

	b := make([]byte, codeLength)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}

	code := string(b)
	expiresAt := time.Now().AddDate(0, 0, 30)

	return &models.VerificationStatus{
		VerificationCode: code,
		Status:           "UNVERIFIED",
		ExpiresAt:        expiresAt,
	}
}

func (s *AuthService) Register(newUser *models.User, password string) error {
	ctx := context.Background()

	params := (&auth.UserToCreate{}).
		Email(newUser.Email).
		Password(password).
		DisplayName(fmt.Sprintf("%s %s", newUser.Name, newUser.Surname))

	fbUser, err := s.FirebaseAuth.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create user in Firebase: %w", err)
	}

	newUser.ID = fbUser.UID
	newUser.VerificationStatus = generateActivationCode()

	if err := s.Db.Create(newUser).Error; err != nil {
		_ = s.FirebaseAuth.DeleteUser(ctx, fbUser.UID)
		return fmt.Errorf("failed to save user to database: %w", err)
	}

	return nil
}
