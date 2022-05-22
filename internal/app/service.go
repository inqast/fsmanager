package app

import (
	pb "github.com/inqast/fsmanager/pkg/api"
)

type tserver struct {
	repo Repository
	pb.UnimplementedFamilySubServer
}

func New(repo Repository) *tserver {
	return &tserver{repo: repo}
}
