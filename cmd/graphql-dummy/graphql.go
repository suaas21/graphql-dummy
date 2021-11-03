package main

import (
	"github.com/arangodb/go-driver"
	"github.com/suaas21/graphql-dummy/infra/sentry"
	"github.com/suaas21/graphql-dummy/repo/author"
	"github.com/suaas21/graphql-dummy/repo/book"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"github.com/suaas21/graphql-dummy/api"
	"github.com/suaas21/graphql-dummy/config"
	infraArango "github.com/suaas21/graphql-dummy/infra/arango"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/service"
	"golang.org/x/net/context"
)

// srvCmd is the serve sub command to start the api server
var srvCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve serves the api server",
	RunE:  serve,
}

func init() {
	srvCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "config.yaml", "config file path")
}

func serve(cmd *cobra.Command, args []string) error {
	cfgApp := config.GetApp(cfgPath)
	cfgArango := config.GetArango(cfgPath)
	cfgSentry := config.GetSentry(cfgPath)

	ctx := context.Background()

	lgr := logger.DefaultOutStructLogger

	db, err := infraArango.NewArangoDB(ctx, cfgArango)
	if err != nil {
		return err
	}

	client := infraArango.NewClient(db)

	bookRepo := book.NewArangoBookRepository(client, lgr)
	authorRepo := author.NewArangoAuthorRepository(client, lgr)

	bookAuthorService := service.NewBookAuthor(bookRepo, authorRepo, lgr)

	err = sentry.NewInit(cfgSentry.URL)
	if err != nil {
		return err
	}

	api.SetLogger(logger.DefaultOutLogger)

	errChan := make(chan error)

	go func() {
		if err := startApiServer(cfgApp, bookAuthorService, lgr); err != nil {
			errChan <- err
		}
	}()
	return <-errChan

}

func ensureDBCollectionForSchema(ctx context.Context, db driver.Database, lgr logger.StructLogger) (*service.BookAuthor, error) {
	client := infraArango.NewClient(db)

	bookRepo := book.NewArangoBookRepository(client, lgr)
	authorRepo := author.NewArangoAuthorRepository(client, lgr)

	return service.NewBookAuthor(bookRepo, authorRepo, lgr), nil
}

func startApiServer(cfg *config.Application, bookAuthorService *service.BookAuthor, lgr logger.StructLogger) error {
	baCtrl := api.NewBookAuthorController(*bookAuthorService, lgr)

	r := chi.NewMux()
	r.Mount("/api/v1/public", api.NewRouter(baCtrl))

	srvr := http.Server{
		Addr:    getAddressFromHostAndPort(cfg.Host, cfg.Port),
		Handler: r,
		//ErrorLog: logger.DefaultErrLogger,
		//WriteTimeout: cfg.WriteTimeout,
		//ReadTimeout:  cfg.ReadTimeout,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return ManageServer(&srvr, 30*time.Second)
}

func ManageServer(srvr *http.Server, gracePeriod time.Duration) error {
	errCh := make(chan error)

	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt}

	graceful := func() error {
		log.Println("Sut down server gracefully with in", gracePeriod)
		log.Println("To shutdown immediately press again")

		ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()

		return srvr.Shutdown(ctx)
	}

	forced := func() error {
		log.Println("Shutting down server forcefully")
		return srvr.Close()
	}

	go func() {
		log.Println("Starting server on", srvr.Addr)
		if err := srvr.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	go func() {
		errCh <- HandleSignals(sigs, graceful, forced)
	}()

	return <-errCh
}

// HandleSignals listen on the registered signals and fires the gracefulHandler for the
// first signal and the forceHandler (if any) for the next this function blocks and
// return any error that returned by any of the handlers first
func HandleSignals(sigs []os.Signal, gracefulHandler, forceHandler func() error) error {
	sigCh := make(chan os.Signal)
	errCh := make(chan error, 1)

	signal.Notify(sigCh, sigs...)
	defer signal.Stop(sigCh)

	grace := true
	for {
		select {
		case err := <-errCh:
			return err
		case <-sigCh:
			if grace {
				grace = false
				go func() {
					errCh <- gracefulHandler()
				}()
			} else if forceHandler != nil {
				err := forceHandler()
				errCh <- err
			}
		}
	}
}

func getAddressFromHostAndPort(host string, port int) string {
	addr := host
	if port != 0 {
		addr = addr + ":" + strconv.Itoa(port)
	}
	return addr
}
