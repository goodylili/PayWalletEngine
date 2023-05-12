package server

import (
	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")
	return nil

}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		log.Println(err)
	}
}
