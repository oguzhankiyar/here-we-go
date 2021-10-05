package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexliesenfeld/health"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", os.TempDir() + "simple.sqlite")
	defer db.Close()

	checker := health.NewChecker(
		health.WithCacheDuration(1*time.Second),

		health.WithTimeout(10*time.Second),

		health.WithCheck(health.Check{
			Name:    "database",
			Timeout: 2 * time.Second,
			Check:   db.PingContext,
		}),

		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "search",
			Check: func(ctx context.Context) error {
				return fmt.Errorf("this makes the check fail")
			},
		}),

		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Println(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)

	http.Handle("/health", health.NewHandler(checker))
	log.Fatalln(http.ListenAndServe(":3000", nil))
}