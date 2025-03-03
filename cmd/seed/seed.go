package main

import (
	"context"
	"github.com/google/uuid"
	"komek/config"
	"komek/internal/domain/client/entity"
	clientRepo "komek/internal/domain/client/repository"
	clientSrv "komek/internal/domain/client/service"
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
	service := clientSrv.New(clientsRepository)
	ctx := context.Background()

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

	animalCategoryID, err := uuid.Parse("cf4001c6-c20f-4933-9234-c7c206642f3a")
	if err != nil {
		l.Error("app - Run - uuid.Parse animalCategoryID:", logger.Err(err))
	}
	oldCategoryID, err := uuid.Parse("a082da62-0076-429a-b2f3-4ecbb4478ced")
	if err != nil {
		l.Error("app - Run - uuid.Parse oldCategoryID:", logger.Err(err))
	}
	creates := []entity.CreateIn{
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
	}

	for _, create := range creates {
		_, err = service.Create(ctx, create)
		if err != nil {
			l.Error("app - Run - service.Create:", logger.Err(err))
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
