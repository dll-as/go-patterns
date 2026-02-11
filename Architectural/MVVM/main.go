package main

import "fmt"

// Model
type User struct {
	ID   int
	Name string
}

// ViewModel
type UserViewModel struct {
	UserName string
}

func NewUserViewModel(user User) *UserViewModel {
	return &UserViewModel{
		UserName: fmt.Sprintf("User: %s (ID: %d)", user.Name, user.ID),
	}
}

// View
type ConsoleView struct{}

func (v *ConsoleView) Render(vm *UserViewModel) {
	fmt.Println(vm.UserName)
}

func main() {
	user := User{ID: 1, Name: "Alice"}
	vm := NewUserViewModel(user)

	view := &ConsoleView{}
	view.Render(vm)
}
