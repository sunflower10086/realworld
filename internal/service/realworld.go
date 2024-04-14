package service

import (
	"context"
	"realworld/internal/biz"

	pb "realworld/api/realworld/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{
		uc: uc,
	}
}

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *emptypb.Empty) (*pb.User, error) {
	return &pb.User{
		User: &pb.User_User{
			Email:    "123",
			Token:    "123",
			Username: "123",
			Bio:      "123",
			Image:    "123",
		},
	}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.Profile, error) {
	return &pb.Profile{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.Profile, error) {
	return &pb.Profile{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.Profile, error) {
	return &pb.Profile{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticles, error) {
	return &pb.MultipleArticles{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticles, error) {
	return &pb.MultipleArticles{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) AddCommentsArticle(ctx context.Context, req *pb.AddCommentsArticleRequest) (*pb.SingleComment, error) {
	return &pb.SingleComment{}, nil
}
func (s *RealWorldService) GetCommentsArticle(ctx context.Context, req *pb.GetCommentsArticleRequest) (*pb.MultipleComments, error) {
	return &pb.MultipleComments{}, nil
}
func (s *RealWorldService) DeleteCommentsArticle(ctx context.Context, req *pb.DeleteCommentsArticleRequest) (*pb.MultipleComments, error) {
	return &pb.MultipleComments{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleArticle, error) {
	return &pb.SingleArticle{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.ListOfTags, error) {
	return &pb.ListOfTags{}, nil
}
