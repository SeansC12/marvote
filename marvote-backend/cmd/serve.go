/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/SeansC12/marvote/pkg/api/v1/character"
	"github.com/SeansC12/marvote/pkg/infra"
	"github.com/SeansC12/marvote/pkg/logging"
	"github.com/SeansC12/marvote/pkg/repository"
	"github.com/SeansC12/marvote/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the backend service",
	Long:  `Starts the backend service`,
	Run:   setupRoutes,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func setupRoutes(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d", appConfig.MongoCfg.Username,
		appConfig.MongoCfg.Password,
		appConfig.MongoCfg.Host,
		appConfig.MongoCfg.Port)

	characterCollections, err := infra.GetCollection(ctx, mongoUri)
	if err != nil {
		logging.Fatalf("%v", err)
	}

	characterRepository := repository.NewCharacterRepository(ctx, characterCollections)

	characterService := service.NewCharacterService(characterRepository)
	characterRoutes := character.NewCharacterRoutes(characterService)
	e := echo.New()

	e.GET("/api/v1/characters/all", characterRoutes.GetAllCharacters)
	e.GET("/api/v1/character/:id", characterRoutes.Get)
	e.POST("/api/v1/character", characterRoutes.Save)
	e.DELETE("/api/v1/character/:id", characterRoutes.Delete)
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logging.Error(err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.ServerCfg.Port)))
}
