package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	port = ":8080"
	host = "http://localhost"
)

type Result struct {
	Ascii string
	Error string
}

func main() {
	// serve static files like css
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// set handle function for home
	http.HandleFunc("/", home)
	// set handle function for ascii-art
	http.HandleFunc("/ascii-art", asciiArt)
	// start server
	log.Printf("Staring application on port%s\n", port)
	log.Println("Open:", host+port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println(err)
		log.Fatalf("%v - internal server error", http.StatusInternalServerError)
	}

}

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	// check if url path is valid
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// reset data
	data := Result{}
	if render(w, "index.html", data) != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

// /ascii-art handler
func asciiArt(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	data := Result{}
	// get ascii-art
	err, value := createAsciiArt(r.FormValue("input"), r.FormValue("banner"))

	if err != nil {
		// set char error
		data = Result{
			Ascii: "",
			Error: err.Error(),
		}
	} else {
		// set ascii-art value
		data = Result{
			Ascii: value,
			Error: "",
		}
	}
	if render(w, "index.html", data) != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

// render template or send error if file doesn't exist
func render(w http.ResponseWriter, tmplName string, data Result) error {
	// create template
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		return err
	}
	// send template
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func createAsciiArt(input string, banner string) (error, string) {
	asciiInput := readFont(banner)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	splitWords := strings.Split(input, "\n")
	var result string
	if input != "" {
		for _, w := range splitWords {
			for i := 0; i < 9; i++ {
				asciiLine := ""
				for _, char := range w {
					if char < 32 || char > 126 {
						return errors.New("Character '" + string(char) + "' out of range for ascii-art."), ""
					} else {
						asciiLine += string(asciiInput[(int(char)-32)*9+i+1])
					}
				}
				result += asciiLine + "\n"
			}
		}
	}
	return nil, result
}

// read in the banner
func readFont(banner string) map[int]string {
	var fontRow map[int]string = map[int]string{}
	readFile, _ := os.Open("banners/" + banner + ".txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for row := 1; fileScanner.Scan(); row++ {
		fontRow[row] = fileScanner.Text()
	}
	readFile.Close()
	return fontRow
}
