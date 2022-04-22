// kern.go default backend meant to be used with ingress-nginx in kubernetes
// (c)copyright 2022 by Gerald Wodni <gerald.wodni@hmail.com>
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/GeraldWodni/kern.go"
    "github.com/GeraldWodni/kern.go/log"
    "github.com/GeraldWodni/kern.go/router"
    "github.com/GeraldWodni/kern.go/view"
)

func main() {
    // create new kern instance on port 5000
    app := kern.New(":5000", []string{ "content" } )

    LayoutViews, _ := app.Hierarchy.LookupDirectory( "views/layout" )

    errorView, err := view.NewHtml( append( LayoutViews, app.Hierarchy.LookupFatal( "views", "error.gohtml" ) )... )
    if err != nil {
        log.Fatal( "Cannot load error.gohtml", err )
    }
    app.Router.Get("/code/", func (res http.ResponseWriter, req *http.Request, next router.RouteNext ) {
        var statusCode StatusCode
        code := strings.TrimPrefix( req.URL.Path, "/code/" )
        fmt.Printf( code )
        statusCode = GetStatusCode( code )
        res.Header().Set("Content-Type", errorView.ContentType )
        errorView.Render( res, req, next, struct {
            StatusCode StatusCode
        }{
            StatusCode: statusCode,
        })
    })

    // mount index.gohtml on "/"
    app.Router.Get("/css/theme.css", view.NewCssHandler( app.Hierarchy.LookupFatal( "css", "theme.gocss" ) ) )
    app.Router.Get("/", view.NewHtmlHandler( append( LayoutViews, app.Hierarchy.LookupFatal( "views", "index.gohtml" ) )... ) )

    // start server
    app.Run()
}
