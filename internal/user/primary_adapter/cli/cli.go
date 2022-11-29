package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core/domain/services"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/secondary_adapter/repositories/movie"
)

var rootCmd = &cobra.Command{
	Use:   "hexa",
	Short: "hexa is a CLI to call a hexagonal architecture code template",
}

var countMoviesCmd = &cobra.Command{
	Use:   "count-movies",
	Short: "Count the number of movies",
	Long:  `Count the number of movies`,
	Run: func(cmd *cobra.Command, args []string) {
		movieRepository := movie.NewTMDBMovieRepository(os.Getenv("TMBD_TOKEN"))
		movieDomainService := services.MovieDomainService{}
		fmt.Println(movieDomainService.CountMovies(movieRepository))
	},
}

var getAllMoviesCmd = &cobra.Command{
	Use:   "get-movies",
	Short: "Get all movies",
	Long:  `Get all movies`,
	Run: func(cmd *cobra.Command, args []string) {
		movieRepository := movie.NewTMDBMovieRepository(os.Getenv("TMBD_TOKEN"))
		movieDomainService := services.MovieDomainService{}
		movies := movieDomainService.RetrieveAllMovies(movieRepository)
		for _, movie := range movies {
			fmt.Printf("Movie: %s (%s)\n", movie.Title, movie.ReleaseDate)
		}
	},
}

func InitCli() *cobra.Command {
	rootCmd.AddCommand(countMoviesCmd)
	rootCmd.AddCommand(getAllMoviesCmd)
	return rootCmd
}
