package main

type GithubWebhook struct {
	HookId     int64 `json:"hook_id"`
	Repository struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	}
}
