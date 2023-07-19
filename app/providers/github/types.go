package github

import "time"

const (
	ACTION_PING = "ping"
)

type GithubEvent struct {
	Sender       GithubUser            `mapstructure:"sender"`
	Repository   GithubRepository      `mapstructure:"repository"`
	Organization GithubOrganization    `mapstructure:"organization"`
	Installation GithubAppInstallation `mapstructure:"installation"`
	Other        map[string]any        `mapstructure:",remain"`
}

type GithubUser struct {
	Name              string `mapstructure:"name"`
	Email             string `mapstructure:"email"`
	Login             string `mapstructure:"login"`
	Id                uint32 `mapstructure:"id"`
	NodeId            string `mapstructure:"node_id"`
	AvatarUrl         string `mapstructure:"avatar_url"`
	GravatarId        string `mapstructure:"gravatar_id"`
	Url               string `mapstructure:"url"`
	HtmlUrl           string `mapstructure:"html_url"`
	FollowersUrl      string `mapstructure:"followers_url"`
	FollowingUrl      string `mapstructure:"following_url"`
	GistsUrl          string `mapstructure:"gists_url"`
	StarredUrl        string `mapstructure:"starred_url"`
	SubscriptionsUrl  string `mapstructure:"subscriptions_url"`
	OrganizationsUrl  string `mapstructure:"organizations_url"`
	ReposUrl          string `mapstructure:"repos_url"`
	EventsUrl         string `mapstructure:"events_url"`
	ReceivedEventsUrl string `mapstructure:"received_events_url"`
	Type              string `mapstructure:"type"`
	SiteAdmin         bool   `mapstructure:"site_admin"`
}

