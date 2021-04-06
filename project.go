package main

import ("fmt"; "net/http"; "html/template")

type User struct {
    Name string
    Age uint16
    Money int16
    Avg_grades, Happiness float64
    Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is %s. " +
    "He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
  u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
  ildar := User{"Ildar", 21, 5000, 4.5, 0.9, []string{"Programming", "Music", "Deutch language"}}
  //stepan.setNewName("Stefan")
  //fmt.Fprintf(w, stepan.getAllInfo())
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, ildar)
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
