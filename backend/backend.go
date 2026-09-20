package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	Config "github.com/ShuffleKamori/go-practice/config"
	"github.com/ShuffleKamori/go-practice/events"
)

const apiBase = "http://127.0.0.1:8010"

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func Start_backend() bool {
	fmt.Println("Backend initialize! Client -> " + apiBase)

	client := &http.Client{Timeout: 5 * time.Second}

	for req := range events.Request {
		switch req.Type {
		case "GET_CATALOG":
			catalog, err := getCatalog(client)
			if err != nil {
				fmt.Println("GET_CATALOG error:", err)
				continue
			}
			events.Result <- catalog
		}
	}

	return true
}

func getCatalog(client *http.Client) ([]Config.Buycard, error) {
	resp, err := client.Get(apiBase + "/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items []product
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, err
	}

	catalog := make([]Config.Buycard, 0, len(items))
	for _, p := range items {
		catalog = append(catalog, Config.Buycard{
			Name:  p.Name,
			Price: p.Price,
			SizeX: 200,
			SizeY: 200,
		})
	}
	return catalog, nil
}
