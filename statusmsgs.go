package main

type StatusMsg struct {
	Status int `json:"state, omitempty"`
	Message string `json:"message, omitempty"`
}

var (
	DatabaseError = StatusMsg{0, "Can't connect to database"}
	EmailSent= StatusMsg{8, "Email Sent Successfully" }
	CreatedSuccess = StatusMsg{1, "creation successful"}
	//CreatedFailed = StatusMsg{2, "failed to create"}
	LoggedInSuccess= StatusMsg{3, "successfully logged in"}
	LoggedInFail= StatusMsg{4, "id and password doesnot match"}
	DatabaseRetriveError= StatusMsg{5, "cannot retrive data"}
	WrongPassword= StatusMsg{6, "Old password do not match"}
	PasswordChanged= StatusMsg{7, "Password Changed Successfully"}
)

