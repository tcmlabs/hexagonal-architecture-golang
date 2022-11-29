package movie

type TMBDMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type TMBDTopRatedMovies struct {
	Page         int         `json:"page"`
	Results      []TMBDMovie `json:"results"`
	TotalPages   int         `json:"total_pages"`
	TotalResults int         `json:"total_results"`
}

type TMBDGenre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TMBDMovieGenres struct {
	Genres []TMBDGenre `json:"genres"`
}

type T struct {
	Genres []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}
