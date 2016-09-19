package main

import (
  "fmt"
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
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
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
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}

// compiles and executes code, creates testpage.tct
// func main() {
//     p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//     p1.save()
//     p2, _ := loadPage("TestPage")
//     fmt.Println(string(p2.Body))
// }
// version with views
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}