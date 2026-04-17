package app

import (
	"context"
	"log"
	"mediasoft/lesson9/internal/app/service"
	"mediasoft/lesson9/internal/bootstrap"
	"mediasoft/lesson9/internal/config"
	"mediasoft/lesson9/internal/repository/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}
	dbWrapper := database.NewDatabase(db)
	personService := service.NewPersonService(dbWrapper)
	groupService := service.NewGroupService(dbWrapper)

	router := http.NewServeMux()
	router.HandleFunc("POST /persons", personService.Create)        //создание человека
	router.HandleFunc("GET /persons/{id}", personService.Get)       //получение человека по индификатору
	router.HandleFunc("GET /persons", personService.GetAllPersons)  //получение списка всех людей
	router.HandleFunc("PUT /persons/{id}", personService.Update)    //редактирование человека
	router.HandleFunc("DELETE /persons/{id}", personService.Delete) //удаление человека

	router.HandleFunc("POST /groups", groupService.Create)                           //создание группы
	router.HandleFunc("GET /groups", groupService.GetAllGroups)                      //получение списка групп
	router.HandleFunc("GET /groups/{id}", groupService.Get)                          //получение группы по индификатору
	router.HandleFunc("PUT /groups/{id}", groupService.Update)                       //редактирование группы
	router.HandleFunc("DELETE /groups/{id}", groupService.Delete)                    //удаление группы
	router.HandleFunc("GET /groupspersons/{id}", groupService.ListPersonLocal)       //получение списка посльзователей группы по индификатору
	router.HandleFunc("GET /allgroupspersons/{id}", groupService.ListPersonAll)      //получение списка посльзователей группы по индификатору (включая дочерние)
	router.HandleFunc("GET /countgroupspersons/{id}", groupService.CountGroupLocal)  //получение количества человек в группе по индификатору
	router.HandleFunc("GET /countallgroupspersons/{id}", groupService.CountGroupAll) //получение количества человек в группе по индификатору (включая дочерние)

	srv := http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("run server: http://localhost%s", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("error when listen and serve", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	sig := <-ch
	log.Printf("received signal: %s", sig)
	return srv.Shutdown(context.Background())
}
