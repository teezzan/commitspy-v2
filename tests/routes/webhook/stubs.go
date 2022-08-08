package webhook_test

var GitlabBody = []byte(
	`{
		"object_kind": "push",
		"event_name": "push",
		"before": "9fc6c4b32e13da46ad7f79dc77e4691668e9bcd2",
		"after": "7f42fba02a6ee11a62ee02ca47895eb7d671ff9a",
		"ref": "refs/heads/master",
		"checkout_sha": "7f42fba02a6ee11a62ee02ca47895eb7d671ff9a",
		"message": null,
		"user_id": 4814445,
		"user_name": "Yusuf Taiwo Hassan",
		"user_username": "teezzan",
		"user_email": "",
		"user_avatar": "https://gitlab.com/uploads/-/system/user/avatar/4814445/avatar.png",
		"project_id": 38048848,
		"project": {
			"id": 38048848,
			"name": "dummy-project",
			"description": null,
			"web_url": "https://gitlab.com/teezzan/dummy-project",
			"avatar_url": null,
			"git_ssh_url": "git@gitlab.com:teezzan/dummy-project.git",
			"git_http_url": "https://gitlab.com/teezzan/dummy-project.git",
			"namespace": "Yusuf Taiwo Hassan",
			"visibility_level": 20,
			"path_with_namespace": "teezzan/dummy-project",
			"default_branch": "master",
			"ci_config_path": "",
			"homepage": "https://gitlab.com/teezzan/dummy-project",
			"url": "git@gitlab.com:teezzan/dummy-project.git",
			"ssh_url": "git@gitlab.com:teezzan/dummy-project.git",
			"http_url": "https://gitlab.com/teezzan/dummy-project.git"
		},
		"commits": [
			{
				"id": "036007866981646805c517794d498bafe94c65a9",
				"message": "Docs Again\n",
				"title": "Docs Again",
				"timestamp": "2022-07-23T18:05:27+01:00",
				"url": "https://gitlab.com/teezzan/dummy-project/-/commit/036007866981646805c517794d498bafe94c65a9",
				"author": {
					"name": "Yusuf Taiwo Hassan",
					"email": "teehazzan@gmail.com"
				},
				"added": [],
				"modified": [
					"README.md"
				],
				"removed": []
			},
			{
				"id": "9fc6c4b32e13da46ad7f79dc77e4691668e9bcd2",
				"message": "Docs\n",
				"title": "Docs",
				"timestamp": "2022-07-23T18:05:09+01:00",
				"url": "https://gitlab.com/teezzan/dummy-project/-/commit/9fc6c4b32e13da46ad7f79dc77e4691668e9bcd2",
				"author": {
					"name": "Yusuf Taiwo Hassan",
					"email": "teehazzan@gmail.com"
				},
				"added": [],
				"modified": [
					"README.md"
				],
				"removed": []
			},
			{
				"id": "9fc6c4b32e13da46ad7f79dc77e5391668e9cbd2",
				"message": "Nope\n",
				"title": "Nope",
				"timestamp": "2022-07-23T19:05:09+01:00",
				"url": "https://gitlab.com/teezzan/dummy-project/-/commit/9fc6c4b32e13da46ad7f79dc77e4691668e9bcd2",
				"author": {
					"name": "Yusuf Taiwo Hassan",
					"email": "teehazzan@gmail.com"
				},
				"added": [],
				"modified": [
					"README.md"
				],
				"removed": []
			},
			{
				"id": "7f42fba02a6ee11a62ee02ca47895eb7d671ff9a",
				"message": "Merge branch 'feat/adding_new_stuff' into 'master'\n\nFeat/adding new stuff\n\nSee merge request teezzan/dummy-project!1",
				"title": "Merge branch 'feat/adding_new_stuff' into 'master'",
				"timestamp": "2022-07-23T17:06:49+00:00",
				"url": "https://gitlab.com/teezzan/dummy-project/-/commit/7f42fba02a6ee11a62ee02ca47895eb7d671ff9a",
				"author": {
					"name": "Yusuf Taiwo Hassan",
					"email": "teehazzan@gmail.com"
				},
				"added": [],
				"modified": [
					"Dockerfile",
					"README.md"
				],
				"removed": []
			}
		],
		"total_commits_count": 4,
		"push_options": {},
		"repository": {
			"name": "dummy-project",
			"url": "git@gitlab.com:teezzan/dummy-project.git",
			"description": null,
			"homepage": "https://gitlab.com/teezzan/dummy-project",
			"git_http_url": "https://gitlab.com/teezzan/dummy-project.git",
			"git_ssh_url": "git@gitlab.com:teezzan/dummy-project.git",
			"visibility_level": 20
		}
	}`)

