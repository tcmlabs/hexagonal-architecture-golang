package services

import (
	services "tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/core/domain"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/core/ports/secondary"
)

type MovieDomainService struct {
}

// CountMovies TODO: Deal with errors
func (movieDomainService MovieDomainService) CountMovies(repository secondary.MovieRepository) int {
	movies, _ := repository.FindAllMovies()
	return len(movies)
}

// RetrieveAllMovies TODO: Deal with errors
func (movieDomainService MovieDomainService) RetrieveAllMovies(repository secondary.MovieRepository) []services.Movie {
	movies, _ := repository.FindAllMovies()
	return movies
}
