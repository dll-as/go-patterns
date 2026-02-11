package main

import "fmt"

// Model
type User struct {
	ID   int
	Name string
}

// View interface
type UserView interface {
	Display(user User)
}

// Concrete View
type ConsoleUserView struct{}

func (v *ConsoleUserView) Display(user User) {
	fmt.Printf("User: %d, Name: %s\n", user.ID, user.Name)
}

// Presenter
type UserPresenter struct {
	view UserView
}

func NewUserPresenter(view UserView) *UserPresenter {
	return &UserPresenter{view: view}
}

func (p *UserPresenter) ShowUser(user User) {
	// Presenter can add logic before showing
	p.view.Display(user)
}

func main() {
	user := User{ID: 1, Name: "Alice"}
	view := &ConsoleUserView{}
	presenter := NewUserPresenter(view)

	presenter.ShowUser(user)
}
