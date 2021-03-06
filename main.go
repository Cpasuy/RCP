package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"regexp"
	"text/template"

	_ "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//var templates = template.Must(template.ParseFiles("index.html"))
var conection *sql.DB
var err error

func main() {
	conection, err = sql.Open("mysql", "root:@/rcp")
	if err != nil {
		panic(err.Error())
	}
	err = conection.Ping()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/login", loadLogin)
	http.HandleFunc("/register", registerComp)
	http.HandleFunc("/index", loadUsers)
	http.HandleFunc("/", loadLogin)
	http.ListenAndServe(":8080", nil)
}

var u string

func registerComp(w http.ResponseWriter, r *http.Request) {
	alerta := Message{}
	alerta.Msg = "Fill all the boxes"
	if r.Method != "POST" {
		http.ServeFile(w, r, "register.html")
		return
	}
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmpwd := r.FormValue("confirmpwd")
	country := r.FormValue("country")
	fmt.Println("PasswordSign " + password)
	err := conection.QueryRow("SELECT Username FROM usuarios WHERE username=?", username).Scan(&u)
	if err == sql.ErrNoRows {
		if (len(username) > 5 && len(username) < 20) || (len(password) > 0 && len(confirmpwd) > 0){
			if len(country) > 0 {
				if password == confirmpwd {
					hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
					if err != nil {
						http.Error(w, "Error", 500)
						return
					}
					em := isEmailValid(email)
					if em == true {
						sendEmail(firstname, email)
						_, err = conection.Exec("INSERT INTO usuarios(Username, Firstname, Lastname, Email,Password,Country) VALUES(?, ?,?,?,?,?)", username, firstname, lastname, email, hash, country)
					} else{
						alerta.Msg = "Digit a valid email"
						t, _ := template.ParseFiles("register.html")
						t.Execute(w, alerta)
						return
					}

				} else{
					alerta.Msg = "Passwords do not match"
					t, _ := template.ParseFiles("register.html")
					t.Execute(w, alerta)
					return
				}
			} else{
				alerta.Msg = "Digit a contry"
				t, _ := template.ParseFiles("register.html")
				t.Execute(w, alerta)
				return
			}
		} else{
			alerta.Msg = "Digit a username and password valids"
			t, _ := template.ParseFiles("register.html")
			t.Execute(w, alerta)
			return
		}
	}  else{
			alerta.Msg = "The username already exists"
			t, _ := template.ParseFiles("register.html")
			t.Execute(w, alerta)
			return
	}
	http.Redirect(w, r, "/", 301)
	return
}
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

var us string
var pw string

func loadLogin(w http.ResponseWriter, r *http.Request) {
	alerta := Message{}
	alerta.Msg = "Welcome, digit your user and password"
	if r.Method != "POST" {
		http.ServeFile(w, r, "login.html")
		return
	}
	username := r.FormValue("username")
	
	password := r.FormValue("password")

	err := conection.QueryRow("SELECT Username, Password FROM usuarios WHERE Username=? ", username).Scan(&us, &pw)
	if err != nil {
		alerta.Msg = "Digit a user and password valids"
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, alerta)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	if err != nil {
		alerta.Msg = "Digit a user and password valids"
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, alerta)
		return
	}
	http.Redirect(w, r, "/index", 301)
}

func loadUsers(w http.ResponseWriter, r *http.Request) {
	comps, err := conection.Query("SELECT Username, Firstname, Lastname FROM usuarios")

	if err != nil {
		http.Error(w, "ERROR", 500)
		return
	}
	competitor := Competitor{}
	competitors := []Competitor{}
	for comps.Next() {
		var u string
		var f string
		var l string
		err = comps.Scan(&u, &f, &l)
		if err != nil {
			http.Error(w, "ERROR", 500)
			return
		}
		competitor.Username = u
		competitor.Firstname = f
		competitor.Lastname = l

		competitors = append(competitors, competitor)

	}

	fmt.Println(competitors)
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, competitors)
}

type Competitor struct {
	Firstname, Lastname, Username string
}

type Message struct {
	Msg string
}

type Dest struct {
	Name string
}

func sendEmail(f string, e string) {
	from := mail.Address{"RCP", "soportercpicesi@gmail.com"}
	to := mail.Address{f, e}
	subject := "Enviando correo desde GO"
	dest := Dest{Name: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("template.html")
	checkErr(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "soportercpicesi@gmail.com", "icesi123", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	err = client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	client.Quit()
}
func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
