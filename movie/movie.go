package movie

import "fmt"

func init() {
	fmt.Println("init : movie")
}
func ReviewMovie(name string, rating float32) {
	fmt.Printf("You vote %v is %v\n", name, rating)
}

func FindMovieName(imdbID string) string {
	switch imdbID {
	case "aaa":
		{
			return "asdasd"
		}

	case "bbb":
		{
			return "aaaa"
		}

	}
	return "not found"
}
