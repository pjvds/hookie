package main

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rawRequest(content string) (*http.Request, error) {
	return http.ReadRequest(bufio.NewReader(bytes.NewBufferString(content)))
}

func mustRawRequest(content string) *http.Request {
	request, err := rawRequest(content)
	if err != nil {
		panic(err.Error())
	}
	return request
}

func TestGithubSignatureValidatorAcceptsValidSignature(t *testing.T) {
	assert := assert.New(t)

	const raw = "POST / HTTP/1.1\n" +
		"Host: hook.publichost.io\n" +
		"Accept: */*\n" +
		"User-Agent: GitHub-Hookshot/40ea4ec\n" +
		"X-GitHub-Event: push\n" +
		"X-GitHub-Delivery: d2ef0380-a7d9-11e6-8b1e-c6ac4194fcb8\n" +
		"content-type: application/json\n" +
		"X-Hub-Signature: sha1=21178161e2fc27f5598342b6759e6a395a1fd008\n" +
		"Content-Length: 5972\n" +
		"\n" +
		"{\"ref\":\"refs/heads/master\",\"before\":\"05ada2f017175bf69b48c089e5b2cf53df8a7353\",\"after\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"created\":false,\"deleted\":false,\"forced\":false,\"base_ref\":null,\"compare\":\"https://github.com/pjvds/publichost/compare/05ada2f01717...3d0be96c8de9\",\"commits\":[{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]}],\"head_commit\":{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]},\"repository\":{\"id\":15182873,\"name\":\"publichost\",\"full_name\":\"pjvds/publichost\",\"owner\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"private\":false,\"html_url\":\"https://github.com/pjvds/publichost\",\"description\":null,\"fork\":false,\"url\":\"https://github.com/pjvds/publichost\",\"forks_url\":\"https://api.github.com/repos/pjvds/publichost/forks\",\"keys_url\":\"https://api.github.com/repos/pjvds/publichost/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/pjvds/publichost/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/pjvds/publichost/teams\",\"hooks_url\":\"https://api.github.com/repos/pjvds/publichost/hooks\",\"issue_events_url\":\"https://api.github.com/repos/pjvds/publichost/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/pjvds/publichost/events\",\"assignees_url\":\"https://api.github.com/repos/pjvds/publichost/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/pjvds/publichost/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/pjvds/publichost/tags\",\"blobs_url\":\"https://api.github.com/repos/pjvds/publichost/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/pjvds/publichost/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/pjvds/publichost/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/pjvds/publichost/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/pjvds/publichost/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/pjvds/publichost/languages\",\"stargazers_url\":\"https://api.github.com/repos/pjvds/publichost/stargazers\",\"contributors_url\":\"https://api.github.com/repos/pjvds/publichost/contributors\",\"subscribers_url\":\"https://api.github.com/repos/pjvds/publichost/subscribers\",\"subscription_url\":\"https://api.github.com/repos/pjvds/publichost/subscription\",\"commits_url\":\"https://api.github.com/repos/pjvds/publichost/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/pjvds/publichost/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/pjvds/publichost/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/pjvds/publichost/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/pjvds/publichost/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/pjvds/publichost/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/pjvds/publichost/merges\",\"archive_url\":\"https://api.github.com/repos/pjvds/publichost/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/pjvds/publichost/downloads\",\"issues_url\":\"https://api.github.com/repos/pjvds/publichost/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/pjvds/publichost/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/pjvds/publichost/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/pjvds/publichost/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/pjvds/publichost/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/pjvds/publichost/releases{/id}\",\"deployments_url\":\"https://api.github.com/repos/pjvds/publichost/deployments\",\"created_at\":1387011790,\"updated_at\":\"2016-04-16T08:01:51Z\",\"pushed_at\":1478846467,\"git_url\":\"git://github.com/pjvds/publichost.git\",\"ssh_url\":\"git@github.com:pjvds/publichost.git\",\"clone_url\":\"https://github.com/pjvds/publichost.git\",\"svn_url\":\"https://github.com/pjvds/publichost\",\"homepage\":null,\"size\":2682,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"open_issues_count\":0,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"master\",\"stargazers\":0,\"master_branch\":\"master\"},\"pusher\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"sender\":{\"login\":\"pjvds\",\"id\":150387,\"avatar_url\":\"https://avatars.githubusercontent.com/u/150387?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/pjvds\",\"html_url\":\"https://github.com/pjvds\",\"followers_url\":\"https://api.github.com/users/pjvds/followers\",\"following_url\":\"https://api.github.com/users/pjvds/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/pjvds/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/pjvds/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/pjvds/subscriptions\",\"organizations_url\":\"https://api.github.com/users/pjvds/orgs\",\"repos_url\":\"https://api.github.com/users/pjvds/repos\",\"events_url\":\"https://api.github.com/users/pjvds/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/pjvds/received_events\",\"type\":\"User\",\"site_admin\":false}}"

	invoked := false
	var upstreamRequest *http.Request
	var upstreamResponse http.ResponseWriter
	upstreamHandler := http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		invoked = true
		upstreamRequest = request
		upstreamResponse = response
	})

	request := mustRawRequest(raw)
	response := httptest.NewRecorder()

	sut := GithubSignatureValidator{
		Handler: upstreamHandler,
		Secret:  []byte("123456"),
	}
	sut.ServeHTTP(response, request)

	assert.True(invoked, "did not invoke upstream hander")
	assert.NotNil(upstreamRequest, "upstream handler did not receive request")
	assert.NotNil(upstreamRequest, "upstream handler did not receive response")
}

