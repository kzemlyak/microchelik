package controllers

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/usecases"
	pb "github.com/kzemlyak/microchelik/social/pkg/api/social"
)

func (c *Controller) CreateOrder(
	ctx context.Context,
	req *pb.CreateProfileRequest,
) (*pb.CreateProfileResponse, error) {
	// 1. validation (in middleware)

	profile, err := c.UseCase.CreateProfile(ctx, &usecases.CreateProfilePayload{
		UserID:    req.UserId,
		Nickname:  req.Nickname,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	})
	if err != nil {
		return nil, err // обработается на уровне middleware
	}

	response := &pb.CreateProfileResponse{
		Id:        profile.ID,
		UserId:    profile.UserID,
		Nickname:  profile.Nickname,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
	}

	return response, nil
}
