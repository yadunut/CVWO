package server

import (
	context "context"
	"errors"
	"net/mail"

	"github.com/jackc/pgconn"
	"github.com/yadunut/CVWO/backend/auth/internal/database"
	"github.com/yadunut/CVWO/backend/auth/internal/proto"
	"github.com/yadunut/CVWO/backend/auth/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	DB  database.DB
	log *zap.SugaredLogger
	proto.UnimplementedAuthServiceServer
}

func NewServer(DB database.DB, log *zap.SugaredLogger) *Server {
	return &Server{
		DB:                             DB,
		log:                            log,
		UnimplementedAuthServiceServer: proto.UnimplementedAuthServiceServer{},
	}
}

// Login implements proto.AuthServiceServer
func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// check if its email or username
	var user database.User
	_, err := mail.ParseAddress(req.UsernameOrEmail)
	if err != nil {
		err = nil
		err = s.DB.Where(&database.User{Username: req.UsernameOrEmail}).First(&user).Error
	} else {
		err = s.DB.Where(&database.User{Email: req.UsernameOrEmail}).First(&user).Error
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
	// TODO: Generate Token

	return &proto.LoginResponse{Status: proto.ResponseStatus_SUCCESS, Token: "Successful Login"}, nil
}

// Register implements proto.AuthServiceServer
func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
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
	if err = s.DB.Create(database.NewUser(req.Email, req.Username, string(hash))).Error; err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			// Postgres error for duplicate
			if pgError.Code == "23505" {
				return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: "username already in use"}, nil
			}
			return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: pgError.Message}, nil
		}

		s.log.Error(err)
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}
	// TODO: generate JWT / other auth token with expiry

	return &proto.RegisterResponse{Status: proto.ResponseStatus_SUCCESS, Error: req.String(), Token: "Correct token!"}, nil
}
