package webhook_test


var GitlabBody= []byte(
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