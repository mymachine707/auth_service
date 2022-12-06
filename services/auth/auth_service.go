package user

import (
	"context"
	"fmt"
	"log"
	"mymachine707/protogen/blogpost"
	"mymachine707/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	stg storage.Interfaces
	blogpost.UnimplementedUserServiceServer
}

func NewAuthService(stg storage.Interfaces) *userService {
	return &userService{
		stg: stg,
	}
}

func (s *userService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "Ok",
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, req *blogpost.CreateUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- CreateUser ---->>>")

	// Todo password hash

	id := uuid.New()

	err := s.stg.AddUser(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *blogpost.UpdateUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- UpdateUser ---->>>")
	// Todo password hash

	err := s.stg.UpdateUser(req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(req.Id) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID---!: %s", err.Error())
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *blogpost.DeleteUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- DeleteUser ---- >>>")

	user, err := s.stg.GetUserByID(req.Id) // maqsad tekshirish rostan  ham create bo'ldimi?

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	err = s.stg.DeleteUser(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteUser: %s", err.Error())
	}

	return user, nil
}

func (s *userService) GetUserList(ctx context.Context, req *blogpost.GetUserListRequest) (*blogpost.GetUserListResponse, error) {
	fmt.Println("<<< ---- GetUserList ---->>>")

	res, err := s.stg.GetUserList(int(req.Offset), int(req.Limit), req.Search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserList: %s", err.Error())
	}

	return res, nil
}
func (s *userService) GetUserById(ctx context.Context, req *blogpost.GetUserByIDRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- GetUserById ---->>>")

	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}
