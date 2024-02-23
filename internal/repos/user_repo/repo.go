package user_repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	user_db "komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/pkg/postgres"
	"log"
	"strings"
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
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}

	u, err := qtx.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("r.q.GetUser :%w", err)
	}
	return domain.User{
		ID:            u.ID,
		Name:          u.Name.String,
		Phone:         domain.Phone(u.Phone.String),
		Login:         u.Login,
		EmailVerified: u.EmailVerified.Bool,
		PasswordHash:  u.PasswordHash,
		Email:         domain.Email(u.Email.String),
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

//
//func (r *Repository) GetWithTX(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error) {
//
//	u, err := qtx.GetUser(ctx, userID)
//	if err != nil {
//		return domain.User{}, fmt.Errorf("r.q.GetUser :%w", err)
//	}
//	return domain.User{
//		ID:            u.ID,
//		Name:          u.Name.String,
//		Phone:         domain.Phone(u.Phone.String),
//		Login:         u.Login,
//		EmailVerified: u.EmailVerified.Bool,
//		PasswordHash:  u.PasswordHash,
//		Email:         domain.Email(u.Email.String),
//		CreatedAt:     u.CreatedAt.Time,
//		UpdatedAt:     u.UpdatedAt.Time,
//	}, nil
//}

func (r *Repository) Save(ctx context.Context, tx pgx.Tx, u domain.User) error {
	qtx := r.q.WithTx(tx)

	_, err := qtx.SaveUser(ctx, user_db.SaveUserParams{
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
		Roles:        u.Roles.ConvString(),
	})
	if err != nil {
		return fmt.Errorf("qtx.SaveUser: %w", err)
	}

	if err != nil {
		return fmt.Errorf("transactional.Commit: %w", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, tx pgx.Tx, req dto.UserUpdateRequest) (domain.User, error) {
	qtx := r.q.WithTx(tx)

	name := checkAndConvertToNullStr(req.Name)
	login := checkAndConvertToNullStr(req.Login)
	email := checkAndConvertToNullStr(string(req.Email))
	phone := checkAndConvertToNullStr(string(req.Phone))
	passwordHash := checkAndConvertToNullStr(req.PasswordHash)
	emailVerified := checkAndConvertToNullBool(req.EmailVerified)

	u, err := qtx.UpdateUser(ctx, user_db.UpdateUserParams{
		ID:            req.ID,
		Name:          name,
		Login:         login,
		Email:         email,
		Phone:         phone,
		PasswordHash:  passwordHash,
		EmailVerified: emailVerified,
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("r.q.UpdateName: %w", err)
	}
	return domain.User{
		ID:            u.ID,
		Name:          u.Name.String,
		Phone:         domain.Phone(u.Phone.String),
		Login:         u.Login,
		Email:         domain.Email(u.Email.String),
		EmailVerified: u.EmailVerified.Bool,
		PasswordHash:  u.PasswordHash,
		Roles:         convertRolesToDomain(u.Roles),
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

func (r *Repository) Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	qtx := r.q.WithTx(tx)

	_, err := qtx.RemoveUser(ctx, id)
	if err != nil {
		return fmt.Errorf("qtx.RemoveUser: %w", err)
	}
	return nil
}

func (r *Repository) Find(ctx context.Context, tx pgx.Tx, req dto.UserFindRequest) ([]domain.User, error) {
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}

	users, err := qtx.FindUsers(ctx, user_db.FindUsersParams{
		Name:  "Jack",
		Login: "jake_buffalo",
		Email: "",
	})
	if err != nil {
		return nil, fmt.Errorf("qtx.FindUsers: %w", err)
	}

	for i, u := range users {
		log.Println(i, u)
	}

	return make([]domain.User, 0, 0), nil
}

func convertRolesToDomain(rolesStr string) domain.Roles {
	rolesInput := strings.Split(rolesStr, ",")
	roles := make(domain.Roles, 0, len(rolesInput))
	for _, r := range rolesInput {
		roles = append(roles, domain.Role(r))
	}
	return roles
}

func checkAndConvertToNullStr(value string) (nullValue sql.NullString) {
	if value != "" {
		nullValue = sql.NullString{
			String: value,
			Valid:  true,
		}
	}
	return
}

func checkAndConvertToNullBool(value bool) (nullValue sql.NullBool) {
	if value {
		nullValue = sql.NullBool{
			Bool:  value,
			Valid: true,
		}
	}
	return
}
