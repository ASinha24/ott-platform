package movies

import (
	"context"
	"errors"

	"github.com/asinha24/ott-platform/movies/model"
)

type MovieStore interface {
	CreateMovie(ctx context.Context, movie *model.Movie) (*model.Movie, error)
	UpdateMovie(ctx context.Context, newMovie *model.Movie) (*model.Movie, error)
	DeleteMovie(ctx context.Context, movieID string) error
	FindMovieByID(ctx context.Context, movieID string) (*model.Movie, error)
	FindALLOTTMovies(ctx context.Context, ottID string) ([]*model.Movie, error)
}

type StoreMovieInMem struct {
	movies map[string]*model.Movie
}

func NewMovieStore() MovieStore {
	return &StoreMovieInMem{movies: make(map[string]*model.Movie)}
}

func (s *StoreMovieInMem) CreateMovie(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	s.movies[movie.ID] = movie
	return movie, nil
}

func (s *StoreMovieInMem) UpdateMovie(ctx context.Context, newMovie *model.Movie) (*model.Movie, error) {
	oldMovie, exist := s.movies[newMovie.ID]
	if !exist {
		return nil, errors.New("movie not found")
	}

	oldMovie.Name = newMovie.Name
	oldMovie.Director = newMovie.Director
	oldMovie.ReleaseYear = newMovie.ReleaseYear
	oldMovie.OTTID = newMovie.OTTID

	s.movies[newMovie.ID] = oldMovie
	return oldMovie, nil
}

func (s *StoreMovieInMem) DeleteMovie(ctx context.Context, movieID string) error {
	_, ok := s.movies[movieID]
	if !ok {
		return errors.New("movie not found")
	}

	delete(s.movies, movieID)
	return nil
}

func (s *StoreMovieInMem) FindMovieByID(ctx context.Context, movieID string) (*model.Movie, error) {
	movie, exist := s.movies[movieID]
	if !exist {
		return nil, errors.New("movie not found")
	}
	return movie, nil
}

func (s *StoreMovieInMem) FindALLOTTMovies(ctx context.Context, ottID string) ([]*model.Movie, error) {
	movies := []*model.Movie{}
	for _, movie := range s.movies {
		if movie.OTTID == ottID {
			movies = append(movies, movie)
		}

	}
	return movies, nil
}
