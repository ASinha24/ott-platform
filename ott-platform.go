package ottplatform

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/asinha24/ott-platform/api"
	"github.com/asinha24/ott-platform/movies"
	"github.com/asinha24/ott-platform/movies/model"
)

type OttManager interface {
	GetMovies(ctx context.Context, ottName string) ([]*api.CreateMoveieResponse, error)
	CreateMovie(ctx context.Context, ottName string, movie *api.MovieRequest) (*api.CreateMoveieResponse, error)
	UpdateMovie(ctx context.Context, ottName string, movieID string, newMovie *api.MovieRequest) (*api.CreateMoveieResponse, error)
	DeleteMovie(ctx context.Context, ottName string, movieId string) error
}

type OttService struct {
	movieStore movies.MovieStore
	ottStore   movies.OTTPlatform
}

func NewOttService(movieStore movies.MovieStore, ottStore movies.OTTPlatform) *OttService {
	return &OttService{
		movieStore: movieStore,
		ottStore:   ottStore,
	}
}

func (s *OttService) GetMovies(ctx context.Context, ottName string) ([]*api.CreateMoveieResponse, error) {
	ott, err := s.ottStore.GetOTTPlatformByName(ctx, ottName)
	if err != nil {
		return nil, &api.OTTError{Code: api.MovieNotFount, Message: fmt.Sprintf("ott not found %s ", ottName), Description: err.Error()}
	}

	movies, err := s.movieStore.FindALLOTTMovies(ctx, ott.ID)
	if err != nil {
		return nil, &api.OTTError{Code: api.MovieNotFount, Message: fmt.Sprintf("no movies found on ott %s ", ottName), Description: err.Error()}
	}

	resp := []*api.CreateMoveieResponse{}

	for _, m := range movies {
		resp = append(resp, &api.CreateMoveieResponse{
			MovieRequest: &api.MovieRequest{Name: m.Name, ReleaseYear: m.ReleaseYear, Director: m.Director},
			ID:           m.ID,
		})
	}

	return resp, nil
}

func (s *OttService) CreateMovie(ctx context.Context, ottName string, movie *api.MovieRequest) (*api.CreateMoveieResponse, error) {
	ott, err := s.ottStore.GetOTTPlatformByName(ctx, ottName)
	if err != nil {
		return nil, &api.OTTError{Code: api.MovieNotFount, Message: fmt.Sprintf("ott not found %s", ottName), Description: err.Error()}
	}

	newMovie, err := s.movieStore.CreateMovie(ctx, &model.Movie{
		ID:          uuid.NewUUID().String(),
		Name:        movie.Name,
		ReleaseYear: movie.ReleaseYear,
		Director:    movie.Director,
		OTTID:       ott.ID,
	})

	if err != nil {
		return nil, &api.OTTError{Code: api.MovieCreationFailed, Message: fmt.Sprintf("can not create new movie %s ", ottName), Description: err.Error()}
	}

	return &api.CreateMoveieResponse{MovieRequest: movie, ID: newMovie.ID}, nil

}

func (s *OttService) UpdateMovie(ctx context.Context, ottName string, movieID string, newMovie *api.MovieRequest) (*api.CreateMoveieResponse, error) {
	ottID, err := s.checkValidMovie(ctx, ottName, movieID)
	if err != nil {
		return nil, err
	}

	updateMovie, err := s.movieStore.UpdateMovie(ctx, &model.Movie{
		ID:          movieID,
		Name:        newMovie.Name,
		ReleaseYear: newMovie.ReleaseYear,
		Director:    newMovie.Director,
		OTTID:       ottID,
	})
	if err != nil {
		return nil, &api.OTTError{Code: api.MovieUpdateFailed, Message: fmt.Sprintf("can not update existing movie %s ", ottName), Description: err.Error()}
	}

	return &api.CreateMoveieResponse{MovieRequest: newMovie, ID: updateMovie.ID}, nil
}

func (s *OttService) DeleteMovie(ctx context.Context, ottName string, movieID string) error {
	_, err := s.checkValidMovie(ctx, ottName, movieID)
	if err != nil {
		return err
	}
	return s.movieStore.DeleteMovie(ctx, movieID)
}

func (s *OttService) checkValidMovie(ctx context.Context, ottName, movieId string) (string, error) {
	ott, err := s.ottStore.GetOTTPlatformByName(ctx, ottName)
	if err != nil {
		return "", &api.OTTError{Code: api.MovieNotFount, Message: fmt.Sprintf("ott not found %s ", ottName), Description: err.Error()}
	}

	movie, err := s.movieStore.FindMovieByID(ctx, movieId)
	if err != nil {
		return "", &api.OTTError{Code: api.MovieNotFount, Message: fmt.Sprintf("movie not found %s ", ottName), Description: err.Error()}
	}

	if movie.OTTID != ott.ID {
		return "", &api.OTTError{Code: api.UnauthorizedMovie, Message: fmt.Sprintf("movie %s is not authorized on OTT %s ", movie.Name, ottName), Description: err.Error()}
	}

	return ott.ID, nil
}
