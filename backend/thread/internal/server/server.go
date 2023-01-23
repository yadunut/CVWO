package server

import (
	context "context"
	"strings"

	"github.com/google/uuid"
	"github.com/yadunut/CVWO/backend/database"
	"github.com/yadunut/CVWO/backend/database/models"
	proto "github.com/yadunut/CVWO/backend/proto/thread"
	"github.com/yadunut/CVWO/backend/thread/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedThreadServiceServer
	DB     database.DB
	log    *zap.SugaredLogger
	config config.Config
}

// CreateThread implements thread.ThreadServiceServer
func (s *Server) CreateThread(ctx context.Context, req *proto.Thread) (*proto.ThreadResponse, error) {
	if strings.TrimSpace(req.Body) == "" {
		return &proto.ThreadResponse{Status: proto.ResponseStatus_FAILURE, Error: "body cannot be empty"}, nil
	}

	if strings.TrimSpace(req.Title) == "" {
		return &proto.ThreadResponse{Status: proto.ResponseStatus_FAILURE, Error: "tile cannot be empty"}, nil
	}

	if strings.TrimSpace(req.OwnerId) == "" {
		return &proto.ThreadResponse{Status: proto.ResponseStatus_FAILURE, Error: "owner cannot be empty"}, nil
	}
	ownerId, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return &proto.ThreadResponse{Status: proto.ResponseStatus_FAILURE, Error: "invalid owner id"}, nil
	}
	thread := models.NewThread(ownerId, req.Title, req.Body)

	if err = s.DB.Create(&thread).Error; err != nil {
		return &proto.ThreadResponse{Status: proto.ResponseStatus_FAILURE, Error: err.Error()}, nil
	}

	return &proto.ThreadResponse{
		Status: proto.ResponseStatus_SUCCESS,
		Thread: dbThreadToProto(thread),
	}, nil
}

// GetThread implements thread.ThreadServiceServer
func (*Server) GetThread(ctx context.Context, req *proto.GetThreadRequest) (*proto.ThreadResponse, error) {
	panic("unimplemented")
}

// GetThreads implements thread.ThreadServiceServer
func (s *Server) GetThreads(ctx context.Context, req *proto.GetThreadsRequest) (*proto.ThreadsResponse, error) {
	// get all dbThreads
	var dbThreads []models.Thread
	if len(req.Ids) == 0 {
		err := s.DB.Find(&dbThreads).Error
		if err != nil {
			return nil, err
		}
	}
	var protoThreads []*proto.Thread
	for _, t := range dbThreads {
		protoThreads = append(protoThreads, dbThreadToProto(t))
	}

	return &proto.ThreadsResponse{
		Status:  proto.ResponseStatus_SUCCESS,
		Error:   "",
		Threads: protoThreads,
	}, nil
}

func NewServer(DB database.DB, log *zap.SugaredLogger, c config.Config) *Server {
	return &Server{
		DB:                               DB,
		log:                              log,
		UnimplementedThreadServiceServer: proto.UnimplementedThreadServiceServer{},
		config:                           c,
	}
}

func dbThreadToProto(t models.Thread) *proto.Thread {
	return &proto.Thread{
		Id:      t.ID.String(),
		OwnerId: t.OwnerId.String(),
		Title:   t.Title,
		Body:    t.Body,
	}
}
