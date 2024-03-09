package user_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/internal/repo"
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

func (r *Repository) GetByID(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByID(ctx, pgtype.UUID{
		Bytes: userID,
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUser: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByLogin(ctx context.Context, tx pgx.Tx, login string) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUserByLogin: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByPhone(ctx context.Context, tx pgx.Tx, phone domain.Phone) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByPhone(ctx, repo.ConvertToNullStr(string(phone)))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUserByPhone: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByEmail(ctx context.Context, tx pgx.Tx, email domain.Email) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByEmail(ctx, repo.ConvertToNullStr(string(email)))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUserByEmail: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetByAccount(ctx context.Context, tx pgx.Tx, accountID int64) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUserByAccount: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) Save(ctx context.Context, tx pgx.Tx, u domain.User) (domain.User, error) {
	qtx := r.queries(tx)

	user, err := qtx.SaveUser(ctx, sqlc.SaveUserParams{
		Login:        u.Login,
		Phone:        repo.ConvertToNullStr(string(u.Phone)),
		PasswordHash: u.PasswordHash,
		Roles:        u.Roles.ConvString(),
	})
	if err != nil {
		if err = checkConstraints(err); err != nil {
			return domain.User{}, err
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("qtx.SaveUser: %w", err)
	}

	if err != nil {
		return domain.User{}, fmt.Errorf("transactional.Commit: %w", err)
	}
	return mapper.ConvUserToDomain(user), nil
}

func (r *Repository) Update(ctx context.Context, tx pgx.Tx, req dto.UserUpdateRequest) (domain.User, error) {
	qtx := r.queries(tx)

	name := repo.ConvertToNullStr(req.Name)
	login := repo.ConvertToNullStr(req.Login)
	email := repo.ConvertToNullStr(string(req.Email))
	phone := repo.ConvertToNullStr(string(req.Phone))
	emailVerified := repo.ConvertToNullBool(req.EmailVerified)
	roles := repo.ConvertToNullStr(req.Roles.ConvString())
	passwordHash := repo.ConvertToNullStr(req.PasswordHash)

	u, err := qtx.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID: pgtype.UUID{
			Bytes: req.ID,
			Valid: true,
		},
		Name:          name,
		PasswordHash:  passwordHash,
		Login:         login,
		Email:         email,
		Phone:         phone,
		Roles:         roles,
		EmailVerified: emailVerified,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		if err = checkConstraints(err); err != nil {
			return domain.User{}, err
		}
		return domain.User{}, fmt.Errorf("r.q.UpdateUser: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	qtx := r.queries(tx)

	_, err := qtx.RemoveUser(ctx, pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
	if err != nil {
		return fmt.Errorf("qtx.RemoveUser: %w", err)
	}
	return nil
}

func (r *Repository) Find(ctx context.Context, tx pgx.Tx, req dto.UserFindRequest) ([]domain.User, error) {
	qtx := r.queries(tx)

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

	return make([]domain.User, 0, 0), nil
}

func checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		switch e.ConstraintName {
		case ConstraintUsersLoginKey:
			return ErrUserLoginAlreadyExists
		case ConstraintUsersPhoneKey:
			return ErrUserPhoneAlreadyExists
		case ConstraintUsersEmailKey:
			return ErrUserEmailAlreadyExists
		default:
			return ErrUserAlreadyExists
		}
	}
	return nil
}

func (r *Repository) queries(tx pgx.Tx) *sqlc.Queries {
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}
	return qtx
}
