package backend

import (
	"log"
	"net/http"
	"os"

	"./controller"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func HandleRequest() {
	http.HandleFunc("/", controller.HomePage)

	http.HandleFunc("/api/get/allbook", controller.GetAllBook)
	http.HandleFunc("/api/create/book", controller.CreateBook)
	http.HandleFunc("/api/delete/bookbyid", controller.DeleteBookByID)
	http.HandleFunc("/api/edit/bookbyid", controller.EditBookByID)

	http.ListenAndServe(getPort(), nil)
}

func main() {
	log.Print("The service is start on http://localhost" + getPort())
	HandleRequest()
}
