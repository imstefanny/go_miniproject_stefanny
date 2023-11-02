package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type MovieController interface{
}

type movieController struct {
	movieUsecase usecase.MovieUsecase
}

func NewMovieController(movieUsecase usecase.MovieUsecase) *movieController {
	return &movieController{movieUsecase}
}

func (u *movieController) Create(c echo.Context) error {
	movie := dto.CreateMovieRequest{}

	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.movieUsecase.Create(movie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": movie,
	})
}

func (u *movieController) GetAll(c echo.Context) error {
	movies, err := u.movieUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": movies,
	})
}

func (u *movieController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := u.movieUsecase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": movie,
	})
}

func (u *movieController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.movieUsecase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted succesfully",
	})
}

func (u *movieController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie := dto.CreateMovieRequest{}

	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	movieUpdated, err := u.movieUsecase.Update(id, movie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": movieUpdated,
		"message": "success update data",
	})
}

func (u *movieController) GetMovieRecommendations(c echo.Context) error {
	recommend, err := u.movieUsecase.GetMovieRecommendations()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"recommended_movie": recommend,
	})
}
