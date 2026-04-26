package main

type Message struct {
	Time    string `json:"time"`
	Channel string `json:"channel"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type result struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
