package types

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Total    int    `json:"total"`
}

type User struct {
	ID       string   `json:"id"`
	Access   string   `json:"access"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Issued   []string `json:"issued"`
}
