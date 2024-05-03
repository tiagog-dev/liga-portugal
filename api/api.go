package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// data structure
type stadium struct {
	Team        string  `json:"team"`
	Stadium     string  `json:"stadium"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Capacity    string  `json:"capacity"`
	Year_Opened string  `json:"year_opened"`
}

var stadiums = []stadium{
	{Team: "Arouca", Stadium: "Estádio Municipal de Arouca", City: "Arouca", Latitude: 40.93288730907449, Longitude: -8.250399827957153, Capacity: "5,600", Year_Opened: "2006"},
	{Team: "Benfica", Stadium: "Estádio da Luz", City: "Lisboa", Latitude: 38.752648359600805, Longitude: -9.184699058532715, Capacity: "64,642", Year_Opened: "2003"},
	{Team: "Boavista", Stadium: "Estádio do Bessa Século XXI", City: "Porto", Latitude: 41.16224317458759, Longitude: -8.642533421516418, Capacity: "28,263", Year_Opened: "1972"},
	{Team: "Casa Pia", Stadium: "Estádio Pina Manique", City: "Lisboa", Latitude: 38.737540246765576, Longitude: -9.203839302062988, Capacity: "5,000", Year_Opened: "1954"},
	{Team: "Estrela Amadora", Stadium: "Estádio José Gomes", City: "Amadora", Latitude: 38.7512186618, Longitude: -9.22344577288, Capacity: "9,288", Year_Opened: "1957"},
	{Team: "Chaves", Stadium: "Estádio Municipal Eng. Manuel Branco Teixeira", City: "Chaves", Latitude: 41.75055999558073, Longitude: -7.46495246887207, Capacity: "8,396", Year_Opened: "1949"},
	{Team: "Estoril Praia", Stadium: "Estádio António Coimbra da Mota", City: "Estoril", Latitude: 38.71585362022704, Longitude: -9.406411796808243, Capacity: "5,000", Year_Opened: "1930"},
	{Team: "Famalicão", Stadium: "Estádio Municipal 22 de Junho", City: "Vila Nova de Famalicão", Latitude: 41.40144303510442, Longitude: -8.52239191532135, Capacity: "5,305", Year_Opened: "1952"},
	{Team: "Farense", Stadium: "Estádio de São Luís", City: "Faro", Latitude: 37.02282647562398, Longitude: -7.928507924079895, Capacity: "7.000", Year_Opened: "1923"},
	{Team: "Gil Vicente", Stadium: "Estádio Cidade de Barcelos", City: "Barcelos", Latitude: 41.55112933644525, Longitude: -8.62303376197815, Capacity: "10,046", Year_Opened: "1924"},
	{Team: "Moreirense", Stadium: "Estádio C. J. de Almeida Freitas", City: "Moreira de Cónegos", Latitude: 41.37808861279996, Longitude: -8.354759216308594, Capacity: "6,153", Year_Opened: "2002"},
	{Team: "Portimonense", Stadium: "Estádio Municipal de Portimão", City: "Portimão", Latitude: 37.135700416749145, Longitude: -8.539837002754211, Capacity: "4,961", Year_Opened: "1937"},
	{Team: "Porto", Stadium: "Estádio do Dragão", City: "Porto", Latitude: 41.1617908503053, Longitude: -8.583739399909973, Capacity: "50,033", Year_Opened: "2003"},
	{Team: "Rio Ave", Stadium: "Estádio dos Arcos", City: "Vila do Conde", Latitude: 41.362774699560454, Longitude: -8.740214109420776, Capacity: "5,300", Year_Opened: "1984"},
	{Team: "Braga", Stadium: "Estádio Municipal de Braga", City: "Braga", Latitude: 41.5624, Longitude: -8.4308, Capacity: "30,86", Year_Opened: "2003"},
	{Team: "Sporting", Stadium: "Estádio José Alvalade XXI", City: "Lisboa", Latitude: 38.761207228440945, Longitude: -9.160875678062439, Capacity: "50,095", Year_Opened: "2003"},
	{Team: "Vitória Guimarães", Stadium: "Estádio D. Afonso Henriques", City: "Guimarães", Latitude: 41.445903097491914, Longitude: -8.300943374633789, Capacity: "30.029", Year_Opened: "1965"},
	{Team: "Vizela", Stadium: "Estádio do Vizela", City: "Caldas de Vizela", Latitude: 41.3887627414459, Longitude: -8.307321667671204, Capacity: "6,565", Year_Opened: "1989"},
}

// getStadium responds with the list of all stadiums as JSON.
func getStadium(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stadiums)
}

// setting up the endpoint.
func main() {
	router := gin.Default()
	router.GET("/stadiums", getStadium)

	router.Run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/v1/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{status: 'running'}")
	})

	log.Println("listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error launching REST API server: %v", err)
	}
}
