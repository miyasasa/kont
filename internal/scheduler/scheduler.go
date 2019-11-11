package scheduler

import (
	"miya/internal/remoterepository/bitbucket"
	"miya/internal/repository"
)

func ListenRepositories() {
	repositories := repository.GetAllRepositories()

	if len(repositories) != 0 {
		schedule(&repositories[0])
	}
}

func schedule(repo *repository.Repository) {

	// scheduler.Every(5).Seconds().Run(bitbucket.Listen)

	if repo.Provider == repository.BITBUCKET {
		bitbucket.Listen(repo)
	}

}
