package movie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core"
	"time"
)

type TMBDMovieRepository struct {
	bearerToken     string
	httpMovieClient http.Client
}

func NewTMDBMovieRepository(bearerToken string) *TMBDMovieRepository {
	return &TMBDMovieRepository{
		bearerToken: bearerToken,
		httpMovieClient: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

func (tmbdMovieRepository *TMBDMovieRepository) tmbdRequest(url string) ([]byte, error) {
	tmbdRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	tmbdRequest.Header.Set("Authorization", tmbdMovieRepository.bearerToken)
	tmbdResponse, err := tmbdMovieRepository.httpMovieClient.Do(tmbdRequest)
	if err != nil {
		return nil, err
	}

	if tmbdResponse.Body != nil {
		defer tmbdResponse.Body.Close()
	}

	if tmbdResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot reach API %s", url)
	}

	rawBody, err := ioutil.ReadAll(tmbdResponse.Body)
	if err != nil {
		return nil, err
	}

	return rawBody, nil
}

func (tmbdMovieRepository *TMBDMovieRepository) retrieveTopRatedMovies() ([]TMBDMovie, error) {
	topRatedBody, err := tmbdMovieRepository.tmbdRequest("https://api.themoviedb.org/3/movie/top_rated")
	if err != nil {
		return nil, err
	}

	topRatedMovies := TMBDTopRatedMovies{}
	err = json.Unmarshal(topRatedBody, &topRatedMovies)
	if err != nil {
		return nil, err
	}

	return topRatedMovies.Results, nil
}

func (tmbdMovieRepository *TMBDMovieRepository) retrieveMovieGenres() (map[int]string, error) {
	movieGenresBody, err := tmbdMovieRepository.tmbdRequest("https://api.themoviedb.org/3/genre/movie/list")
	if err != nil {
		return nil, err
	}

	movieGenres := TMBDMovieGenres{}
	err = json.Unmarshal(movieGenresBody, &movieGenres)
	if err != nil {
		return nil, err
	}
	genres := make(map[int]string)
	for _, genre := range movieGenres.Genres {
		genres[genre.Id] = genre.Name
	}

	return genres, nil
}

// FindAllMovies TODO: Return non adapter specific errors
func (tmbdMovieRepository *TMBDMovieRepository) FindAllMovies() ([]core.Movie, error) {
	topRatedMovies, err := tmbdMovieRepository.retrieveTopRatedMovies()
	if err != nil {
		return nil, err
	}

	movieGenres, err := tmbdMovieRepository.retrieveMovieGenres()
	if err != nil {
		return nil, err
	}

	var coreMovies []core.Movie
	for _, movie := range topRatedMovies {
		coreMovie := core.Movie{
			Genres:           nil,
			OriginalLanguage: movie.OriginalLanguage,
			OriginalTitle:    movie.OriginalTitle,
			ReleaseDate:      movie.ReleaseDate,
			Title:            movie.Title,
		}

		for _, genre := range movie.GenreIds {
			coreMovie.Genres = append(coreMovie.Genres, movieGenres[genre])
		}
		coreMovies = append(coreMovies, coreMovie)
	}

	return coreMovies, nil
}
