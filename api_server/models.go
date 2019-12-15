package main

// User ...
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Task ...
type Task struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	UserID      string `json:"user_id,omitempty"`
}