var GithubBody = []byte(
	`{
	"ref": "refs/heads/feat/dummy_branch",
	"before": "af02839887791014f577198be60b73fa4cb9f23e",
	"after": "f8a3d8428d28ea7d3f6ec8b4e922cc15948c4bc0",
	"repository": {
		"id": 289352038,
		"node_id": "MDEwOlJlcG9zaXRvcnkyODkzNTIwMzg=",
		"name": "dummyTestRepo",
		"full_name": "teezzan/dummyTestRepo",
		"private": false,
		"owner": {
			"name": "teezzan",
			"email": "teehazzan@gmail.com",
			"login": "teezzan",
			"id": 25867172,
			"node_id": "MDQ6VXNlcjI1ODY3MTcy",
			"avatar_url": "https://avatars.githubusercontent.com/u/25867172?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/teezzan",
			"html_url": "https://github.com/teezzan",
			"followers_url": "https://api.github.com/users/teezzan/followers",
			"following_url": "https://api.github.com/users/teezzan/following{/other_user}",
			"gists_url": "https://api.github.com/users/teezzan/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/teezzan/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/teezzan/subscriptions",
			"organizations_url": "https://api.github.com/users/teezzan/orgs",
			"repos_url": "https://api.github.com/users/teezzan/repos",
			"events_url": "https://api.github.com/users/teezzan/events{/privacy}",
			"received_events_url": "https://api.github.com/users/teezzan/received_events",
			"type": "User",
			"site_admin": false
		},
		"html_url": "https://github.com/teezzan/dummyTestRepo",
		"description": null,
		"fork": false,
		"url": "https://github.com/teezzan/dummyTestRepo",
		"forks_url": "https://api.github.com/repos/teezzan/dummyTestRepo/forks",
		"keys_url": "https://api.github.com/repos/teezzan/dummyTestRepo/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/teezzan/dummyTestRepo/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/teezzan/dummyTestRepo/teams",
		"hooks_url": "https://api.github.com/repos/teezzan/dummyTestRepo/hooks",
		"issue_events_url": "https://api.github.com/repos/teezzan/dummyTestRepo/issues/events{/number}",
		"events_url": "https://api.github.com/repos/teezzan/dummyTestRepo/events",
		"assignees_url": "https://api.github.com/repos/teezzan/dummyTestRepo/assignees{/user}",
		"branches_url": "https://api.github.com/repos/teezzan/dummyTestRepo/branches{/branch}",
		"tags_url": "https://api.github.com/repos/teezzan/dummyTestRepo/tags",
		"blobs_url": "https://api.github.com/repos/teezzan/dummyTestRepo/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/teezzan/dummyTestRepo/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/teezzan/dummyTestRepo/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/teezzan/dummyTestRepo/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/teezzan/dummyTestRepo/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/teezzan/dummyTestRepo/languages",
		"stargazers_url": "https://api.github.com/repos/teezzan/dummyTestRepo/stargazers",
		"contributors_url": "https://api.github.com/repos/teezzan/dummyTestRepo/contributors",
		"subscribers_url": "https://api.github.com/repos/teezzan/dummyTestRepo/subscribers",
		"subscription_url": "https://api.github.com/repos/teezzan/dummyTestRepo/subscription",
		"commits_url": "https://api.github.com/repos/teezzan/dummyTestRepo/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/teezzan/dummyTestRepo/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/teezzan/dummyTestRepo/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/teezzan/dummyTestRepo/issues/comments{/number}",
		"contents_url": "https://api.github.com/repos/teezzan/dummyTestRepo/contents/{+path}",
		"compare_url": "https://api.github.com/repos/teezzan/dummyTestRepo/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/teezzan/dummyTestRepo/merges",
		"archive_url": "https://api.github.com/repos/teezzan/dummyTestRepo/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/teezzan/dummyTestRepo/downloads",
		"issues_url": "https://api.github.com/repos/teezzan/dummyTestRepo/issues{/number}",
		"pulls_url": "https://api.github.com/repos/teezzan/dummyTestRepo/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/teezzan/dummyTestRepo/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/teezzan/dummyTestRepo/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/teezzan/dummyTestRepo/labels{/name}",
		"releases_url": "https://api.github.com/repos/teezzan/dummyTestRepo/releases{/id}",
		"deployments_url": "https://api.github.com/repos/teezzan/dummyTestRepo/deployments",
		"created_at": 1598039133,
		"updated_at": "2022-01-20T18:58:00Z",
		"pushed_at": 1658612442,
		"git_url": "git://github.com/teezzan/dummyTestRepo.git",
		"ssh_url": "git@github.com:teezzan/dummyTestRepo.git",
		"clone_url": "https://github.com/teezzan/dummyTestRepo.git",
		"svn_url": "https://github.com/teezzan/dummyTestRepo",
		"homepage": null,
		"size": 21,
		"stargazers_count": 0,
		"watchers_count": 0,
		"language": null,
		"has_issues": true,
		"has_projects": true,
		"has_downloads": true,
		"has_wiki": true,
		"has_pages": false,
		"forks_count": 0,
		"mirror_url": null,
		"archived": false,
		"disabled": false,
		"open_issues_count": 12,
		"license": null,
		"allow_forking": true,
		"is_template": false,
		"web_commit_signoff_required": false,
		"topics": [],
		"visibility": "public",
		"forks": 0,
		"open_issues": 12,
		"watchers": 0,
		"default_branch": "master",
		"stargazers": 0,
		"master_branch": "master"
	},
	"pusher": {
		"name": "teezzan",
		"email": "teehazzan@gmail.com"
	},
	"sender": {
		"login": "teezzan",
		"id": 25867172,
		"node_id": "MDQ6VXNlcjI1ODY3MTcy",
		"avatar_url": "https://avatars.githubusercontent.com/u/25867172?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/teezzan",
		"html_url": "https://github.com/teezzan",
		"followers_url": "https://api.github.com/users/teezzan/followers",
		"following_url": "https://api.github.com/users/teezzan/following{/other_user}",
		"gists_url": "https://api.github.com/users/teezzan/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/teezzan/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/teezzan/subscriptions",
		"organizations_url": "https://api.github.com/users/teezzan/orgs",
		"repos_url": "https://api.github.com/users/teezzan/repos",
		"events_url": "https://api.github.com/users/teezzan/events{/privacy}",
		"received_events_url": "https://api.github.com/users/teezzan/received_events",
		"type": "User",
		"site_admin": false
	},
	"created": false,
	"deleted": false,
	"forced": false,
	"base_ref": null,
	"compare": "https://github.com/teezzan/dummyTestRepo/compare/af0283988779...f8a3d8428d28",
	"commits": [
		{
			"id": "f8a3d8428d28ea7d3f6ec8b4e922cc15948c4bc0",
			"tree_id": "32968a3123f409b332e460e9b4167cadf5fc221f",
			"distinct": true,
			"message": "addedd",
			"timestamp": "2022-07-23T22:40:36+01:00",
			"url": "https://github.com/teezzan/dummyTestRepo/commit/f8a3d8428d28ea7d3f6ec8b4e922cc15948c4bc0",
			"author": {
				"name": "Yusuf Taiwo Hassan",
				"email": "teehazzan@gmail.com",
				"username": "teezzan"
			},
			"committer": {
				"name": "Yusuf Taiwo Hassan",
				"email": "teehazzan@gmail.com",
				"username": "teezzan"
			},
			"added": [],
			"removed": [],
			"modified": [
				"test.txt"
			]
		}
	],
	"head_commit": {
		"id": "f8a3d8428d28ea7d3f6ec8b4e922cc15948c4bc0",
		"tree_id": "32968a3123f409b332e460e9b4167cadf5fc221f",
		"distinct": true,
		"message": "addedd",
		"timestamp": "2022-07-23T22:40:36+01:00",
		"url": "https://github.com/teezzan/dummyTestRepo/commit/f8a3d8428d28ea7d3f6ec8b4e922cc15948c4bc0",
		"author": {
			"name": "Yusuf Taiwo Hassan",
			"email": "teehazzan@gmail.com",
			"username": "teezzan"
		},
		"committer": {
			"name": "Yusuf Taiwo Hassan",
			"email": "teehazzan@gmail.com",
			"username": "teezzan"
		},
		"added": [],
		"removed": [],
		"modified": [
			"test.txt"
		]
	}
}`)