func TestGithubSignatureValidatorRejectsMissingSignature(t *testing.T) {
	assert := assert.New(t)

	const raw = "POST / HTTP/1.1\n" +
		"Host: hook.publichost.io\n" +
		"Accept: */*\n" +
		"User-Agent: GitHub-Hookshot/40ea4ec\n" +
		"X-GitHub-Event: push\n" +
		"X-GitHub-Delivery: d2ef0380-a7d9-11e6-8b1e-c6ac4194fcb8\n" +
		"content-type: application/json\n" +
		"Content-Length: 5972\n" +
		"\n" +
		"{\"ref\":\"refs/heads/master\",\"before\":\"05ada2f017175bf69b48c089e5b2cf53df8a7353\",\"after\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"created\":false,\"deleted\":false,\"forced\":false,\"base_ref\":null,\"compare\":\"https://github.com/pjvds/publichost/compare/05ada2f01717...3d0be96c8de9\",\"commits\":[{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]}],\"head_commit\":{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]},\"repository\":{\"id\":15182873,\"name\":\"publichost\",\"full_name\":\"pjvds/publichost\",\"owner\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"private\":false,\"html_url\":\"https://github.com/pjvds/publichost\",\"description\":null,\"fork\":false,\"url\":\"https://github.com/pjvds/publichost\",\"forks_url\":\"https://api.github.com/repos/pjvds/publichost/forks\",\"keys_url\":\"https://api.github.com/repos/pjvds/publichost/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/pjvds/publichost/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/pjvds/publichost/teams\",\"hooks_url\":\"https://api.github.com/repos/pjvds/publichost/hooks\",\"issue_events_url\":\"https://api.github.com/repos/pjvds/publichost/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/pjvds/publichost/events\",\"assignees_url\":\"https://api.github.com/repos/pjvds/publichost/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/pjvds/publichost/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/pjvds/publichost/tags\",\"blobs_url\":\"https://api.github.com/repos/pjvds/publichost/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/pjvds/publichost/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/pjvds/publichost/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/pjvds/publichost/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/pjvds/publichost/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/pjvds/publichost/languages\",\"stargazers_url\":\"https://api.github.com/repos/pjvds/publichost/stargazers\",\"contributors_url\":\"https://api.github.com/repos/pjvds/publichost/contributors\",\"subscribers_url\":\"https://api.github.com/repos/pjvds/publichost/subscribers\",\"subscription_url\":\"https://api.github.com/repos/pjvds/publichost/subscription\",\"commits_url\":\"https://api.github.com/repos/pjvds/publichost/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/pjvds/publichost/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/pjvds/publichost/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/pjvds/publichost/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/pjvds/publichost/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/pjvds/publichost/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/pjvds/publichost/merges\",\"archive_url\":\"https://api.github.com/repos/pjvds/publichost/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/pjvds/publichost/downloads\",\"issues_url\":\"https://api.github.com/repos/pjvds/publichost/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/pjvds/publichost/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/pjvds/publichost/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/pjvds/publichost/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/pjvds/publichost/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/pjvds/publichost/releases{/id}\",\"deployments_url\":\"https://api.github.com/repos/pjvds/publichost/deployments\",\"created_at\":1387011790,\"updated_at\":\"2016-04-16T08:01:51Z\",\"pushed_at\":1478846467,\"git_url\":\"git://github.com/pjvds/publichost.git\",\"ssh_url\":\"git@github.com:pjvds/publichost.git\",\"clone_url\":\"https://github.com/pjvds/publichost.git\",\"svn_url\":\"https://github.com/pjvds/publichost\",\"homepage\":null,\"size\":2682,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"open_issues_count\":0,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"master\",\"stargazers\":0,\"master_branch\":\"master\"},\"pusher\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"sender\":{\"login\":\"pjvds\",\"id\":150387,\"avatar_url\":\"https://avatars.githubusercontent.com/u/150387?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/pjvds\",\"html_url\":\"https://github.com/pjvds\",\"followers_url\":\"https://api.github.com/users/pjvds/followers\",\"following_url\":\"https://api.github.com/users/pjvds/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/pjvds/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/pjvds/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/pjvds/subscriptions\",\"organizations_url\":\"https://api.github.com/users/pjvds/orgs\",\"repos_url\":\"https://api.github.com/users/pjvds/repos\",\"events_url\":\"https://api.github.com/users/pjvds/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/pjvds/received_events\",\"type\":\"User\",\"site_admin\":false}}"

	invoked := false
	upstreamHandler := http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		invoked = true
	})

	request := mustRawRequest(raw)
	response := httptest.NewRecorder()

	sut := GithubSignatureValidator{
		Handler: upstreamHandler,
		Secret:  []byte("123456"),
	}
	sut.ServeHTTP(response, request)

	assert.False(invoked, "handler did invoke upstream hander")
	assert.Equal(http.StatusForbidden, response.Code)
}

