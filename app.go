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

    // extend globals (available to all view-templates)
    view.Globals["AppName"] = "kern.go nginx-error-server (default-backend)"

    // mount index.gohtml on "/"
    app.Router.Get("/", view.NewHandler( app.Hierarchy.LookupFatal( "views", "index.gohtml" ) ) )

    // start server
    app.Run()
}
