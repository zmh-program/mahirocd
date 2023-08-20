package main

import "encoding/json"

type GithubWebhook struct {
	HookId     int64 `json:"hook_id"`
	Repository struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	}
}

func (g *GithubWebhook) ToBytes() []byte {
	if data, err := json.Marshal(g); err != nil {
		return nil
	} else {
		return data
	}
}
