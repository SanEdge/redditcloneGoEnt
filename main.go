package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/renaldyhidayatt/redditgoent/routes"
	"github.com/renaldyhidayatt/redditgoent/utils"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	err := utils.Viper()

	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := utils.Database(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	routes.NewAuthRoutes("/auth", db, r, ctx)
	routes.NewSubRedditRoutes("/subreddit", db, r, ctx)
	routes.NewPostRoutes("/post", db, r, ctx)
	routes.NewCommentRoutes("/comment", db, r, ctx)

	routes.NewTestRoutes("/test", r)

	serve := &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("PORT")),
		ReadTimeout:    time.Duration(time.Second) * 60,
		WriteTimeout:   time.Duration(time.Second) * 30,
		IdleTimeout:    time.Duration(time.Second) * 120,
		MaxHeaderBytes: 3145728,
		Handler:        r,
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Connected to port:", viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)

}
