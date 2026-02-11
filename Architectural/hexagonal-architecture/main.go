package main

import "fmt"

// Domain / Core
type User struct {
	ID   int
	Name string
}

// Port: Repository interface
type UserRepository interface {
	FindByID(id int) (User, error)
}

// Use Case / Service
type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) string {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "Unknown"
	}
	return user.Name
}

// Adapter: In-memory repository
type InMemoryUserRepo struct {
	users []User
}

func (r *InMemoryUserRepo) FindByID(id int) (User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func main() {
	repo := &InMemoryUserRepo{
		users: []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		},
	}

	service := NewUserService(repo)

	fmt.Println("User 1:", service.GetUserName(1))
	fmt.Println("User 3:", service.GetUserName(3))
}
