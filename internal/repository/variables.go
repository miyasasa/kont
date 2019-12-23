package repository

const (
	BitbucketBasePath       = "/rest/api/1.0"
	BitbucketProjectPath    = "/projects/"
	BitbucketRepositoryPath = "/repos/"
	BitbucketPrPath         = "/pull-requests"
	BitbucketUserPath       = "/permissions/users"
)

// get project name & repo name dynamically from repo
func bitbucketFetchRepoUsersURL(host string, projectName string, repoName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketRepositoryPath +
		repoName +
		BitbucketUserPath

}

// get project name dynamically from repo
func bitbucketFetchProjectUsersURL(host string, projectName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketUserPath
}

func bitbucketFetchPrListURL(host string, projectName string, repoName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketRepositoryPath +
		repoName +
		BitbucketPrPath
}
