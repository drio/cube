package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type Message struct {
	Msg string
}

func getPage() string {
	hostname, _ := os.Hostname()

	s := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>cube demo</title>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            font-family: Arial, sans-serif;
            background-color: #f2f2f2;
        }
        .container {
            text-align: center;
        }
        h1 {
            font-size: 50px;
            color: #333;
        }
        p {
            font-size: 20px;
            color: #666;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1> 👋 hello for %s</h1>
    </div>
</body>
</html>`, hostname)
	return s
}

func main() {
	r := chi.NewRouter()
	r.Post("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(getPage()))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check called")
		w.Write([]byte("OK"))
	})
	r.Get("/healthfail", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check failed")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:7777",
		Handler: r,
	}

	go func() {
		log.Println("Listening on http://localhost:7777")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Setup handler for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	<-c

	log.Println("Shutting down")
}
