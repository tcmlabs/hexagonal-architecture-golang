package secondary

import (
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core/domain"
)

type MovieRepository interface {
	FindAllMovies() ([]services.Movie, error)
}
