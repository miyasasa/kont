package remoterepository

import (
	"kont/internal/remoterepository/bitbucket"
	"kont/internal/repository"
	"kont/storage"
)

func ListenRemoteRepositories() {
	repositories := storage.Storage.GetAllRepositories()

	if len(repositories) != 0 {
		listenRepo(&repositories[0])
	}
}

func listenRepo(repo *repository.Repository) {
	if repo.Provider == repository.BITBUCKET {
		bitbucket.Listen(repo)
	}
}
