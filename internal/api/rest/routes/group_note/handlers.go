package group_note

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/entities"
	"net/http"
)

func (controller *restGroupNoteController) GetList(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")
	list, err := controller.useCases.GetList(
		ctx,
		&entities.NoteGetList{
			UserID:  userID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, list)
}

func (controller *restGroupNoteController) Get(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")
	noteID := echoCTX.Param("noteID")

	list, err := controller.useCases.Get(
		ctx,
		&entities.NoteGet{
			ID:      noteID,
			UserID:  userID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, list)
}

func (controller *restGroupNoteController) Add(echoCTX echo.Context, ctx context.Context, userID string, dto *AddDTO) error {
	groupID := echoCTX.Param("groupID")

	err := controller.useCases.Add(
		ctx,
		&entities.NoteAdd{
			UserID:      userID,
			GroupID:     groupID,
			Title:       dto.Title,
			Description: dto.Description,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupNoteController) Patch(echoCTX echo.Context, ctx context.Context, userID string, dto *PatchDTO) error {
	groupID := echoCTX.Param("groupID")
	noteID := echoCTX.Param("noteID")

	err := controller.useCases.Patch(
		ctx,
		&entities.NotePatch{
			ID:          noteID,
			UserID:      userID,
			GroupID:     groupID,
			Title:       dto.Title,
			Description: dto.Description,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupNoteController) Delete(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")
	noteID := echoCTX.Param("noteID")

	err := controller.useCases.Delete(
		ctx,
		&entities.NoteDelete{
			ID:      noteID,
			UserID:  userID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}
