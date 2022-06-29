package user_meili

import (
	"api/internal/entity"
	rep "api/internal/repository"
	"encoding/json"

	"github.com/meilisearch/meilisearch-go"
)

type repository struct {
	meili *meilisearch.Client
}

func New(meili *meilisearch.Client) rep.UserMeiliRepository {
	return &repository{meili: meili}
}

func (r *repository) UpdateUser(u *entity.User) error {
	index := r.meili.Index("users")
	var data map[string]interface{}
	json_data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(json_data, &data); err != nil {
		return err
	}
	_, err = index.UpdateDocuments(data)
	return err
}

// func (r *repository) GetUserByID(id int) (*entity.User, error) {
// 	index := r.meili.Index("users")

// 	res, err := index.Search("", &meilisearch.SearchRequest{
// 		Filter: fmt.Sprintf("id = %d", id),
// 		Limit:  1,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	b, err := json.Marshal(res.Hits[0])
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := entity.User{}
// 	err = json.Unmarshal(b, &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
