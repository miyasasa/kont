package util

const (
	BitbucketBasePath       = "/rest/api/1.0"
	BitbucketProjectPath    = "/projects/"
	BitbucketRepositoryPath = "/repos/"
	BitbucketPrPath         = "/pull-requests"
	BitbucketUserPath       = "/permissions/users"
)

// get project name & repo name dynamically from repo
func BitbucketFetchRepoUsersURL(host string, projectName string, repoName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketRepositoryPath +
		repoName +
		BitbucketUserPath

}

// get project name dynamically from repo
func BitbucketFetchProjectUsersURL(host string, projectName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketUserPath
}

func BitbucketFetchPrListURL(host string, projectName string, repoName string) string {
	return host +
		BitbucketBasePath +
		BitbucketProjectPath +
		projectName +
		BitbucketRepositoryPath +
		repoName +
		BitbucketPrPath
}
