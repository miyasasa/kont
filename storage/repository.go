package storage

import (
	"encoding/json"
	"miya/internal/repository"
)

func (s *Store) GetAllRepositories() []repository.Repository {

	var repositories []repository.Repository

	_ = Storage.ForEach(func(k, v []byte) error {
		var repo repository.Repository
		_ = json.Unmarshal(v, &repo)

		repositories = append(repositories, repo)
		return nil
	})

	return repositories
}
