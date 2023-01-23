package server

import (
	"context"
	"errors"
	"net/mail"

	"github.com/jackc/pgconn"
	"github.com/yadunut/CVWO/backend/auth/internal/config"
	"github.com/yadunut/CVWO/backend/auth/internal/utils"
	"github.com/yadunut/CVWO/backend/database"
	"github.com/yadunut/CVWO/backend/database/models"
	proto "github.com/yadunut/CVWO/backend/proto/auth"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	proto.UnimplementedAuthServiceServer
	DB     database.DB
	log    *zap.SugaredLogger
	config config.Config
}

func NewServer(DB database.DB, log *zap.SugaredLogger, c config.Config) *Server {
	return &Server{
		DB:                             DB,
		log:                            log,
		UnimplementedAuthServiceServer: proto.UnimplementedAuthServiceServer{},
		config:                         c,
	}
}

// Login implements proto.AuthServiceServer
func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	s.log.Infof("Received Connection")
	// check if its email or username
	var user models.User
	_, err := mail.ParseAddress(req.UsernameOrEmail)
	if err != nil {
		err = s.DB.Where(&models.User{Username: req.UsernameOrEmail}).First(&user).Error
	} else {
		err = s.DB.Where(&models.User{Email: req.UsernameOrEmail}).First(&user).Error
	}

	if err != nil {
		return &proto.LoginResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}
	if user.PasswordHash == "" {
		return &proto.LoginResponse{Status: proto.ResponseStatus_FAILURE, Error: "username or password invalid"}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return &proto.LoginResponse{Status: proto.ResponseStatus_FAILURE, Error: "username or password invalid"}, nil
	}

	tokenString, err := utils.GenerateJwtToken(user.ID.String(), s.config)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Status: proto.ResponseStatus_SUCCESS, Token: tokenString}, nil
}

// Register implements proto.AuthServiceServer
func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	s.log.Infof("Received Connection")
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: "invalid email address"}, nil
	}
	// a validation library should be used here, but the surface area is small, so manual validation is good enough
	if err := utils.ValidateUsername(req.Username); err != nil {
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	user := models.NewUser(req.Email, req.Username, string(hash))
	if err = s.DB.Create(&user).Error; err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			// Postgres error for duplicate
			if pgError.Code == utils.PG_DUPLICATE {
				return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: "username already in use"}, nil
			}
			return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: pgError.Message}, nil
		}

		s.log.Error(err)
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}
	tokenString, err := utils.GenerateJwtToken(user.ID.String(), s.config)
	if err != nil {
		return nil, err
	}

	return &proto.RegisterResponse{Status: proto.ResponseStatus_SUCCESS, Error: req.String(), Token: tokenString}, nil
}

func (s *Server) Verify(ctx context.Context, req *proto.VerifyRequest) (*proto.VerifyResponse, error) {
	id, err := utils.ParseJwtToken(req.Token, s.config)
	if err != nil {
		return &proto.VerifyResponse{Status: proto.ResponseStatus_FAILURE, Message: err.Error()}, nil
	}
	var user models.User
	err = s.DB.Where(&models.User{Model: models.Model{ID: id}}).First(&user).Error
	if err != nil {
		return &proto.VerifyResponse{Status: proto.ResponseStatus_FAILURE, Message: err.Error()}, nil
	}
	return &proto.VerifyResponse{Status: proto.ResponseStatus_SUCCESS, Id: id.String()}, nil
}
