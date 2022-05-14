package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func asciiartHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// fmt.Fprintf(w, "POST request successful")
	input := r.FormValue("input")
	banner := r.FormValue("drone")
	path := "Banners/" + banner + ".txt"
	fmt.Fprintf(w, "%s\n", print(input, path))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/asciiart", asciiartHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// func main() {

// implementing our newline function on the argument which is the string after main.go

// Also logging errors if too many or too little arguments.

func banner(str string) map[int][]string {
	file, err := os.Open(str)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	// Returns new Scanner
	scanner.Split(bufio.ScanLines)
	// scans line by line
	var txtlines []string
	// create new slice of string variable
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	// append scanned lines (line by line) to new slices of string
	file.Close()

	asciimap := make(map[int][]string)
	num := 0

	asciival := 32
	// starting ascii value of 'space'
	emptslice := []string{}
	// empty slice to append too, when pulling lines.

	// loop up to the ascii value of 126.
	for asciival <= 126 {
		section := (9 * num)
		// variable to break up code into seections/lines of 9.
		for j := section; j < section+9; j++ {
			// loops through the lines
			emptslice = append(emptslice, txtlines[j])
		}
		asciimap[asciival] = emptslice
		// this maps the value to emptslice. e.g. num = 1, ascii value 32 = space.
		emptslice = []string{}
		// after successfully mapping. Then emptying the variable which is mapped too.
		num++
		// every loop num goes up. which is the next group of lines.
		asciival++
		// as the 'num' does go up every loop, so does the 'ascii value' we are focusing on.

	}

	return asciimap
	// returning the mapped value
}

func print(str string, b string) string {
	m := banner(b)
	// associating value 'm' with the mapping function, with the standard.txt. file.
	art := ""
	// empty variable 'art'

	for i := 1; i < 9; i++ {
		// loop start at line 1 up to line 8
		for j := range str {
			// ranging through the string
			art += (m[int(str[j])][i])
		}
		// mapping into empty 'art', 'm' casted as an integar the
		art += ("\n")
	}
	return art
}

/*
func newline(str string) {
	splitstr := strings.Split(str, "\\n")
	b := "Banners/" + banner + ".txt"
	for i := 0; i < len(splitstr); i++ {
		if splitstr[i] == "" {
			fmt.Print("\n")
		} else {
			fmt.Print(print(splitstr[i], b))
		}
	}
}
*/
