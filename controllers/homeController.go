package controllers
import (
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/tigerbeatle/scorpionApi/models"
	"encoding/json"
)


type HomeContext struct {
	Db *mgo.Database
}

func (c *HomeContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"Scorpion-API", "Home", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}

func (c *HomeContext) LoginHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"Scorpion-API", "Login", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}