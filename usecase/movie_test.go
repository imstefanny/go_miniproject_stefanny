package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
	"time"
)

func TestCreateMovie(t *testing.T) {
	dateStr := "2023-10-07T18:45:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	data := dto.CreateMovieRequest{
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	movieData := model.Movie{
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("Create", movieData).Return(nil)

	service := NewMovieUsecase(mockMovieRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateMovieWithEmptyValue(t *testing.T) {
	data := dto.CreateMovieRequest{}

	movieData := model.Movie{}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("Create", movieData).Return(nil)

	service := NewMovieUsecase(mockMovieRepository)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllMovie(t *testing.T) {
	dateStr := "2023-10-07T18:45:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	mockMovie := []model.Movie{
		{
			ID:          1,
			Title:       "THE PILOT",
			Duration:    87,
			ReleaseDate: date,
			Genre:       "Drama, War",
			Rating:      4,
			Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
			Producer:    "Eric Gendarme, Thomas Lubeau",
			Director:    "Paul Doucet",
			Writer:      "Paul Doucet",
			Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
			Distributor: "Fulltime Studio",
			Type:        "13+",
		},
	}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("GetAll").Return(mockMovie, nil)

	service := NewMovieUsecase(mockMovieRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetMovie(t *testing.T) {
	dateStr := "2023-10-07T18:45:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	mockMovie := model.Movie{
		ID:          1,
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("Find").Return(mockMovie, nil)

	service := NewMovieUsecase(mockMovieRepository)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateMovie(t *testing.T) {
	dateStr := "2023-10-07T18:45:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	data := dto.CreateMovieRequest{
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	mockMovie := model.Movie{
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("Update", 1, mockMovie).Return(nil)
	mockMovieRepository.On("Find").Return(mockMovie, nil)

	service := NewMovieUsecase(mockMovieRepository)

	_, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestDeleteMovie(t *testing.T) {
	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("Delete", 1).Return(nil)

	service := NewMovieUsecase(mockMovieRepository)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetMovieByName(t *testing.T) {
	dateStr := "2023-10-07T18:45:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	mockMovie := model.Movie{
		Title:       "THE PILOT",
		Duration:    87,
		ReleaseDate: date,
		Genre:       "Drama, War",
		Rating:      4,
		Synopsis:    "Daniel (Hugo Becker), seorang pilot drone Prancis yang bertugas di Mali, mendapati istri dan putrinya diculik oleh teroris. Mereka akan dieksekusi jika Daniel tidak menggagalkan penangkapan seorang kepala jaringan jihad, Coulibaly (Yombi Hermann), yang dilakukan oleh timnya.",
		Producer:    "Eric Gendarme, Thomas Lubeau",
		Director:    "Paul Doucet",
		Writer:      "Paul Doucet",
		Cast:        "Hugo Becker, Eye Haidara, Kahina Carina, Grégory Fitoussi, Lalie Lanners, Jean-baptiste Anoumon, Elisabeth Mbai Mbanzani, Sayyid El Alami",
		Distributor: "Fulltime Studio",
		Type:        "13+",
	}

	mockMovieRepository := repository.NewMockMovieRepository()
	mockMovieRepository.On("GetMovieByName").Return(mockMovie, nil)

	service := NewMovieUsecase(mockMovieRepository)

	_, err := service.GetMovieByName("THE PILOT")
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}