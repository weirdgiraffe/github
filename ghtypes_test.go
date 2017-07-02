//
// ghtypes_test.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestParseStarredResponse(t *testing.T) {
	r := bytes.NewBufferString(exampleResponseStarredRepos)
	var repo []StarredRepo
	err := json.NewDecoder(r).Decode(&repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(repo) == 0 {
		t.Errorf("no repositories decoded")
	}
}

func TestParseReadme(t *testing.T) {
	r := bytes.NewBufferString(exampleResponseReadme)
	var readme Readme
	err := json.NewDecoder(r).Decode(&readme)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseUser(t *testing.T) {
	r := bytes.NewBufferString(exampleResponseReadme)
	var user User
	err := json.NewDecoder(r).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}
}

const exampleResponseStarredRepos = `[
    {
        "starred_at": "2017-06-26T19:03:48Z",
        "repo": {
            "id": 19989605,
            "name": "certificate-transparency",
            "full_name": "google/certificate-transparency",
            "owner": {
                "login": "google",
                "id": 1342004,
                "avatar_url": "https://avatars2.githubusercontent.com/u/1342004?v=3",
                "gravatar_id": "",
                "url": "https://api.github.com/users/google",
                "html_url": "https://github.com/google",
                "followers_url": "https://api.github.com/users/google/followers",
                "following_url": "https://api.github.com/users/google/following{/other_user}",
                "gists_url": "https://api.github.com/users/google/gists{/gist_id}",
                "starred_url": "https://api.github.com/users/google/starred{/owner}{/repo}",
                "subscriptions_url": "https://api.github.com/users/google/subscriptions",
                "organizations_url": "https://api.github.com/users/google/orgs",
                "repos_url": "https://api.github.com/users/google/repos",
                "events_url": "https://api.github.com/users/google/events{/privacy}",
                "received_events_url": "https://api.github.com/users/google/received_events",
                "type": "Organization",
                "site_admin": false
            },
            "private": false,
            "html_url": "https://github.com/google/certificate-transparency",
            "description": "Auditing for TLS certificates.",
            "fork": false,
            "url": "https://api.github.com/repos/google/certificate-transparency",
            "forks_url": "https://api.github.com/repos/google/certificate-transparency/forks",
            "keys_url": "https://api.github.com/repos/google/certificate-transparency/keys{/key_id}",
            "collaborators_url": "https://api.github.com/repos/google/certificate-transparency/collaborators{/collaborator}",
            "teams_url": "https://api.github.com/repos/google/certificate-transparency/teams",
            "hooks_url": "https://api.github.com/repos/google/certificate-transparency/hooks",
            "issue_events_url": "https://api.github.com/repos/google/certificate-transparency/issues/events{/number}",
            "events_url": "https://api.github.com/repos/google/certificate-transparency/events",
            "assignees_url": "https://api.github.com/repos/google/certificate-transparency/assignees{/user}",
            "branches_url": "https://api.github.com/repos/google/certificate-transparency/branches{/branch}",
            "tags_url": "https://api.github.com/repos/google/certificate-transparency/tags",
            "blobs_url": "https://api.github.com/repos/google/certificate-transparency/git/blobs{/sha}",
            "git_tags_url": "https://api.github.com/repos/google/certificate-transparency/git/tags{/sha}",
            "git_refs_url": "https://api.github.com/repos/google/certificate-transparency/git/refs{/sha}",
            "trees_url": "https://api.github.com/repos/google/certificate-transparency/git/trees{/sha}",
            "statuses_url": "https://api.github.com/repos/google/certificate-transparency/statuses/{sha}",
            "languages_url": "https://api.github.com/repos/google/certificate-transparency/languages",
            "stargazers_url": "https://api.github.com/repos/google/certificate-transparency/stargazers",
            "contributors_url": "https://api.github.com/repos/google/certificate-transparency/contributors",
            "subscribers_url": "https://api.github.com/repos/google/certificate-transparency/subscribers",
            "subscription_url": "https://api.github.com/repos/google/certificate-transparency/subscription",
            "commits_url": "https://api.github.com/repos/google/certificate-transparency/commits{/sha}",
            "git_commits_url": "https://api.github.com/repos/google/certificate-transparency/git/commits{/sha}",
            "comments_url": "https://api.github.com/repos/google/certificate-transparency/comments{/number}",
            "issue_comment_url": "https://api.github.com/repos/google/certificate-transparency/issues/comments{/number}",
            "contents_url": "https://api.github.com/repos/google/certificate-transparency/contents/{+path}",
            "compare_url": "https://api.github.com/repos/google/certificate-transparency/compare/{base}...{head}",
            "merges_url": "https://api.github.com/repos/google/certificate-transparency/merges",
            "archive_url": "https://api.github.com/repos/google/certificate-transparency/{archive_format}{/ref}",
            "downloads_url": "https://api.github.com/repos/google/certificate-transparency/downloads",
            "issues_url": "https://api.github.com/repos/google/certificate-transparency/issues{/number}",
            "pulls_url": "https://api.github.com/repos/google/certificate-transparency/pulls{/number}",
            "milestones_url": "https://api.github.com/repos/google/certificate-transparency/milestones{/number}",
            "notifications_url": "https://api.github.com/repos/google/certificate-transparency/notifications{?since,all,participating}",
            "labels_url": "https://api.github.com/repos/google/certificate-transparency/labels{/name}",
            "releases_url": "https://api.github.com/repos/google/certificate-transparency/releases{/id}",
            "deployments_url": "https://api.github.com/repos/google/certificate-transparency/deployments",
            "created_at": "2014-05-20T17:03:47Z",
            "updated_at": "2017-06-28T02:35:56Z",
            "pushed_at": "2017-06-27T15:41:41Z",
            "git_url": "git://github.com/google/certificate-transparency.git",
            "ssh_url": "git@github.com:google/certificate-transparency.git",
            "clone_url": "https://github.com/google/certificate-transparency.git",
            "svn_url": "https://github.com/google/certificate-transparency",
            "homepage": "http://www.certificate-transparency.org/",
            "size": 27731,
            "stargazers_count": 428,
            "watchers_count": 428,
            "language": "C++",
            "has_issues": true,
            "has_projects": true,
            "has_downloads": true,
            "has_wiki": false,
            "has_pages": false,
            "forks_count": 155,
            "mirror_url": null,
            "open_issues_count": 94,
            "forks": 155,
            "open_issues": 94,
            "watchers": 428,
            "default_branch": "master"
        }
    }
]`

