package main

// User ...
type User struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Task ...
type Task struct {
	ID          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
}
