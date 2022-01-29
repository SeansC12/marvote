/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/SeansC12/marvote/pkg/api/v1/character"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setupRoutes(cmd *cobra.Command, args []string) {
	characterRepository := repository.NewCharacterRepository()

	characterService := service.NewCharacterService(characterRepository)
	characterRoutes := character.NewCharacterRoutes(characterService)
	e := echo.New()

	e.GET("/api/v1/characters/all", characterRoutes.GetAllCharacters)
	e.GET("/api/v1/character/:id", characterRoutes.Get)
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		fmt.Println(c.Path(), c.QueryParams(), err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Logger.Fatal(e.Start(":1323"))
}

// Never gonna give you up, never gonna let you down, never gonna run around and desert you.