func TestGithubSignatureValidatorRejectsInvalidSignature(t *testing.T) {
	assert := assert.New(t)

	const raw = "POST / HTTP/1.1\n" +
		"Host: hook.publichost.io\n" +
		"Accept: */*\n" +
		"User-Agent: GitHub-Hookshot/40ea4ec\n" +
		"X-GitHub-Event: push\n" +
		"X-GitHub-Delivery: d2ef0380-a7d9-11e6-8b1e-c6ac4194fcb8\n" +
		"X-Hub-Signature: sha1=4e231252d2fb36f4487231a7848d5b486b0fe117\n" +
		"content-type: application/json\n" +
		"Content-Length: 5972\n" +
		"\n" +
		"{\"ref\":\"refs/heads/master\",\"before\":\"05ada2f017175bf69b48c089e5b2cf53df8a7353\",\"after\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"created\":false,\"deleted\":false,\"forced\":false,\"base_ref\":null,\"compare\":\"https://github.com/pjvds/publichost/compare/05ada2f01717...3d0be96c8de9\",\"commits\":[{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]}],\"head_commit\":{\"id\":\"3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"tree_id\":\"e9720fc80d5e58cfa0f9e3c34dbd62529fb0a5b0\",\"distinct\":true,\"message\":\"get tunnel address from header\",\"timestamp\":\"2016-11-11T07:40:50+01:00\",\"url\":\"https://github.com/pjvds/publichost/commit/3d0be96c8de9632738a87c81ee5ec9e4d638135b\",\"author\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"committer\":{\"name\":\"Pieter Joost van de Sande\",\"email\":\"pj@born2code.net\",\"username\":\"pjvds\"},\"added\":[],\"removed\":[],\"modified\":[\"client/main.go\"]},\"repository\":{\"id\":15182873,\"name\":\"publichost\",\"full_name\":\"pjvds/publichost\",\"owner\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"private\":false,\"html_url\":\"https://github.com/pjvds/publichost\",\"description\":null,\"fork\":false,\"url\":\"https://github.com/pjvds/publichost\",\"forks_url\":\"https://api.github.com/repos/pjvds/publichost/forks\",\"keys_url\":\"https://api.github.com/repos/pjvds/publichost/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/pjvds/publichost/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/pjvds/publichost/teams\",\"hooks_url\":\"https://api.github.com/repos/pjvds/publichost/hooks\",\"issue_events_url\":\"https://api.github.com/repos/pjvds/publichost/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/pjvds/publichost/events\",\"assignees_url\":\"https://api.github.com/repos/pjvds/publichost/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/pjvds/publichost/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/pjvds/publichost/tags\",\"blobs_url\":\"https://api.github.com/repos/pjvds/publichost/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/pjvds/publichost/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/pjvds/publichost/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/pjvds/publichost/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/pjvds/publichost/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/pjvds/publichost/languages\",\"stargazers_url\":\"https://api.github.com/repos/pjvds/publichost/stargazers\",\"contributors_url\":\"https://api.github.com/repos/pjvds/publichost/contributors\",\"subscribers_url\":\"https://api.github.com/repos/pjvds/publichost/subscribers\",\"subscription_url\":\"https://api.github.com/repos/pjvds/publichost/subscription\",\"commits_url\":\"https://api.github.com/repos/pjvds/publichost/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/pjvds/publichost/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/pjvds/publichost/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/pjvds/publichost/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/pjvds/publichost/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/pjvds/publichost/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/pjvds/publichost/merges\",\"archive_url\":\"https://api.github.com/repos/pjvds/publichost/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/pjvds/publichost/downloads\",\"issues_url\":\"https://api.github.com/repos/pjvds/publichost/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/pjvds/publichost/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/pjvds/publichost/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/pjvds/publichost/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/pjvds/publichost/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/pjvds/publichost/releases{/id}\",\"deployments_url\":\"https://api.github.com/repos/pjvds/publichost/deployments\",\"created_at\":1387011790,\"updated_at\":\"2016-04-16T08:01:51Z\",\"pushed_at\":1478846467,\"git_url\":\"git://github.com/pjvds/publichost.git\",\"ssh_url\":\"git@github.com:pjvds/publichost.git\",\"clone_url\":\"https://github.com/pjvds/publichost.git\",\"svn_url\":\"https://github.com/pjvds/publichost\",\"homepage\":null,\"size\":2682,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"open_issues_count\":0,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"master\",\"stargazers\":0,\"master_branch\":\"master\"},\"pusher\":{\"name\":\"pjvds\",\"email\":\"pj@born2code.net\"},\"sender\":{\"login\":\"pjvds\",\"id\":150387,\"avatar_url\":\"https://avatars.githubusercontent.com/u/150387?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/pjvds\",\"html_url\":\"https://github.com/pjvds\",\"followers_url\":\"https://api.github.com/users/pjvds/followers\",\"following_url\":\"https://api.github.com/users/pjvds/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/pjvds/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/pjvds/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/pjvds/subscriptions\",\"organizations_url\":\"https://api.github.com/users/pjvds/orgs\",\"repos_url\":\"https://api.github.com/users/pjvds/repos\",\"events_url\":\"https://api.github.com/users/pjvds/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/pjvds/received_events\",\"type\":\"User\",\"site_admin\":false}}"

	invoked := false
	upstreamHandler := http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		invoked = true
	})

	request := mustRawRequest(raw)
	response := httptest.NewRecorder()

	sut := GithubSignatureValidator{
		Handler: upstreamHandler,
		Secret:  []byte("123456"),
	}
	sut.ServeHTTP(response, request)

	assert.False(invoked, "handler did invoke upstream hander")
	assert.Equal(http.StatusForbidden, response.Code)
}
