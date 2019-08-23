package try_gqlgen

type User struct {
	ID   string
	Name string
	Age  int
}

type Todo struct {
	ID     string
	UserID string
	Text   string
	Done   bool
}
