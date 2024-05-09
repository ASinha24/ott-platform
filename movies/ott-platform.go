package movies

import (
	"context"
	"errors"

	"github.com/asinha24/ott-platform/movies/model"
)

type OTTPlatform interface {
	CreateOTT(context.Context, *model.OTT) error
	GetAllOTTPLatform(context.Context) ([]*model.OTT, error)
	GetOTTPlatformByName(context.Context, string) (*model.OTT, error)
	GetOTTPLatformByID(context.Context, string) (*model.OTT, error)
}

type StoreOttInMem struct {
	OTTPlatforms map[string]*model.OTT
}

func NewOTTPlatform() OTTPlatform {
	return &StoreOttInMem{OTTPlatforms: make(map[string]*model.OTT)}
}

func (s *StoreOttInMem) CreateOTT(ctx context.Context, ott *model.OTT) error {
	s.OTTPlatforms[ott.ID] = ott
	return nil
}

func (s *StoreOttInMem) GetOTTPlatformByName(ctx context.Context, ottName string) (*model.OTT, error) {
	for _, ott := range s.OTTPlatforms {
		if ott.Name == ottName {
			return ott, nil
		}
	}

	return nil, errors.New("ott not found")
}

func (s *StoreOttInMem) GetOTTPLatformByID(ctx context.Context, ottOId string) (*model.OTT, error) {
	ott, exist := s.OTTPlatforms[ottOId]
	if !exist {
		return nil, errors.New("ott not found")
	}

	return ott, nil
}

func (s *StoreOttInMem) GetAllOTTPLatform(ctx context.Context) ([]*model.OTT, error) {
	resp := []*model.OTT{}
	for _, v := range s.OTTPlatforms {
		resp = append(resp, &model.OTT{ID: v.ID, Name: v.Name})
	}

	return resp, nil
}
