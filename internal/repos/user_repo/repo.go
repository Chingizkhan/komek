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
	user_db "komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/pkg/postgres"
	"log"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *user_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, user_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUser(ctx, pgtype.UUID{
		Bytes: userID,
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUser :%w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) GetUserByLogin(ctx context.Context, tx pgx.Tx, login string) (domain.User, error) {
	qtx := r.queries(tx)

	u, err := qtx.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("r.q.GetUserByLogin :%w", err)
	}
	return domain.User{
		ID:            u.ID.Bytes,
		Name:          u.Name.String,
		Phone:         domain.Phone(u.Phone.String),
		Login:         u.Login,
		EmailVerified: u.EmailVerified.Bool,
		PasswordHash:  u.PasswordHash,
		Email:         domain.Email(u.Email.String),
		CreatedAt:     u.CreatedAt.Time,
		UpdatedAt:     u.UpdatedAt.Time,
	}, nil
}

func (r *Repository) Save(ctx context.Context, tx pgx.Tx, u domain.User) (domain.User, error) {
	qtx := r.queries(tx)

	phone := checkAndConvertToNullStr(string(u.Phone))

	user, err := qtx.SaveUser(ctx, user_db.SaveUserParams{
		Login:        u.Login,
		Phone:        phone,
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

	name := checkAndConvertToNullStr(req.Name)
	login := checkAndConvertToNullStr(req.Login)
	email := checkAndConvertToNullStr(string(req.Email))
	phone := checkAndConvertToNullStr(string(req.Phone))
	emailVerified := checkAndConvertToNullBool(req.EmailVerified)
	roles := checkAndConvertToNullStr(req.Roles.ConvString())
	passwordHash := checkAndConvertToNullStr(req.PasswordHash)

	u, err := qtx.UpdateUser(ctx, user_db.UpdateUserParams{
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

	users, err := qtx.FindUsers(ctx, user_db.FindUsersParams{
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

func checkAndConvertToNullStr(value string) (nullValue pgtype.Text) {
	if value != "" {
		nullValue = pgtype.Text{
			String: value,
			Valid:  true,
		}
	}
	return
}

func checkAndConvertToNullBool(value *bool) (nullValue pgtype.Bool) {
	if value != nil {
		nullValue = pgtype.Bool{
			Bool:  *value,
			Valid: true,
		}
	}
	return
}

func (r *Repository) queries(tx pgx.Tx) *user_db.Queries {
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}
	return qtx
}
