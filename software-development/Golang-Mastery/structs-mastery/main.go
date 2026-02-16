package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Password string
}

//creating a login function which is a method of User

func (u User) Login() string {
	return u.Email
}

func (u *User) UpdateUser(newEmail string, newPassword string) string {
	u.Name = newEmail
	u.Password = newPassword

	return fmt.Sprintf("%s %s", u.Name, u.Password)
}

func main() {
	user := User{"Mubarak", "mubaraklouis@gmail.com", "mubarak12"}

	//login the user
	fmt.Println(user.Login())

	//change the password and email

	fmt.Println(user.UpdateUser("johndoe@gmail.com", "johndoe123"))

}
