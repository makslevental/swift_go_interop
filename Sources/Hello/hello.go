package main

/*
struct Vertex {
    int X;
    int Y;
};
*/
import     "C"

import (
    "fmt"
 	"fyne.io/fyne/widget"
 	"fyne.io/fyne/app"
    "net/http"
    "github.com/go-chi/chi"
 )

//export getVertex
func getVertex(X, Y C.int) C.struct_Vertex {
    return C.struct_Vertex{X, Y}
}

//export takeVertex
func takeVertex(v C.struct_Vertex) {
    fmt.Println(v.X + v.Y)
}

//export helloWindow
func helloWindow() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

//export helloChi
func helloChi() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	fmt.Println("serving at http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

//export helloGo
func helloGo() {
    fmt.Println("Hello from Go!")
}
func main() {}