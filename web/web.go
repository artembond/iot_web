package web

import (
	_ "bufio"
	"fmt"
	_ "log"
	_ "os"
	_ "reflect"
	_ "strconv"
	_ "strings"
	_ "sync"
	_ "time"

	"html/template"
	"io/ioutil"
	"net/http"
	//"regexp"
	//"io"

	_ "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	_ "github.com/fatih/color"
	_ "github.com/tarm/goserial"
	"golang.org/x/net/websocket"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	p, err := loadPage("aaa")
	if err != nil {
		http.Redirect(w, r, "/edit/"+"aaa", http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	p, err := loadPage("aaa")
	if err != nil {
		p = &Page{Title: "aaa"}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	body := r.FormValue("body")
	//fmt.Println(body)
	p := &Page{Title: "aaa", Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+"aaa", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("111111111111111111")
	//p, err := loadPage("aaa")
	//p, err := &Page{Title: "aaa"}
	//	if err != nil {
	//		http.Redirect(w, r, "/edit/"+"aaa", http.StatusFound)
	//		return
	//	}
	//	Title := "TITLE"
	t, _ := template.ParseFiles("w/index.html")
	items := struct {
		Title string
		Test2 string
		Test3 string
	}{
		Title: "MyName",
		Test2: "MyCity",
		Test3: "Test3",
	}

	t.Execute(w, items)
	//renderTemplate(w, "index", p, Title)
}

var templates = template.Must(template.ParseFiles("w/edit.html", "w/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		m := validPath.FindStringSubmatch(r.URL.Path)
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
//		fn(w, r, m[2])
//	}
//}

func Init() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/index.html", indexHandler)

	http.Handle("/ws_test", websocket.Handler(EchoServer))

	fmt.Println("web start ")
	http.ListenAndServe(":8081", nil)
}
