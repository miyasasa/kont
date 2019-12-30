package remoterepository

import (
	"kont/internal/remoterepository/bitbucket"
	"kont/internal/repository"
	"kont/storage"
)

func ListenRemoteRepositories() {
	repositories := storage.Storage.GetAllRepositories()

	for i := range repositories {
		go listenRepo(&repositories[i])
	}

}

func listenRepo(repo *repository.Repository) {
	if repo.Provider == repository.BITBUCKET {
		bitbucket.Listen(repo)
	}
}