type GithubRepository struct {
	Id                       uint32     `mapstructure:"id"`
	NodeId                   string     `mapstructure:"node_id"`
	Name                     string     `mapstructure:"name"`
	FullName                 string     `mapstructure:"full_name"`
	Private                  bool       `mapstructure:"private"`
	Owner                    GithubUser `mapstructure:"owner"`
	HtmlUrl                  string     `mapstructure:"html_url"`
	Description              string     `mapstructure:"description"`
	Fork                     bool       `mapstructure:"fork"`
	Url                      string     `mapstructure:"url"`
	ForksUrl                 string     `mapstructure:"forks_url"`
	KeysUrl                  string     `mapstructure:"keys_url"`
	CollaboratorsUrl         string     `mapstructure:"collaborators_url"`
	TeamsUrl                 string     `mapstructure:"teams_url"`
	HooksUrl                 string     `mapstructure:"hooks_url"`
	IssueEventsUrl           string     `mapstructure:"issue_events_url"`
	EventsUrl                string     `mapstructure:"events_url"`
	AssigneesUrl             string     `mapstructure:"assignees_url"`
	BranchesUrl              string     `mapstructure:"branches_url"`
	TagsUrl                  string     `mapstructure:"tags_url"`
	BlobsUrl                 string     `mapstructure:"blobs_url"`
	GitTagsUrl               string     `mapstructure:"git_tags_url"`
	GitRefsUrl               string     `mapstructure:"git_refs_url"`
	TreesUrl                 string     `mapstructure:"trees_url"`
	StatusesUrl              string     `mapstructure:"statuses_url"`
	LanguagesUrl             string     `mapstructure:"languages_url"`
	StargazersUrl            string     `mapstructure:"stargazers_url"`
	ContributorsUrl          string     `mapstructure:"contributors_url"`
	SubscribersUrl           string     `mapstructure:"subscribers_url"`
	SubscriptionUrl          string     `mapstructure:"subscription_url"`
	CommitsUrl               string     `mapstructure:"commits_url"`
	GitCommitsUrl            string     `mapstructure:"git_commits_url"`
	CommentsUrl              string     `mapstructure:"comments_url"`
	IssueCommentUrl          string     `mapstructure:"issue_comment_url"`
	ContentsUrl              string     `mapstructure:"contents_url"`
	CompareUrl               string     `mapstructure:"compare_url"`
	MergesUrl                string     `mapstructure:"merges_url"`
	ArchiveUrl               string     `mapstructure:"archive_url"`
	DownloadsUrl             string     `mapstructure:"downloads_url"`
	IssuesUrl                string     `mapstructure:"issues_url"`
	PullsUrl                 string     `mapstructure:"pulls_url"`
	MilestonesUrl            string     `mapstructure:"milestones_url"`
	NotificationsUrl         string     `mapstructure:"notifications_url"`
	LabelsUrl                string     `mapstructure:"labels_url"`
	ReleasesUrl              string     `mapstructure:"releases_url"`
	DeploymentsUrl           string     `mapstructure:"deployments_url"`
	CreatedAt                time.Time  `mapstructure:"created_at"`
	UpdatedAt                time.Time  `mapstructure:"updated_at"`
	PushedAt                 time.Time  `mapstructure:"pushed_at"`
	GitUrl                   string     `mapstructure:"git_url"`
	SshUrl                   string     `mapstructure:"ssh_url"`
	CloneUrl                 string     `mapstructure:"clone_url"`
	SvnUrl                   string     `mapstructure:"svn_url"`
	Homepage                 string     `mapstructure:"homepage"`
	Size                     uint32     `mapstructure:"size"`
	StargazersCount          uint32     `mapstructure:"stargazers_count"`
	WatchersCount            uint32     `mapstructure:"watchers_count"`
	Language                 string     `mapstructure:"language"`
	HasIssues                bool       `mapstructure:"has_issues"`
	HasProjects              bool       `mapstructure:"has_projects"`
	HasDownloads             bool       `mapstructure:"has_downloads"`
	HasWiki                  bool       `mapstructure:"has_wiki"`
	HasPages                 bool       `mapstructure:"has_pages"`
	HasDiscussions           bool       `mapstructure:"has_discussions"`
	ForksCount               uint32     `mapstructure:"forks_count"`
	MirrorUrl                string     `mapstructure:"mirror_url"`
	Archived                 bool       `mapstructure:"archived"`
	Disabled                 bool       `mapstructure:"disabled"`
	OpenIssuesCount          uint32     `mapstructure:"open_issues_count"`
	License                  string     `mapstructure:"license"`
	AllowForking             bool       `mapstructure:"allow_forking"`
	IsTemplate               bool       `mapstructure:"is_template"`
	WebCommitSignoffRequired bool       `mapstructure:"web_commit_signoff_required"`
	Topics                   []string   `mapstructure:"topics"`
	Visibility               string     `mapstructure:"visibility"`
	Forks                    uint32     `mapstructure:"forks"`
	OpenIssues               uint32     `mapstructure:"open_issues"`
	Watchers                 uint32     `mapstructure:"watchers"`
	DefaultBranch            string     `mapstructure:"default_branch"`
}

type GithubOrganization struct {
}

type GithubAppInstallation struct {
}

type GitUser struct {
	Name     string `mapstructure:"name"`
	Email    string `mapstructure:"email"`
	Username string `mapstructure:"username"`
}

type GithubPushEvent struct {
	GithubEvent
	Ref        string         `mapstructure:"ref"`
	Before     string         `mapstructure:"before"`
	After      string         `mapstructure:"after"`
	Pusher     GitUser        `mapstructure:"pusher"`
	Created    bool           `mapstructure:"created"`
	Deleted    bool           `mapstructure:"deleted"`
	Forced     bool           `mapstructure:"forced"`
	BaseRef    string         `mapstructure:"base_ref"`
	Compare    string         `mapstructure:"compare"`
	Commits    []GithubCommit `mapstructure:"commits"`
	HeadCommit GithubCommit   `mapstructure:"head_commit"`
}

type GithubCommit struct {
	Id        string    `mapstructure:"id"`
	TreeId    string    `mapstructure:"tree_id"`
	Distinct  bool      `mapstructure:"distinct"`
	Message   string    `mapstructure:"message"`
	Timestamp time.Time `mapstructure:"timestamp"`
	Url       string    `mapstructure:"url"`
	Author    GitUser   `mapstructure:"author"`
	Committer GitUser   `mapstructure:"committer"`
	Added     []string  `mapstructure:"added"`
	Removed   []string  `mapstructure:"removed"`
	Modified  []string  `mapstructure:"modified"`
}
