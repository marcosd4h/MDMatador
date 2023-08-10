package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/marcosd4h/MDMatador/internal/database"
	"github.com/marcosd4h/MDMatador/internal/mdm"
	"github.com/marcosd4h/MDMatador/internal/version"

	"github.com/gorilla/sessions"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {

		//TODO: Revert this when needed
		//trace := string(debug.Stack())
		trace := ""
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL   string
	httpPort  int
	autoHTTPS struct {
		domain  string
		email   string
		staging bool
	}
	tls struct {
		certFile string
		keyFile  string
	}
	db struct {
		dsn         string
		automigrate bool
	}
	session struct {
		secretKey    string
		oldSecretKey string
	}
	debug struct {
		log  bool
		http bool
	}
}

type application struct {
	config          config
	db              *database.DB
	logger          *slog.Logger
	sessionStore    *sessions.CookieStore
	identityManager *mdm.WSTEPManager
	cmdManager      *mdm.CommandManager
	wg              sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	// Command line flags setup

	flag.StringVar(&cfg.baseURL, "base-url", "https://demomatador.io", "base URL for the application")
	//flag.StringVar(&cfg.baseURL, "base-url", "http://localhost", "base URL for the application")

	flag.IntVar(&cfg.httpPort, "http-port", 443, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.autoHTTPS.domain, "auto-https-domain", "", "domain to enable automatic HTTPS for")
	flag.StringVar(&cfg.autoHTTPS.email, "auto-https-email", "admin@github.com", "contact email address for problems with LetsEncrypt certificates")
	flag.BoolVar(&cfg.autoHTTPS.staging, "auto-https-staging", false, "use LetsEncrypt staging environment")

	//flag.StringVar(&cfg.tls.certFile, "custom-cert-file", "", "custom certificate to use by TLS server")
	//flag.StringVar(&cfg.tls.keyFile, "custom-key-file", "", "custom cert key to use by TLS server")
	flag.StringVar(&cfg.tls.certFile, "custom-cert-file", "", "certificate to use by TLS server")
	flag.StringVar(&cfg.tls.keyFile, "custom-key-file", "", "key to use by TLS server")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "db.sqlite", "sqlite3 DSN")
	flag.BoolVar(&cfg.db.automigrate, "db-automigrate", true, "run migrations on startup")
	flag.StringVar(&cfg.session.secretKey, "session-secret-key", "xknasfx3psbegn1aayy5as4grc15devu", "secret key for session cookie authentication")
	flag.StringVar(&cfg.session.oldSecretKey, "session-old-secret-key", "", "previous secret key for session cookie authentication")
	flag.BoolVar(&cfg.debug.log, "debug-log", true, "debug logging enabled for application")
	flag.BoolVar(&cfg.debug.http, "debug-http", true, "debug logging enabled for http traffic")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	// Version request handling
	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	// Logger setup - only log debug messages if debug.log is true
	if !cfg.debug.log {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	// Database connection and migration setup
	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	// MDM identity certificate and private key setup
	identityCert, dentityPrivateKey, err := db.GetIdentityCert()
	if err != nil {
		return err
	}

	certManager, err := mdm.NewCertManager(identityCert, dentityPrivateKey)
	if err != nil {
		return err
	}

	// MDM command manager setup
	cmdManager, err := mdm.GetCommandManager(cfg.baseURL, logger, db)
	if err != nil {
		return err
	}

	// Session store setup
	keyPairs := [][]byte{[]byte(cfg.session.secretKey), nil}
	if cfg.session.oldSecretKey != "" {
		keyPairs = append(keyPairs, []byte(cfg.session.oldSecretKey), nil)
	}

	sessionStore := sessions.NewCookieStore(keyPairs...)
	sessionStore.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   86400 * 7,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	// Application setup
	app := &application{
		config:          cfg,
		db:              db,
		logger:          logger,
		sessionStore:    sessionStore,
		identityManager: certManager,
		cmdManager:      cmdManager,
	}

	// Application Start
	if cfg.autoHTTPS.domain != "" {
		return app.serveAutoHTTPS()
	}

	if len(app.config.tls.certFile) > 0 && len(app.config.tls.keyFile) > 0 {
		return app.serveSelfSignedHTTPS(app.config.tls.certFile, app.config.tls.keyFile)
	}

	return app.serveHTTP()
}
