package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/eiannone/keyboard"
)

type GameStatus struct {
	Paused bool
	Mu     sync.Mutex
}

var game GameStatus

func main() {
	port := 12372
	keyToPause := 'p'

	http.HandleFunc("/pause", pauseHandler)

	ip := getIP()
	if ip == nil {
		log.Fatal("it was not possible to get the local ip")
	}
	fmt.Printf("Access http://%s:%d/pause from any device in your network to pause/resume the game (just refresh the browser to pause/resume)\n", ip.String(), port)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			log.Fatal(err)
		}
	}()

	keyboardWatch(keyToPause)
}

func keyboardWatch(keyToPause rune) {
	keysEvents, err := keyboard.GetKeys(1)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Printf("Press %q to pause the game (this window must be selected, this requires alt+tab while playing)\n", keyToPause)
	fmt.Println("Press ESC to quit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Rune == keyToPause {
			err := pauseResumeGame()
			if err != nil {
				log.Fatal(err)
			}
		}

		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

func getIP() *net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return &ipv4
		}
	}
	return nil
}

func pauseHandler(w http.ResponseWriter, r *http.Request) {
	err := pauseResumeGame()
	if err != nil {
		fmt.Fprint(w, "could not update game status, please wait until the game responds")
		return
	}
	msg := "game paused"
	if !game.Paused {
		msg = "game resumed"
	}
	fmt.Fprint(w, msg)
}

func pauseResumeGame() error {
	locked := game.Mu.TryLock()
	if !locked {
		return errors.New("could not acquire lock")
	}
	defer game.Mu.Unlock()

	app := "HogwartsLegacy"

	args := []string{"-NoProfile", "-NonInteractive", "pssuspend"}
	if game.Paused {
		args = append(args, "-r")
		fmt.Println("resuming game")
	} else {
		fmt.Println("pausing game")
	}
	args = append(args, app)

	cmd := exec.Command("powershell.exe", args...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	game.Paused = !game.Paused

	return nil
}
