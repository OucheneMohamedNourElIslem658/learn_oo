package routers

import (
	"context"
	"encoding/json"
	"errors"

	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
	usersRepo "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProfilesServiceServer struct {
	authpb.UnimplementedProfilesServiceServer
	usersRepo *usersRepo.ProfilesRepository
}

func NewProfilesServiceServer() *ProfilesServiceServer {
	return &ProfilesServiceServer{
		usersRepo: usersRepo.NewProfilesRepository(),
	}
}

func (s *ProfilesServiceServer) GetProfile(ctx context.Context, req *emptypb.Empty) (*authpb.Profile, error) {
	id, ok := ctx.Value("id").(string)
	if !ok || id == "" {
		return nil, errors.New("requester is not a user")
	}

	profile, apiErr := s.usersRepo.GetUser(id, "courses,author_profile,image")
	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	BioToJson, err := json.Marshal(profile.AuthorProfile.Bio)

	if err != nil {
		return nil, errors.New("failed to marshal bio")
	}

	return &authpb.Profile{
		Id:            profile.ID,
		FullName:      profile.FullName,
		Email:         profile.Email,
		EmailVerified: profile.EmailVerified,
		Image: func() *authpb.File {
			if profile.Image == nil {
				return nil
			}
			return &authpb.File{
				Id:           uint64(profile.Image.ID),
				Url:          profile.Image.URL,
				ThumbnailUrl: profile.Image.ThumbnailURL,
				Height:       int32(profile.Image.Height),
				Width:        int32(profile.Image.Width),
			}
		}(),
		AuthorProfile: func() *authpb.Author {
			if profile.AuthorProfile == nil {
				return nil
			}
			return &authpb.Author{
				Id:      profile.AuthorProfile.ID,
				Bio:     string(BioToJson),
				Balance: int32(profile.AuthorProfile.Balance),
			}
		}(),
		Courses: func() []*authpb.Course {
			courses := make([]*authpb.Course, len(profile.Courses))
			for i, course := range profile.Courses {
				bio, err := json.Marshal(course.Author.Bio)
				if err != nil {
					continue
				}
				courses[i] = &authpb.Course{
					Id:          uint64(course.ID),
					Title:       course.Title,
					Description: course.Description,
					AuthorId:    course.AuthorID,
					Image: func() *authpb.File {
						if course.Image == nil {
							return nil
						}
						return &authpb.File{
							Id:           uint64(course.Image.ID),
							Url:          course.Image.URL,
							ThumbnailUrl: course.Image.ThumbnailURL,
							Height:       int32(course.Image.Height),
							Width:        int32(course.Image.Width),
						}
					}(),
					Price:       course.Price,
					Language:    string(course.Language),
					Level:       string(course.Level),
					Duration:    uint64(course.Duration),
					Rate:        course.Rate,
					RatersCount: uint64(course.RatersCount),
					IsCompleted: course.IsCompleted,
					Video: func() *authpb.File {
						if course.Video == nil {
							return nil
						}
						return &authpb.File{
							Id:           uint64(course.Video.ID),
							Url:          course.Video.URL,
							ThumbnailUrl: course.Video.ThumbnailURL,
							Height:       int32(course.Video.Height),
							Width:        int32(course.Video.Width),
						}
					}(),
					Author: func() *authpb.Author {
						if course.Author == nil {
							return nil
						}
						return &authpb.Author{
							Id:      course.Author.ID,
							Bio:     string(bio),
							Balance: int32(course.Author.Balance),
						}
					}(),
				}
			}
			return courses
		}(),
	}, nil
}

func (s *ProfilesServiceServer) UpgradeToAuthor(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	id, ok := ctx.Value("id").(string)
	if !ok || id == "" {
		return nil, errors.New("requester is not a user")
	}

	apiErr := s.usersRepo.UpgradeToAuthor(id)
	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return nil, nil
}

func (s *ProfilesServiceServer) DowngradeToUser(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	id, ok := ctx.Value("author_id").(string)
	if !ok || id == "" {
		return nil, errors.New("requester is not an author")
	}

	apiErr := s.usersRepo.DowngradeFromAuthor(id)
	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return nil, nil
}

func (s *ProfilesServiceServer) GetAuthor(ctx context.Context, req *authpb.GetAuthorRequest) (*authpb.Author, error) {
	author, apiErr := s.usersRepo.GetAuthor(req.Id, "accomplishments,user")
	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	bio, err := json.Marshal(author.Bio)
	if err != nil {
		return nil, errors.New("error marshaling bio")
	}

	user := author.User

	accomplishments := []*authpb.File{};
	for _, a := range author.Accomplishments {
		accomplishments = append(accomplishments, &authpb.File{
			Url: a.URL,
			Id:           uint64(a.ID),
			Height:       int32(a.Height),
			Width:        int32(a.Width),
			ThumbnailUrl: a.ThumbnailURL,
		})
	}

	response := &authpb.Author{
		Id: author.ID,
		Bio: string(bio),
		Balance: int32(author.Balance),
		UserProfile: &authpb.Profile{
			Id: author.UserID,
			Email: user.Email,
			FullName: user.FullName,
			EmailVerified: user.EmailVerified,
		},
		Accomplishments: accomplishments,
	}

	return response, nil
}