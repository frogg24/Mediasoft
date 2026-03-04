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
	router.HandleFunc("POST /persons", personService.Create)        //работает
	router.HandleFunc("GET /persons/{id}", personService.Get)       //работает
	router.HandleFunc("GET /persons", personService.GetAllPersons)  //работает
	router.HandleFunc("PUT /persons/{id}", personService.Update)    //работает
	router.HandleFunc("DELETE /persons/{id}", personService.Delete) //работает

	router.HandleFunc("POST /groups", groupService.Create)                           //работает
	router.HandleFunc("GET /groups", groupService.GetAllGroups)                      //работает
	router.HandleFunc("GET /groups/{id}", groupService.Get)                          //работает
	router.HandleFunc("PUT /groups/{id}", groupService.Update)                       //работает
	router.HandleFunc("DELETE /groups/{id}", groupService.Delete)                    //работает
	router.HandleFunc("GET /groupspersons/{id}", groupService.ListPersonLocal)       //работает
	router.HandleFunc("GET /allgroupspersons/{id}", groupService.ListPersonAll)      //работает
	router.HandleFunc("GET /countgroupspersons/{id}", groupService.CountGroupLocal)  //работает
	router.HandleFunc("GET /countallgroupspersons/{id}", groupService.CountGroupAll) //работает

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
