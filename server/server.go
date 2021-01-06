package server

import (
	"context"
	"github.com/FerdinaKusumah/fastcrud/route"
	"github.com/FerdinaKusumah/fastcrud/utils"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	var wait = time.Second * time.Duration(utils.ReadWriteTimeout)
	var address = fmt.Sprintf(`%s:%s`, viper.GetString("app.host"), viper.GetString("app.port"))
	var r = route.Router()

	srv := &http.Server{
		Addr: address,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(utils.ReadWriteTimeout),
		ReadTimeout:  time.Second * time.Duration(utils.ReadWriteTimeout),
		IdleTimeout:  time.Second * time.Duration(utils.IdleTimeout),
		Handler:      http.TimeoutHandler(r, time.Duration(utils.IdleTimeout)*time.Second, "Request Timeout !\n"), // Pass our instance of gorilla/mux in.
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			_ = utils.SendLogError("unable to start the server: %s", err)
		}
	}()

	fmt.Println(fmt.Sprintf("Server Running on %s...", address))
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	fmt.Println("shutting down")
	os.Exit(0)

}
