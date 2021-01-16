package v1

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/fboccacini/goauth/pkg/api/google.golang.org/protobuf/v1"
)

const (
	apiVersion = "v1"
)

type goAuthServiceServer struct {
	db *sql.DB
}

func NewGoAuthServiceServer(db *sql.DB) v1.GoAuthServiceServer {
	return &goAuthServiceServer{db: db}
}

func (s *goAuthServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'",
				apiVersion, api)
		}
	}
	return nil
}

func (s *goAuthServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *goAuthServiceServer) Signup(ctx context.Context, req *v1.LoginRequest) (*v1.AuthenticateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	// }

	// // insert ToDo entity data
	// res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
	// 	req.ToDo.Title, req.ToDo.Description, reminder)
	// if err != nil {
	// 	return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
	// }

	// // get ID of creates ToDo
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
	// }

	return &v1.AuthenticateResponse{
		Api:     apiVersion,
		Status:  true,
		Token:   "50m3t0k3n",
		Message: "",
	}, nil
}

func (s *goAuthServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.AuthenticateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return &v1.AuthenticateResponse{
		Api:     apiVersion,
		Status:  true,
		Token:   "50m3t0k3n",
		Message: "",
	}, nil
}

func (s *goAuthServiceServer) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return &v1.AuthenticateResponse{
		Api:     apiVersion,
		Status:  true,
		Token:   "50m3t0k3n",
		Message: "",
	}, nil
}
