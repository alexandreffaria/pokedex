package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

const firstPageLocation string = "https://pokeapi.co/api/v2/location-area/"

func FetchLocations() {
	response, err := http.Get(firstPageLocation)
	if err != nil {
		fmt.Println("error making the request: ", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading the body: ", err)
		return
	}

	fmt.Println(string(body))
}
