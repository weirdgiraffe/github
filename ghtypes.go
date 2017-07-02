//
// ghtypes.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//
package github

import "time"

type StarredRepo struct {
	Repo      Repo      `json:"repo"`
	StarredAt time.Time `json:"starred_at"`
}

type Repo struct {
	ArchiveURL       string    `json:"archive_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BlobsURL         string    `json:"blobs_url"`
	BranchesURL      string    `json:"branches_url"`
	CloneURL         string    `json:"clone_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	CommentsURL      string    `json:"comments_url"`
	CommitsURL       string    `json:"commits_url"`
	CompareURL       string    `json:"compare_url"`
	ContentsURL      string    `json:"contents_url"`
	ContributorsURL  string    `json:"contributors_url"`
	CreatedAt        time.Time `json:"created_at"`
	DefaultBranch    string    `json:"default_branch"`
	DeploymentsURL   string    `json:"deployments_url"`
	Description      string    `json:"description"`
	DownloadsURL     string    `json:"downloads_url"`
	EventsURL        string    `json:"events_url"`
	Fork             bool      `json:"fork"`
	Forks            int64     `json:"forks"`
	ForksCount       int64     `json:"forks_count"`
	ForksURL         string    `json:"forks_url"`
	FullName         string    `json:"full_name"`
	GitCommitsURL    string    `json:"git_commits_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitURL           string    `json:"git_url"`
	HasDownloads     bool      `json:"has_downloads"`
	HasIssues        bool      `json:"has_issues"`
	HasPages         bool      `json:"has_pages"`
	HasProjects      bool      `json:"has_projects"`
	HasWiki          bool      `json:"has_wiki"`
	Homepage         string    `json:"homepage"`
	HooksURL         string    `json:"hooks_url"`
	HTMLURL          string    `json:"html_url"`
	ID               int64     `json:"id"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	IssuesURL        string    `json:"issues_url"`
	KeysURL          string    `json:"keys_url"`
	LabelsURL        string    `json:"labels_url"`
	Language         string    `json:"language"`
	LanguagesURL     string    `json:"languages_url"`
	MergesURL        string    `json:"merges_url"`
	MilestonesURL    string    `json:"milestones_url"`
	MirrorURL        string    `json:"mirror_url"`
	Name             string    `json:"name"`
	NotificationsURL string    `json:"notifications_url"`
	OpenIssues       int64     `json:"open_issues"`
	OpenIssuesCount  int64     `json:"open_issues_count"`
	Owner            User      `json:"owner"`
	Private          bool      `json:"private"`
	PullsURL         string    `json:"pulls_url"`
	PushedAt         time.Time `json:"pushed_at"`
	ReleasesURL      string    `json:"releases_url"`
	Size             int64     `json:"size"`
	SSHURL           string    `json:"ssh_url"`
	StargazersCount  int64     `json:"stargazers_count"`
	StargazersURL    string    `json:"stargazers_url"`
	StatusesURL      string    `json:"statuses_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	SvnURL           string    `json:"svn_url"`
	TagsURL          string    `json:"tags_url"`
	TeamsURL         string    `json:"teams_url"`
	TreesURL         string    `json:"trees_url"`
	UpdatedAt        time.Time `json:"updated_at"`
	URL              string    `json:"url"`
	Watchers         int64     `json:"watchers"`
	WatchersCount    int64     `json:"watchers_count"`
}

type Readme struct {
	Links       Links  `json:"_links"`
	Content     []byte `json:"content"`
	DownloadURL string `json:"download_url"`
	Encoding    string `json:"encoding"`
	GitURL      string `json:"git_url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int64  `json:"size"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type Links struct {
	Git  string `json:"git"`
	HTML string `json:"html"`
	Self string `json:"self"`
}

type User struct {
	AvatarURL               string      `json:"avatar_url"`
	Bio                     *string     `json:"bio"`
	Blog                    string      `json:"blog"`
	Collaborators           int64       `json:"collaborators"`
	Company                 interface{} `json:"company"`
	CreatedAt               string      `json:"created_at"`
	DiskUsage               int64       `json:"disk_usage"`
	Email                   string      `json:"email"`
	EventsURL               string      `json:"events_url"`
	Followers               int64       `json:"followers"`
	FollowersURL            string      `json:"followers_url"`
	Following               int64       `json:"following"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	GravatarID              string      `json:"gravatar_id"`
	Hireable                interface{} `json:"hireable"`
	HTMLURL                 string      `json:"html_url"`
	ID                      int64       `json:"id"`
	Location                string      `json:"location"`
	Login                   string      `json:"login"`
	Name                    interface{} `json:"name"`
	OrganizationsURL        string      `json:"organizations_url"`
	OwnedPrivateRepos       int64       `json:"owned_private_repos"`
	Plan                    *UserPlan   `json:"plan,omitempty"`
	PrivateGists            int64       `json:"private_gists"`
	PublicGists             int64       `json:"public_gists"`
	PublicRepos             int64       `json:"public_repos"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	ReposURL                string      `json:"repos_url"`
	SiteAdmin               bool        `json:"site_admin"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	TotalPrivateRepos       int64       `json:"total_private_repos"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Type                    string      `json:"type"`
	UpdatedAt               string      `json:"updated_at"`
	URL                     string      `json:"url"`
}

type UserPlan struct {
	Collaborators int64  `json:"collaborators"`
	Name          string `json:"name"`
	PrivateRepos  int64  `json:"private_repos"`
	Space         int64  `json:"space"`
}
