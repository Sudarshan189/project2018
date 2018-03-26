package main

type StatusMsg struct {
	Status int `json:"status, omitempty"`
	Message string `json:"message, omitempty"`
}

var (
	DatabaseError = StatusMsg{0, "Can't connect to database"}
	CreatedSuccess = StatusMsg{1, "creation successful"}
	CreatedFailed = StatusMsg{2, "failed to create"}
	LoggedInSuccess= StatusMsg{3, "successfilly logged in"}
	LoggedInFail= StatusMsg{4, "id and password doesnot match"}
)