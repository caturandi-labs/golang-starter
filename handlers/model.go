package handlers

type registerRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email	string `json:"email"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email 	string	`json:"email"`
	Password string	`json:"password"`
}
