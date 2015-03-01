package main

import (
	//"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	//"text/template"
)


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	//homeTempl = template.Must(template.ParseFiles("console.html"))
)

type WeightMap [][][]float64

func (weights *WeightMap) init(letters, rows, cols int) {
	r := rand.New(rand.NewSource(999))
	*weights = make([][][]float64, letters)
	for i := 0; i < letters; i++ {
		(*weights)[i] = make([][]float64, rows)
		for j := 0; j < rows; j++ {
			(*weights)[i][j] = make([]float64, cols)
			for k := 0; k < cols; k++ {
				(*weights)[i][j][k] = r.Float64()
			}
		}
	}
}

/*
func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}
*/

func serveWS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	weights := WeightMap{}
	weights.init(26, 9, 9)
	//js, err := json.Marshal(weights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	conn.WriteJSON(weights)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	//http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWS)
	log.Println("Listening...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
