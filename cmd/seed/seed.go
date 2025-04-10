package main

import (
	"context"
	"github.com/google/uuid"
	"komek/config"
	accountRepo "komek/internal/domain/account/repository"
	accountSrv "komek/internal/domain/account/service"
	"komek/internal/domain/client/entity"
	clientRepo "komek/internal/domain/client/repository"
	clientSrv "komek/internal/domain/client/service"
	fundraise "komek/internal/domain/fundraise/entity"
	fundraise_repo "komek/internal/domain/fundraise/repository"
	fundraise_service "komek/internal/domain/fundraise/service"
	"komek/internal/service/transactional"
	"komek/internal/usecase/client"
	"komek/pkg/logger"
	"komek/pkg/postgres"
	"log"
	"os"
)

func main() {
	cfg, err := config.New("./config/config.yml")
	if err != nil {
		log.Fatalf("new config: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Debug("debug messages are enabled")

	pg, err := postgres.New(
		cfg.PG.DSN(),
		postgres.MaxPoolSize(cfg.PG.PoolMax),
	)
	if err != nil {
		l.Error("app - Run - postgres.New:", logger.Err(err))
		os.Exit(1)
	}
	defer pg.Close()

	clientsRepository := clientRepo.New(pg)
	accountRepository := accountRepo.New(pg)
	fundraiseRepository := fundraise_repo.New(pg)
	transactionService := transactional.New(pg)

	accountService := accountSrv.New(accountRepository)
	clientService := clientSrv.New(clientsRepository)
	fundraiseService := fundraise_service.New(fundraiseRepository)

	clientUseCase := client.New(clientService, fundraiseService, accountService, transactionService)

	ctx := context.Background()

	// create client categories

	//categories := []entity.Category{
	//	{Name: "Животные"},
	//	{Name: "Пожилые"},
	//}
	//
	//_, err = clientsRepository.SaveCategories(ctx, categories)
	//if err != nil {
	//	l.Error("app - Run - clients.SaveCategories:", logger.Err(err))
	//	return
	//}

	// // create fundraise types

	//if err = clientUseCase.CreateFundraiseType(ctx, "monthly"); err != nil {
	//	l.Error("app - Run - clientUseCase.CreateFundraiseType:", logger.Err(err))
	//	return
	//}

	monthlyFundraiseType, err := uuid.Parse("e4c61f9f-d249-459d-aea0-4b2206960fe8")
	if err != nil {
		l.Error("app - Run - uuid.Parse monthlyFundraiseType:", logger.Err(err))
		return
	}

	animalCategoryID, err := uuid.Parse("1e15f5c2-31e9-4967-b47a-0650bb9b8f62")
	if err != nil {
		l.Error("app - Run - uuid.Parse animalCategoryID:", logger.Err(err))
		return
	}
	oldCategoryID, err := uuid.Parse("daa9a412-e5df-4d25-9290-7e708a17bd93")
	if err != nil {
		l.Error("app - Run - uuid.Parse oldCategoryID:", logger.Err(err))
		return
	}
	creates := []entity.CreateIn{
		{
			Name:          "Талас (обычный көт)",
			Phone:         "87058113795",
			Email:         "talas@mail.ru",
			Age:           60,
			City:          "Kostanay",
			Address:       "Москва, ул. Ленина, д. 10",
			Description:   "Потерял работу из-за болезни, нуждается в поддержке.",
			Circumstances: "Татьяна Георгиевна живет в однокомнатной квартире и нуждается в помощи.",
			ImageURL:      "/talas.png",
			CategoryIDs:   []uuid.UUID{oldCategoryID},
		},
		{
			Name:          "Волосатый көт",
			Phone:         "",
			Email:         "",
			Age:           3,
			City:          "Almaty",
			Address:       "Санкт-Петербург, Невский пр., д. 25",
			Description:   "Работала учителем, воспитывала двоих детей...",
			Circumstances: "После пожара осталась без жилья и нуждается в помощи.",
			ImageURL:      "/kot.png",
			CategoryIDs:   []uuid.UUID{animalCategoryID},
		},
	}

	for _, create := range creates {
		cl, err := clientUseCase.CreateClient(ctx, create)
		if err != nil {
			l.Error("app - Run - clientUseCase.CreateClient:", logger.Err(err))
			return
		}

		_, err = clientUseCase.CreateFundraise(ctx, fundraise.CreateIn{
			Goal:      350000,
			Collected: 0,
			AccountID: cl.Account.ID,
			TypeID:    monthlyFundraiseType,
		})
		if err != nil {
			l.Error("app - Run - clientUseCase.CreateFundraise:", logger.Err(err))
			return
		}
	}

	// list clients
	clients, err := clientsRepository.List(ctx)
	if err != nil {
		l.Error("app - Run - clients.List:", logger.Err(err))
		return
	}

	log.Println("len(clients):", len(clients))

	for _, client := range clients {
		log.Printf("client - %+v", client)
	}
}
