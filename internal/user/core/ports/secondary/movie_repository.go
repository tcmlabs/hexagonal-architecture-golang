package secondary

import "tcmlabs.fr/hexagonal_architecture_golang/internal/user/core"

type MovieRepository interface {
	FindAllMovies() []core.Movie
}
