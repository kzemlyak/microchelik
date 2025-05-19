package controllers

import (
	"github.com/kzemlyak/microchelik/social/internal/app/usecases"
	pb "github.com/kzemlyak/microchelik/social/pkg/api/social"
)

type Deps struct {
	UseCase usecases.UseCase
}

type Controller struct {
	pb.UnimplementedSocialServiceServer
	Deps
}

func NewController(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
