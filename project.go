package main

import ("fmt"
        "net/http")

type User struct {
    name string
    age uint16
    money int16
    avg_grades, happiness float64
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is %s. " +
    "He is %d and he has money equal: %d", u.name, u.age, u.money)
}
func (u *User) setNewName(newName string) {
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
  stepan := User{"Stepan", 21, 5000, 4.5, 0.9}
  stepan.setNewName("Stefan")
  fmt.Fprintf(w, stepan.getAllInfo())
}
func help_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "It is a help page")
}

func handleReguest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/help/", help_page)
  http.ListenAndServe(":8080", nil)
}

func main() {

  //bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.6}


  handleReguest()
}
