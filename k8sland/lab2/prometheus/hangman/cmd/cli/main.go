package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/k8sland/lab2/prom/hangman"
	"github.com/pkg/errors"
)

var g hangman.Game
var base string

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Printf("\n\n 🙅‍♀️ No way!! Seriously??\n\n")
		os.Exit(0)
	}()

	flag.StringVar(&base, "hm", "localhost:5000", "Specify a Hangman service host:port")
	flag.Parse()

	err := call(http.DefaultClient, "GET", urlFor("new_game"), nil, &g)
	if err != nil {
		panic(err)
	}

	old := g.Tally.TurnsLeft
	good := true
	for !gameOver() {
		print("\033[H\033[2J")
		fmt.Printf("\nYour Word: %10s\n", string(g.Tally.Letters))
		c := prompt(good)
		if c != '\n' {
			issueGuess(c)
		}
		if old != g.Tally.TurnsLeft {
			good = false
			old = g.Tally.TurnsLeft
		} else {
			good = true
		}
	}
}

func gameOver() bool {
	if g.Tally.Status == hangman.Won {
		fmt.Printf("\n👏  Noace!! You've just won\n\n")
		return true
	}

	if g.Tally.Status == hangman.Lost {
		fmt.Printf("\n😿  Meow! You've just lost. It was `%s\n\n", g.Word)
		return true
	}
	return false
}

func urlFor(path string) string {
	return "http://" + base + "/" + path
}

func issueGuess(guess rune) {
	body := struct {
		Game  *hangman.Game `json:"game"`
		Guess rune          `json:"guess"`
	}{
		Game:  &g,
		Guess: guess,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = call(http.DefaultClient, "POST", urlFor("guess"), bytes.NewReader(payload), &g)
	if err != nil {
		panic(err)
	}
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

func prompt(s bool) rune {
	reader := bufio.NewReader(os.Stdin)

	ic := "😃"
	if !s {
		ic = "😡"
	}
	fmt.Printf("\n%s %10s [%d/%d]? ", ic, "Your Guess", g.Tally.TurnsLeft, hangman.MaxGuesses)
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	return char
}