const exampleResponseReadme = `{
    "name": "README.md",
    "path": "README.md",
    "sha": "0cdc1fa1ae5318b32a7b417eda6b4ea19db4272a",
    "size": 1728,
    "url": "https://api.github.com/repos/sgur/vim-editorconfig/contents/README.md?ref=master",
    "html_url": "https://github.com/sgur/vim-editorconfig/blob/master/README.md",
    "git_url": "https://api.github.com/repos/sgur/vim-editorconfig/git/blobs/0cdc1fa1ae5318b32a7b417eda6b4ea19db4272a",
    "download_url": "https://raw.githubusercontent.com/sgur/vim-editorconfig/master/README.md",
    "type": "file",
    "content": "dmltLWVkaXRvcmNvbmZpZw0KPT09PT09PT09PT09PT09PQ0KDQpZZXQgYW5v\ndGhlciBWaW0gcGx1Z2luIGZvciBbRWRpdG9yQ29uZmlnXShodHRwOi8vZWRp\ndG9yY29uZmlnLm9yZykNCg0KRGVzY3JpcHRpb24NCi0tLS0tLS0tLS0tDQoN\nCiMjIyBTdXBwb3J0ZWQgUHJvcGVydGllcw0KDQotIGBjaGFyc2V0YA0KLSBg\nZW5kX29mX2xpbmVgDQotIGBpbmRlbnRfc2l6ZWANCi0gYGluZGVudF9zdHls\nZWANCi0gYGluc2VydF9maW5hbF9uZXdsaW5lYA0KLSBgbWF4X2xpbmVfbGVu\nZ3RoYA0KLSBgcm9vdGANCi0gYHRhYl93aWR0aGANCi0gYHRyaW1fdHJhaWxp\nbmdfd2hpdGVzcGFjZWANCg0KUHJvcGVydGllcyBiZWxvdyBhcmUgZW5hYmxl\nZCBvbmx5IGluIHRoaXMgcGx1Z2luOg0KDQotIGBsb2NhbF92aW1yY2ANCg0K\nVi5TLg0KLS0tLQ0KDQotIFtlZGl0b3Jjb25maWcvZWRpdG9yY29uZmlnLXZp\nbV0oaHR0cHM6Ly9naXRodWIuY29tL2VkaXRvcmNvbmZpZy9lZGl0b3Jjb25m\naWctdmltKQ0KDQpbZWRpdG9yY29uZmlnLXZpbV0oaHR0cHM6Ly9naXRodWIu\nY29tL2VkaXRvcmNvbmZpZy9lZGl0b3Jjb25maWctdmltKSBpcyBvZmZpY2lh\nbCB2aW0gcGx1Z2luIGZvciBFZGl0b3JDb25maWcuDQpUaGlzIHJlcXVpcmVz\nIGBpZl9weXRob25gIGludGVyZmFjZSBvciBleHRlcm5hbCBweXRob24gaW50\nZXJwcmV0ZXIuDQoNClt2aW0tZWRpdG9yY29uZmlnXShodHRwczovL2dpdGh1\nYi5jb20vc2d1ci92aW0tZWRpdG9yY29uZmlnKSBpcyB3cml0dGVuIGluIHB1\ncmUgdmltc2NyaXB0Lg0KWW91IGNhbiB1c2UgZWRpdG9yY29uZmlnIHdpdGhv\ndXQgYW55IGV4dGVybmFsIGludGVyZmFjZXMgc3VjaCBhcyBgaWZfcHl0aG9u\nYC4NCg0KVXNhZ2UNCi0tLS0tDQoNCiAxLiBJbnN0YWxsIHRoZSBwbHVnaW4N\nCiAyLiBMb2NhdGUgYC5lZGl0b3Jjb25maWdgDQogMy4gRWRpdCBzb21lIGZp\nbGVzDQoNCkZlYXR1cmVzDQotLS0tLS0tDQoNCiMjIyBsb2NhbFxfdmltcmMN\nCg0KU291cmNlIHNwZWNpZmllZCB2aW1yYyBmaWxlIHdoZW4geW91IGVkaXQg\nYSBidWZmZXIuDQoNCkl0IGJlaGF2ZXMgbGlrZSBbdGhpbmNhL3ZpbS1sb2Nh\nbHJjXShodHRwczovL2dpdGh1Yi5jb20vdGhpbmNhL3ZpbS1sb2NhbHJjKS4N\nCg0KYGBgDQpbKi5tZF0NCmxvY2FsX3ZpbXJjID0gLmxvY2FsLnZpbXJjDQpg\nYGANCg0KT3B0aW9ucw0KLS0tLS0tLQ0KDQojIyMgZzplZGl0b3Jjb25maWdc\nX2JsYWNrbGlzdA0KDQpFeGNsdWRlIHJlZ2V4cCBwYXR0ZXJucyBmb3IgZmls\nZXR5cGVzIG9yIGZpbGVwYXRocw0KDQpgYGB2aW0NCmxldCBnOmVkaXRvcmNv\nbmZpZ19ibGFja2xpc3QgPSB7DQogICAgXCAnZmlsZXR5cGUnOiBbJ2dpdC4q\nJywgJ2Z1Z2l0aXZlJ10sDQogICAgXCAncGF0dGVybic6IFsnXC51bn4kJ119\nDQpgYGANCg0KIyMjIGc6ZWRpdG9yY29uZmlnXF9yb290XF9jaGRpcg0KDQpB\ndXRvbWF0aWNhbGx5IGA6bGNkYCBJZiBgcm9vdCA9IHRydWVgIGV4aXN0cyBp\nbiBgLmVkaXRvcmNvbmZpZ2AuDQoNCiMjIyBnOmVkaXRvcmNvbmZpZ1xfdmVy\nYm9zZQ0KDQpTaG93IHZlcmJvc2UgbWVzc2FnZXMNCg0KYGBgdmltDQpsZXQg\nZzplZGl0b3Jjb25maWdfdmVyYm9zZSA9IDENCmBgYA0KDQpJbnN0YWxsDQot\nLS0tLS0tDQoNClVzZSB5b3VyIGZhdm9yaXRlIHBsdWdpbiBtYW5hZ2VyLg0K\nDQpMaWNlbnNlDQotLS0tLS0tDQoNCk1JVCBMaWNlbnNlDQoNCkF1dGhvcg0K\nLS0tLS0tDQoNCnNndXINCg0K\n",
    "encoding": "base64",
    "_links": {
        "self": "https://api.github.com/repos/sgur/vim-editorconfig/contents/README.md?ref=master",
        "git": "https://api.github.com/repos/sgur/vim-editorconfig/git/blobs/0cdc1fa1ae5318b32a7b417eda6b4ea19db4272a",
        "html": "https://github.com/sgur/vim-editorconfig/blob/master/README.md"
    }
}`

