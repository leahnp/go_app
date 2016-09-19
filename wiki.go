package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
)

// allows users to view a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request) {
  // extracts page title from path
  // slices path to drop view
    title := r.URL.Path[len("/view/"):]
    // loads page data
    p, _ := loadPage(title)
    // formats page to html
    // writes to w
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

// this struct describes how
// the data will be stored in memory
type Page struct {
    Title string
    Body  []byte
}

// save method gives the app
// persistent storage

// This method's signature reads: 
// "This is a method named save that takes as
// its receiver p, a pointer to Page . 
// It takes no parameters, and returns a
// value of type error."
func (p *Page) save() error {
    filename := p.Title + ".txt"
    // will return nil if all goes well
    // 0600 (octal intger literal) makes file read/write 
    // for current user
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// this function constructs the file
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    // io.ReadFile returns []byte and error
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    // if second param is nil, page successfully loaded
    return &Page{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    // function reads contents of file and returns
    // a *template
    t, _ := template.ParseFiles("edit.html")
    // executes the template, writing html to 
    // http.ResponseWriter
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}