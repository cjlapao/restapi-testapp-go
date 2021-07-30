package repositories

import "github.com/cjlapao/restapi-testapp-go/database"

// Collection
const (
	ArticlesCollectionName = "articles"
	PeopleCollectionName   = "people"
	VehicleCollectionName  = "vehicles"
	StarshipCOllectionName = "starships"
	UsersCollectionName    = "users"
)

// Repository Entity
type Repository struct {
	Factory *database.MongoFactory
}

// NewRepo Creates a new Article Repository object
func NewRepo(factory *database.MongoFactory) Repository {
	result := Repository{
		Factory: factory,
	}

	return result
}
