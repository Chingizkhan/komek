package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/user/entity"
	"komek/internal/errs"
	"komek/internal/mapper"
	"komek/internal/service/transactional"
	"komek/pkg/null_value"
	"komek/pkg/postgres"
	"log"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) GetByID(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.GetUserByID(ctx, null_value.UUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUser: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.GetUserByLogin(ctx, null_value.String(login))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUserByLogin: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByPhone(ctx context.Context, phone phone.Phone) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.GetUserByPhone(ctx, null_value.String(string(phone)))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUserByPhone: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByEmail(ctx context.Context, email email.Email) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.GetUserByEmail(ctx, null_value.String(string(email)))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUserByEmail: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByAccount(ctx context.Context, accountID int64) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.GetUserByAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUserByAccount: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) Save(ctx context.Context, u entity.User) (entity.User, error) {
	qtx := r.queries(ctx)

	user, err := qtx.SaveUser(ctx, sqlc.SaveUserParams{
		Login:        null_value.String(u.Login),
		Phone:        null_value.String(string(u.Phone)),
		PasswordHash: u.PasswordHash,
		Roles:        u.Roles.ToString(),
	})
	if err != nil {
		if err = checkConstraints(err); err != nil {
			return entity.User{}, err
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("qtx.SaveUser: %w", err)
	}

	return mapper.ConvUserToDomain(user), nil
}

func (r *Repository) Update(ctx context.Context, req entity.UpdateIn) (entity.User, error) {
	qtx := r.queries(ctx)

	u, err := qtx.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:            null_value.UUID(req.ID),
		Name:          null_value.String(req.Name),
		PasswordHash:  null_value.String(req.PasswordHash),
		Login:         null_value.String(req.Login),
		Email:         null_value.String(string(req.Email)),
		Phone:         null_value.String(string(req.Phone)),
		Roles:         null_value.String(req.Roles.ToString()),
		EmailVerified: null_value.Bool(req.EmailVerified),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		if err = checkConstraints(err); err != nil {
			return entity.User{}, err
		}
		return entity.User{}, fmt.Errorf("r.q.UpdateUser: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	qtx := r.queries(ctx)

	_, err := qtx.RemoveUser(ctx, null_value.UUID(id))
	if err != nil {
		return fmt.Errorf("qtx.RemoveUser: %w", err)
	}
	return nil
}

func (r *Repository) Find(ctx context.Context, req entity.FindRequest) ([]entity.User, error) {
	qtx := r.queries(ctx)

	users, err := qtx.FindUsers(ctx, sqlc.FindUsersParams{
		Name:  req.Name,
		Login: req.Login,
		Email: string(req.Email),
	})
	if err != nil {
		return nil, fmt.Errorf("qtx.FindUsers: %w", err)
	}

	for i, u := range users {
		log.Println(i, u)
	}

	return make([]entity.User, 0, 0), nil
}

func checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		switch e.ConstraintName {
		case ConstraintUsersLoginKey:
			return errs.ErrUserLoginAlreadyExists
		case ConstraintUsersPhoneKey:
			return errs.ErrUserPhoneAlreadyExists
		case ConstraintUsersEmailKey:
			return errs.ErrUserEmailAlreadyExists
		default:
			return errs.ErrUserAlreadyExists
		}
	}
	return nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
