package app

import (
	pb "github.com/inqast/fsmanager/pkg/api"
)

type tserver struct {
	lastID int64
	repo   Repository
	pb.UnimplementedFamilySubServer
}

func New(repo Repository) *tserver {
	return &tserver{repo: repo}
}
