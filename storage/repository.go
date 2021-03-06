package storage

import (
	"encoding/json"
	"kont/internal/repository"
)

func (s *Store) GetAllRepositories() []repository.Repository {

	var repositories = make([]repository.Repository, 0)

	_ = Storage.ForEach(func(k, v []byte) error {
		var repo repository.Repository
		_ = json.Unmarshal(v, &repo)

		repositories = append(repositories, repo)
		return nil
	})

	return repositories
}
