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
	var p PlayerState
	/*
		var myx = v.Arena.State["https://cloudbowl-samples-java-quarkus-yngbkt2j3a-uc.a.run.app"].X
		var myy = v.Arena.State["https://cloudbowl-samples-java-quarkus-yngbkt2j3a-uc.a.run.app"].Y
		var myd = v.Arena.State["https://cloudbowl-samples-java-quarkus-yngbkt2j3a-uc.a.run.app"].Direction
		var myh = v.Arena.State["https://cloudbowl-samples-java-quarkus-yngbkt2j3a-uc.a.run.app"].WasHit
	*/
	
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//resp := play(v) // action happen here
	resp2 := actiononedge(v, p)

	//	fmt.Fprint(w, resp)
	fmt.Fprint(w, resp2)

}

/*
func play(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	commands := []string{"F", "R", "L", "T", "T", "T"}
	rand := rand2.Intn(6)
	return commands[rand]
}
*/

func actiononedge(arena ArenaUpdate, player PlayerState) (response string) {
	//player :=  PlayerState{}
	//arena := ArenaUpdate{}
	log.Printf("IN: %#v", arena) // log only
	//log.Printf("IN: %#p", player)
	if player.X == 0 && player.Y == 0 {
		return "R"
	} else if player.X == 0 && player.Y == arena.Arena.Dimensions[1] {
		commands := []string{"L", "T"}
		rand := rand2.Intn(2)
		return commands[rand]
	} else if player.Y == arena.Arena.Dimensions[1] && player.X == arena.Arena.Dimensions[0] {
		commands := []string{"F", "L", "T", "T"}
		rand := rand2.Intn(4)
		return commands[rand]
	} else if player.Y == 0 && player.X == arena.Arena.Dimensions[0] {
		commands := []string{"F", "R", "T", "T"}
		rand := rand2.Intn(4)
		return commands[rand]
	} else {
		commands := []string{"F", "R", "L", "T"}
		rand := rand2.Intn(4)
		return commands[rand]
	}
}
