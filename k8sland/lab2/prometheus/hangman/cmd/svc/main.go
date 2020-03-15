package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/k8sland/lab2/prom/hangman"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port = ":5000"

var dicURL string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.StringVar(&dicURL, "d", "localhost:4000", "Specifies the host:port for the dictionary service")
	flag.Parse()

	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.Handle("/new_game", http.HandlerFunc(newGameHandler)).Methods("GET")
	r.Handle("/guess", http.HandlerFunc(guessHandler)).Methods("POST")
	r.Handle("/healthz", http.HandlerFunc(healthHandler)).Methods("GET")

	log.Printf("Handman listening on port %s...\n", port)
	log.Panic(http.ListenAndServe(port, m))
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Game  *hangman.Game `json:"game"`
		Guess rune          `json:"guess"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	req.Game.Guess(rune(req.Guess))
	raw, err := json.Marshal(req.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(raw))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	raw, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(raw))
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	g := hangman.NewGame(pick())

	buff, err := json.Marshal(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}

func urlFor(u, p string) string {
	return "http://" + u + "/" + p
}

func loadDictionary(u string) []string {
	var resp struct {
		Words []string `json:"words"`
		Size  int      `json:"size"`
	}

	if err := call(http.DefaultClient, "GET", urlFor(u, "words"), nil, &resp); err != nil {
		log.Panic(err)
	}

	return resp.Words
}

func pick() string {
	words := loadDictionary(dicURL)
	return words[rand.Intn(len(words))]
}

func call(c *http.Client, method, url string, payload io.Reader, res interface{}) error {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("boom: remote call %s crapped out!", url))
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("boom: url call `%s failed with code (%d)", url, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(&res)
}
