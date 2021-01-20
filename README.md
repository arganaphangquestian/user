### Add this for letter at `model.go`

```go
// InputRole model
type InputRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Role model
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UserRole model
type UserRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
```