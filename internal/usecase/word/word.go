package word

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"komek/internal/domain/word"
)

type IRepository interface {
	Get(ctx context.Context, value string, userID uuid.UUID) (word.Word, error)
	Save(ctx context.Context, word word.Word) error
	Update(ctx context.Context, pkey string, word word.Word) (word.Word, error)
	Delete(ctx context.Context, value string, userID uuid.UUID) error
}

type UseCase struct {
	r IRepository
}

func New(r IRepository) *UseCase {
	return &UseCase{r}
}

func (u *UseCase) Get(ctx context.Context, value string, userId uuid.UUID) (word.Word, error) {
	w, err := u.r.Get(ctx, value, userId)
	if err != nil {
		return word.Word{}, err
	}
	return w, nil
}

func (u *UseCase) Save(ctx context.Context, word word.Word) error {
	err := u.r.Save(ctx, word)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) Update(ctx context.Context, oldValue string, w word.Word) (word.Word, error) {
	w, err := u.r.Update(ctx, oldValue, w)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return word.Word{}, ErrNothingUpdated
		}
		return word.Word{}, err
	}
	return w, nil
}

func (u *UseCase) Delete(ctx context.Context, value string, userId uuid.UUID) error {
	err := u.r.Delete(ctx, value, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNothingDeleted
		}
		return err
	}
	return nil
}
