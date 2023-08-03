/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GlobalWebIndex/data-services-pair-challenge/domain"
	"github.com/GlobalWebIndex/data-services-pair-challenge/infrastructure"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts HTTP CRUD API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	err = dbpool.Ping(rootCmd.Context())
	if err != nil {
		log.Fatalf("Database check failed: %+v", err)
	}

	// Initialize the AudienceRepository with the database connection.
	audienceRepository := infrastructure.NewAudienceRepository(dbpool)

	// Initialize the AudienceService with the AudienceRepository.
	audienceService := domain.NewAudienceService(audienceRepository)

	// Create the CreateAudienceHandler with the AudienceService.
	createAudienceHandler := infrastructure.NewCreateAudienceHandler(audienceService)
	getAudienceByIDHandler := infrastructure.NewGetAudienceByIDHandler(audienceService)

	// Create the AudienceRouter with the CreateAudienceHandler and initialize the router.
	audienceRouter := infrastructure.NewAudienceRouter(createAudienceHandler, getAudienceByIDHandler)
	audienceRouter.SetupRoutes()

	// Start the HTTP server using the initialized router.
	log.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %+v", err)
	}
}
