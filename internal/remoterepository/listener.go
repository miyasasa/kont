package remoterepository

import (
	"miya/internal/remoterepository/bitbucket"
	"miya/internal/repository"
	"miya/storage"
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
