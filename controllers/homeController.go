package controllers
import (
	"net/http"
	"github.com/tigerbeatle/cards/models"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/mongo"
)


type HomeContext struct {
	Db *mongo.Database

}

func (c *HomeContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"Cards-API", "Home", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}

func (c *HomeContext) LoginHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"Cards-API", "Login", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}