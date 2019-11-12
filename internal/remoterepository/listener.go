package remoterepository

import (
	"miya/internal/remoterepository/bitbucket"
	"miya/internal/repository"
)

func ListenRemoteRepositories() {
	repositories := repository.GetAllRepositories()

	if len(repositories) != 0 {
		listenRepo(&repositories[0])
	}
}

func listenRepo(repo *repository.Repository) {

	if repo.Provider == repository.BITBUCKET {
		bitbucket.Listen(repo)
	}

}
