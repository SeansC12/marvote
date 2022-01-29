package character

import (
	"net/http"
	"strconv"

	"github.com/SeansC12/marvote/pkg/service"
	"github.com/labstack/echo/v4"
)

type CharacterRoutes struct {
	characterService service.ICharacterService
}

func NewCharacterRoutes(characterService service.ICharacterService) *CharacterRoutes {
	return &CharacterRoutes{
		characterService: characterService,
	}
}

func (cr *CharacterRoutes) GetAllCharacters(c echo.Context) error {

	characters, err := cr.characterService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, characters)
}

func (cr *CharacterRoutes) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	characters, err := cr.characterService.Get(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, characters)
}
