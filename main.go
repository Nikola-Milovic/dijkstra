package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Primer inputa
// '{ "graf": [ { "izvor": "A", "destinacija": "B", "tezina": 2 }, { "izvor": "A", "destinacija": "D", "tezina": 1 }, { "izvor": "A", "destinacija": "C", "tezina": 5 }, { "izvor": "B", "destinacija": "C", "tezina": 3 }, { "izvor": "B", "destinacija": "D", "tezina": 2 }, { "izvor": "D", "destinacija": "E", "tezina": 1 }, { "izvor": "D", "destinacija": "C", "tezina": 3 }, { "izvor": "E", "destinacija": "C", "tezina": 1 } , { "izvor": "E", "destinacija": "F", "tezina": 2 } , { "izvor": "C", "destinacija": "F", "tezina": 5 } ], "od": "A", "do": "C" }'
// '{ "graf": [ { "izvor": "A", "destinacija": "B", "tezina": 7 }, { "izvor": "A", "destinacija": "C", "tezina": 9 }, { "izvor": "A", "destinacija": "F", "tezina": 14 }, { "izvor": "B", "destinacija": "C", "tezina": 10 }, { "izvor": "B", "destinacija": "D", "tezina": 15 }, { "izvor": "C", "destinacija": "D", "tezina": 11 }, { "izvor": "C", "destinacija": "F", "tezina": 2 }, { "izvor": "D", "destinacija": "E", "tezina": 6 }, { "izvor": "E", "destinacija": "F", "tezina": 9 } ], "od": "A", "do": "E" }'
// '{ "graf": [ { "izvor": "X", "destinacija": "A", "tezina": 9 }, { "izvor": "X", "destinacija": "B", "tezina": 14 }, { "izvor": "X", "destinacija": "C", "tezina": 15 }, { "izvor": "A", "destinacija": "B", "tezina": 24 }, { "izvor": "A", "destinacija": "D", "tezina": 18 }, { "izvor": "B", "destinacija": "C", "tezina": 3 }, { "izvor": "B", "destinacija": "D", "tezina": 19 }, { "izvor": "C", "destinacija": "D", "tezina": 6 }, { "izvor": "C", "destinacija": "E", "tezina": 30 }, { "izvor": "D", "destinacija": "E", "tezina": 15 } ], "od": "X", "do": "E" }'

// primer zahteva
//   curl -X POST 'localhost:3000/api/path'  --data '{ "graf": [ { "izvor": "A", "destinacija": "B", "tezina": 2 }, { "izvor": "A", "destinacija": "D", "tezina": 1 }, { "izvor": "A", "destinacija": "C", "tezina": 5 }, { "izvor": "B", "destinacija": "C", "tezina": 3 }, { "izvor": "B", "destinacija": "D", "tezina": 2 }, { "izvor": "D", "destinacija": "E", "tezina": 1 }, { "izvor": "D", "destinacija": "C", "tezina": 3 }, { "izvor": "E", "destinacija": "C", "tezina": 1 } , { "izvor": "E", "destinacija": "F", "tezina": 2 } , { "izvor": "C", "destinacija": "F", "tezina": 5 } ], "od": "A", "do": "C" }'

// main funkcija koja slusa na portu 3000 i registruje rutu '/api/path' za PathHandler
func main() {
	var port = 3000
	http.Handle("/api/path", http.HandlerFunc(PathHandler))
	strPort := ":" + strconv.Itoa(port)
	fmt.Printf("server slusa na portu: %v", port)
	if err := http.ListenAndServe(strPort, nil); err != nil {
		log.Fatal(err)
	}
}

// PathHandler uzima grafove i težinu temena kao ulaz i generiše cenu i najkraći put
func PathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var reqobj InputGraph
		var resp *APIResponse
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			graph := CreateGraph(reqobj)
			resp = GetShortestPath(reqobj.From, reqobj.To, graph)
			w.WriteHeader(http.StatusOK)
			byteresp, _ := json.Marshal(resp)
			w.Write(byteresp)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := APIResponse{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}
