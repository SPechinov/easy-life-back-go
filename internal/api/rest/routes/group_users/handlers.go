package group_users

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/entities"
	"net/http"
)

func (controller *restGroupUsersController) handlerGetList(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")
	list, err := controller.useCases.GetList(
		ctx,
		entities.GroupGetUsersList{
			GroupID: groupID,
			UserID:  userID,
		},
	)
	if err != nil {
		return err
	}
	return echoCTX.JSON(http.StatusOK, list)
}

func (controller *restGroupUsersController) handlerInviteUserInGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *InviteUserDTO) error {
	groupID := echoCTX.Param("groupID")
	err := controller.useCases.Invite(
		ctx,
		entities.GroupInviteUser{
			GroupID:        groupID,
			UserID:         userID,
			InvitingUserID: dto.UserID,
		},
	)
	if err != nil {
		return err
	}
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupUsersController) handleExcludeUserInGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *ExcludeUserDTO) error {
	groupID := echoCTX.Param("groupID")
	err := controller.useCases.Exclude(
		ctx,
		entities.GroupExcludeUser{
			GroupID:         groupID,
			UserID:          userID,
			ExcludingUserID: dto.UserID,
		},
	)
	if err != nil {
		return err
	}
	return echoCTX.NoContent(http.StatusNoContent)
}