const exampleResponseUser = `{
    "login": "weirdgiraffe",
    "id": 8282530,
    "avatar_url": "https://avatars0.githubusercontent.com/u/8282530?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/weirdgiraffe",
    "html_url": "https://github.com/weirdgiraffe",
    "followers_url": "https://api.github.com/users/weirdgiraffe/followers",
    "following_url": "https://api.github.com/users/weirdgiraffe/following{/other_user}",
    "gists_url": "https://api.github.com/users/weirdgiraffe/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/weirdgiraffe/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/weirdgiraffe/subscriptions",
    "organizations_url": "https://api.github.com/users/weirdgiraffe/orgs",
    "repos_url": "https://api.github.com/users/weirdgiraffe/repos",
    "events_url": "https://api.github.com/users/weirdgiraffe/events{/privacy}",
    "received_events_url": "https://api.github.com/users/weirdgiraffe/received_events",
    "type": "User",
    "site_admin": false,
    "name": null,
    "company": null,
    "blog": "",
    "location": "Berlin",
    "email": "giraffe@cyberzoo.xyz",
    "hireable": null,
    "bio": null,
    "public_repos": 20,
    "public_gists": 9,
    "followers": 1,
    "following": 12,
    "created_at": "2014-07-27T16:42:48Z",
    "updated_at": "2017-04-22T13:26:18Z",
    "private_gists": 29,
    "total_private_repos": 0,
    "owned_private_repos": 0,
    "disk_usage": 17237,
    "collaborators": 0,
    "two_factor_authentication": false,
    "plan": {
        "name": "free",
        "space": 976562499,
        "collaborators": 0,
        "private_repos": 0
    }
}`
