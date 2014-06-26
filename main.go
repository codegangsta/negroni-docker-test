package main

import (
	_ "github.com/codegangsta/envy/autoload"
	"github.com/codegangsta/envy/lib"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	n := negroni.Classic()
	r := httprouter.New()

	r.GET("/", HelloWorld)

	n.UseHandler(r)
	n.Run(":" + envy.MustGet("PORT"))
}

func HelloWorld(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rw.Write([]byte("Hello Negroni"))
}
