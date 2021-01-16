package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	_ "github.com/go-sql-driver/mysql"

	grpc "github.com/fboccacini/goauth/pkg/protocol/grpc"
	v1 "github.com/fboccacini/goauth/pkg/service/v1"
)

type Config struct {
	GRPCPort            string
	DatastoreDBHost     string
	DatastoreDBUser     string
	DatastoreDBPassword string
	DatastoreDBSchema   string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.MigrationsFolder, "migrations-folder", "", "migrations folder")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: %s", cfg.GRPCPort)
	}

	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	log.Println("Running migrations")

	// Read migrations directory
	files, err := ioutil.ReadDir(cfg.MigrationsFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Reverse file order
	sort.Slice(files, func(i, j int) bool { return files[i].Name() > files[j].Name() })

	for _, file := range files {
		log.Print(file.Name() + "...          ")

		// Extract query
		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", cfg.MigrationsFolder, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// Execute query
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Done.")
	}

	v1API := v1.NewGoAuthServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}

func (r *postgresRepository) RunMigrations() error {
	log.Println("Running migrations")

	// Read migrations directory
	files, err := ioutil.ReadDir(r.migrationsFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Reverse file order
	sort.Slice(files, func(i, j int) bool { return files[i].Name() > files[j].Name() })

	for _, file := range files {
		log.Print(file.Name() + "...          ")

		// Extract query
		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", r.migrationsFolder, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// Execute query
		_, err = r.db.Exec(string(content))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Done.")
	}

	log.Println("Migrations complete.")
	return err
}
