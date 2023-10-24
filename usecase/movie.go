package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"reflect"

	// "miniproject/middlewares"
	"errors"
)

type MovieUsecase interface {
	Create(movie dto.CreateMovieRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	Delete(id int) error
	Update(id int, movie dto.CreateMovieRequest) (model.Movie, error)
}

type movieUsecase struct {
	movieRepository		repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) *movieUsecase {
	return &movieUsecase{movieRepository: movieRepo}
}

func isEmptyField(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func validateCreateMovieRequest(req dto.CreateMovieRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *movieUsecase) Create(movie dto.CreateMovieRequest) error {
	e := validateCreateMovieRequest(movie)
	
	if e!= nil {
		return e
	}

	movieData := model.Movie{
		Title: movie.Title,
		Duration: movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		Genre: movie.Genre,
		Rating: movie.Rating,
		Synopsis: movie.Synopsis,
		Producer: movie.Producer,
		Director: movie.Director,
		Writer: movie.Writer,
		Cast: movie.Cast,
		Distributor: movie.Distributor,
		Type: movie.Type,
	}
	err := s.movieRepository.Create(movieData)

	if err != nil {
		return err
	}

	return nil
}

func (s *movieUsecase) GetAll() (interface{}, error) {
	movies, err := s.movieRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *movieUsecase) Find(id int) (interface{}, error) {
	movie, err := s.movieRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieUsecase) Delete(id int) error {
	err := s.movieRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *movieUsecase) Update(id int, movie dto.CreateMovieRequest) (model.Movie, error) {
	movieData, err := s.movieRepository.Find(id)

	if err != nil {
		return model.Movie{}, err
	}

	movieData.Title = movie.Title
	movieData.Duration = movie.Duration
	movieData.ReleaseDate = movie.ReleaseDate
	movieData.Genre = movie.Genre
	movieData.Rating = movie.Rating
	movieData.Synopsis = movie.Synopsis
	movieData.Producer = movie.Producer
	movieData.Director = movie.Director
	movieData.Writer = movie.Writer
	movieData.Cast = movie.Cast
	movieData.Distributor = movie.Distributor
	movieData.Type = movie.Type
	
	e := s.movieRepository.Update(id, movieData)

	if e != nil {
		return model.Movie{}, e
	}

	movieUpdated, err := s.movieRepository.Find(id)

	if err != nil {
		return model.Movie{}, err
	}
	return movieUpdated, nil
}
