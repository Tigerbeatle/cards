package main

import (
	"github.com/tigerbeatle/cards/models"
	"github.com/justinas/alice"
	"github.com/tigerbeatle/cards/middleware"
	"github.com/tigerbeatle/cards/routes"
	controller "github.com/tigerbeatle/cards/controllers"
	"net/http"
	"log"
)

func main() {
	
	db := models.NewMongoDB()
/*
	collection := db.Collection("test")

	res, err := collection.InsertOne(context.Background(), map[string]string{"hello000ttt": "world Tiger"})
	if err != nil { log.Fatal(err) }
	id := res.InsertedID

log.Println("id:",id)

*/

	// Lets set some routes

	commonHandlers := alice.New(middleware.RecoverHandler, middleware.AcceptHandler)
	router := routes.NewRouter()

	appH := controller.HomeContext{db.Database}

	router.Get("/", commonHandlers.ThenFunc(appH.HomeHandler))


	log.Println("API Starting...")

	http.ListenAndServe(":8001", router)



/*

	commonHandlers := alice.New(context.ClearHandler, middleware.LoggingHandler, middleware.RecoverHandler, middleware.AcceptHandler)
	router := routes.NewRouter()

	appA := controller.AccountContext{session.DB("scorpion")}
	appH := controller.HomeContext{session.DB("scorpion")}
	appU := controller.UserContext{session.DB("scorpion")}
	//appQ := controller.QuestionContext{session.DB("scorpion")}
	appAct := controller.ActivityContext{session.DB("scorpion")}
	//appS := controller.UtilityContext{session.DB("scorpion")}
	appGeo := controller.GeoContext{redisPool}
	// Home

	router.Get("/", commonHandlers.ThenFunc(appH.HomeHandler))


	//router.Post("/api/1.0/accounts/register", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.User{})).ThenFunc(appA.CreateUser))
	//router.Post("/api/1.0/accounts/login", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.User{})).ThenFunc(appA.Login))
	router.Get("/api/1.0/accounts/user", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appA.UserProfile))
	router.Post("/api/1.0/accounts/updateUser", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.User{})).ThenFunc(appU.UpdateUser))
	router.Get("/api/1.0/accounts/createProfile", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appA.CreateUserProfile))


	// toDo Create a Log out function so we can track time

	router.Get("/api/1.0/interview/unanswered", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appU.Unanswered))
	router.Put("/api/1.0/interview/answered", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.QuestionSubmitted{})).ThenFunc(appU.Answered))

	router.Get("/api/1.0/activities/userActivities", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appAct.UserActivities))
	router.Put("/api/1.0/activities/record", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.SubmittedValues{})).ThenFunc(appAct.RecordActivityList))

	router.Post("/api/1.0/accounts/postTest", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.TestStruct{})).ThenFunc(appU.MatchTest))

	router.Post("/api/1.0/accounts/postTestGeo", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.TestStruct{})).ThenFunc(appGeo.GeoTestHandler))


	// Users

	//router.Get("/api/users/:id", commonHandlers.ThenFunc(appU.UserHandler))
	//router.Put("/api/users/:id", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.UserResource{})).ThenFunc(appU.UpdateUserHandler))
	//router.Delete("/api/users/:id", commonHandlers.ThenFunc(appU.DeleteUserHandler))


	// Questions

	//router.Get("/api/questions/:id", commonHandlers.ThenFunc(appQ.QuestionHandler))
	//router.Put("/api/questions/:id", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.QuestionResource{})).ThenFunc(appQ.UpdateQuestionHandler))
	//router.Delete("/api/questions/:id", commonHandlers.ThenFunc(appQ.DeleteQuestionHandler))
	//router.Get("/api/questions", commonHandlers.ThenFunc(appQ.QuestionsHandler))
	//router.Post("/api/questions", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.QuestionResource{})).ThenFunc(appQ.CreateQuestionHandler))


	// Scorpion Utilities


	//router.Get("/api/geocoding", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appS.GeoGeocodingHandler))

	//router.Post("/api/geocoding", commonHandlers.Append(middleware.AuthorizationHandler).ThenFunc(appS.GeoGeocodingHandler))
	//router.Post("/api/genTestBed", commonHandlers.ThenFunc(appS.GenTestBedHandler))

	http.ListenAndServe(":8001", router)

*/
}





