package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetData(url string) (mapPage, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var data mapPage

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	return data, nil
}
