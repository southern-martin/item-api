package application

import (
	"github.com/southern-martin/item-api/src/controller"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/item", controller.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/item/{id}", controller.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/item/search", controller.ItemController.Search).Methods(http.MethodPost)
}
