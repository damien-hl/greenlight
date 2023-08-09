package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	fmt.Fprintln(w, "create a new movie")
=======
	fmt.Fprintf(w, "create a new movie")
>>>>>>> e379ed7b8197e6a601634dc440d8d7e1dfae5220
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

<<<<<<< HEAD
	fmt.Fprintf(w, "show the details of movie %d\n", id)
=======
	fmt.Fprintf(w, "show the details of movie %d", id)
>>>>>>> e379ed7b8197e6a601634dc440d8d7e1dfae5220
}
