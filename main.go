package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "only POST method supported")
		return
	}

	var v ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	resp := play(v)   // action happen here
	resp := actiononedge(nextDirection)
	fmt.Fprint(w, resp)
}

func play(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	commands := []string{"F", "R", "L", "T"}
	rand := rand2.Intn(4)
	return commands[rand]
}

func actiononedge (input ArenaUpdate, input PlayerState) ( return nextDirection string) {
	player := PlayerState{}
	arena  := ArenaUpdate{}	
	log.Printf("IN: %#v", input) // log only

	if ( player.X == 0 && player.y == 0 ){
	return "R"
	}
	else if ( player.X == 0 && player.y == arena.dimension[1]){
	commands := []string{"L", "T"}
	rand := rand2.Intn(2)
	return commands[rand]
	}
	else if ( player.Y == arena.dimension[1] && player.X == arena.dimension[0]){
		commands := []string{"F", "L", "T"}
		rand := rand2.Intn(3)
		return commands[rand]
	}
	else if ( player.Y == 0] && player.X == arena.dimension[0]){
		commands := []string{"F", "R", "T"}
		rand := rand2.Intn(3)
		return commands[rand]
	}
	else {
		commands := []string{"F", "R", "L", "T"}
		rand := rand2.Intn(4)
		return commands[rand]
	}
}
