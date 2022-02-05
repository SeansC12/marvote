package character

import (
	"net/http"

	"github.com/SeansC12/marvote/pkg/logging"
	"github.com/SeansC12/marvote/pkg/model"
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
	ctx := c.Request().Context()
	logging.Info("retrieving all characters...")
	characters, err := cr.characterService.GetAll(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, characters)
}

func (cr *CharacterRoutes) Get(c echo.Context) error {
	idStr := c.Param("id")
	ctx := c.Request().Context()
	logging.Infof("retrieving a characters %s...", idStr)
	characters, err := cr.characterService.Get(ctx, idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, characters)
}

func (cr *CharacterRoutes) Save(c echo.Context) (err error) {
	ctx := c.Request().Context()
	mc := new(CharacterInfoDto)
	if err = c.Bind(mc); err != nil {
		return
	}
	// To avoid security flaws try to avoid passing bound structs directly to other methods
	// if these structs contain fields that should not be bindable.
	charInfo := model.CharacterInfo{
		Name: mc.Name,
		Aka:  mc.Aka,
	}
	x, err := cr.characterService.Save(ctx, charInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	logging.Debugf("Original Id %d", charInfo.Id)
	logging.Debugf("Returned Id: %d", x.Id)
	return c.JSON(http.StatusOK, x)
}

func (cr *CharacterRoutes) Delete(c echo.Context) error {
	idStr := c.Param("id")
	ctx := c.Request().Context()
	logging.Infof("deleting a character %s...", idStr)
	characters, err := cr.characterService.Delete(ctx, idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, characters)
}
