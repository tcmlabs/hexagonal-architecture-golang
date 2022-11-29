package primary

import (
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/core/domain"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/core/ports/secondary"
)

type MovieCommand interface {
	CountMovies(repository secondary.MovieRepository) int
	RetrieveAllMovies(repository secondary.MovieRepository) []services.Movie
}
