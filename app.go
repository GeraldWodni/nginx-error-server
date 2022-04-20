// kern.go default backend meant to be used with ingress-nginx in kubernetes
// (c)copyright 2022 by Gerald Wodni <gerald.wodni@hmail.com>
package main

import (
    "boolshit.net/kern"
    "boolshit.net/kern/view"
)

func main() {
    // create new kern instance on port 5000
    app := kern.New(":5000", []string{ "content" } )

    // mount index.gohtml on "/"
    app.Router.Get("/css/theme.css", view.NewCssHandler( app.Hierarchy.LookupFatal( "css", "theme.gocss" ) ) )
    app.Router.Get("/", view.NewHtmlHandler( app.Hierarchy.LookupFatal( "views", "index.gohtml" ) ) )

    // start server
    app.Run()
}
