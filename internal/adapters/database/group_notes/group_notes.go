package group_notes

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"go-clean/pkg/postgres"
	"time"
)

type GroupNotes struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *GroupNotes {
	return &GroupNotes{
		postgres: postgres,
	}
}

func (g *GroupNotes) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	query := `
		SELECT
		    id,
		    title,
		    public.notes_info.description,
		    group_id,
		    user_creator_id,
		    user_updater_id,
		    created_at,
		    updated_at,
		    deleted_at
		FROM public.notes
		RIGHT JOIN public.notes_info ON public.notes_info.note_id = notes.id
		WHERE public.notes.group_id = $1 AND public.notes.deleted_at IS NULL
	`

	rows, err := g.postgres.Query(ctx, query, entity.GroupID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	notes := make([]entities.Note, 0)
	for rows.Next() {
		var note dataNote
		err = rows.Scan(
			&note.id,
			&note.title,
			&note.description,
			&note.groupID,
			&note.creatorID,
			&note.updaterID,
			&note.createdAt,
			&note.updatedAt,
			&note.deletedAt,
		)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		notes = append(
			notes,
			entities.Note{
				ID:    note.id,
				Title: note.title,
				Info: &entities.NoteInfo{
					Description: note.description,
				},
				GroupID:   note.groupID,
				CreatorID: note.creatorID,
				UpdaterID: note.updaterID,
				CreatedAt: note.createdAt.Format(time.RFC3339),
				UpdatedAt: note.updatedAt.Format(time.RFC3339),
				DeletedAt: helpers.GetPtrValueFromSQLNullTime(note.deletedAt, time.RFC3339),
			},
		)
	}

	return notes, nil
}
func (g *GroupNotes) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	query := `
		SELECT
		   id,
		   title,
		   public.notes_info.description,
		   group_id,
		   user_creator_id,
		   user_updater_id,
		   created_at,
		   updated_at,
		   deleted_at
		FROM public.notes
		RIGHT JOIN public.notes_info ON public.notes_info.note_id = public.notes.id
		WHERE 
		    public.notes.group_id = $1
			AND public.notes.id = $2
			AND public.notes.deleted_at IS NULL
	`

	var note dataNote
	err := g.postgres.QueryRow(ctx, query, entity.GroupID, entity.ID).Scan(
		&note.id,
		&note.title,
		&note.description,
		&note.groupID,
		&note.creatorID,
		&note.updaterID,
		&note.createdAt,
		&note.updatedAt,
		&note.deletedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, client_error.ErrNoteDeleted
		}
		logger.Error(ctx, err)
		return nil, err
	}
	return &entities.Note{
		ID:    note.id,
		Title: note.title,
		Info: &entities.NoteInfo{
			Description: note.description,
		},
		GroupID:   note.groupID,
		CreatorID: note.creatorID,
		UpdaterID: note.updaterID,
		CreatedAt: note.createdAt.Format(time.RFC3339),
		UpdatedAt: note.updatedAt.Format(time.RFC3339),
		DeletedAt: helpers.GetPtrValueFromSQLNullTime(note.deletedAt, time.RFC3339),
	}, nil
}

func (g *GroupNotes) Add(ctx context.Context, entity *entities.NoteAdd) (*entities.Note, error) {
	queryAdd := `
		INSERT INTO public.notes (user_creator_id, user_updater_id, group_id, title)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, group_id, user_creator_id, user_updater_id, created_at, updated_at
	`

	tx, err := g.postgres.Begin(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var note dataNote
	err = tx.QueryRow(ctx, queryAdd, entity.UserID, entity.UserID, entity.GroupID, entity.Title).Scan(
		&note.id,
		&note.title,
		&note.groupID,
		&note.creatorID,
		&note.updaterID,
		&note.createdAt,
		&note.updatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, client_error.ErrGroupNotExists
		}
		logger.Error(ctx, err)
		return nil, err
	}

	queryAddInfo := `
		INSERT INTO public.notes_info (note_id, description)
		VALUES ($1, $2)
		RETURNING description
	`

	err = tx.QueryRow(ctx, queryAddInfo, note.id, entity.Description).Scan(&note.description)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &entities.Note{
		ID:    note.id,
		Title: note.title,
		Info: &entities.NoteInfo{
			Description: note.description,
		},
		GroupID:   note.groupID,
		CreatorID: note.creatorID,
		UpdaterID: note.updaterID,
		CreatedAt: note.createdAt.Format(time.RFC3339),
		UpdatedAt: note.updatedAt.Format(time.RFC3339),
		DeletedAt: nil,
	}, nil
}

func (g *GroupNotes) Patch(ctx context.Context, entity *entities.NotePatch) error {
	err := g.patchTitle(ctx, entity)
	if err != nil {
		return err
	}

	err = g.patchInfo(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupNotes) patchTitle(ctx context.Context, entity *entities.NotePatch) error {
	if entity.Title == nil {
		return nil
	}

	query := `
		UPDATE public.notes SET title=$2 WHERE id=$1 AND deleted_at IS NULL
	`

	row, err := g.postgres.Exec(ctx, query, entity.ID, *entity.Title)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if row.RowsAffected() == 0 {
		return client_error.ErrNoteNotExists
	}

	return nil
}

func (g *GroupNotes) patchInfo(ctx context.Context, entity *entities.NotePatch) error {
	if entity.Description == nil {
		return nil
	}

	query := `
        WITH notes AS (
            SELECT id 
            FROM public.notes 
            WHERE id = $1 AND deleted_at IS NULL
        )
        UPDATE public.notes_info 
        SET description = $2 
        WHERE note_id IN (SELECT id FROM notes)
	`

	row, err := g.postgres.Exec(ctx, query, entity.ID, *entity.Description)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if row.RowsAffected() == 0 {
		return client_error.ErrNoteNotExists
	}

	return nil
}

func (g *GroupNotes) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	query := `
		UPDATE public.notes
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	row, err := g.postgres.Exec(ctx, query, entity.ID)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if row.RowsAffected() == 0 {
		return client_error.ErrNoteNotExists
	}
	return nil
}

func (g *GroupNotes) IsCreator(ctx context.Context, userID, noteID string) bool {
	query := `
		SELECT user_creator_id 
		FROM public.notes 
		WHERE 
		    id = $2 
		  	AND notes.user_creator_id = $1
		  	AND deleted_at IS NULL
	`

	cmt, err := g.postgres.Exec(ctx, query, userID, noteID)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return cmt.RowsAffected() > 0
}
