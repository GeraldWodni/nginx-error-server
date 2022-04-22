// kern.go default backend meant to be used with ingress-nginx in kubernetes
// (c)copyright 2022 by Gerald Wodni <gerald.wodni@hmail.com>
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"

    "github.com/GeraldWodni/kern.go"
    "github.com/GeraldWodni/kern.go/log"
    "github.com/GeraldWodni/kern.go/router"
    "github.com/GeraldWodni/kern.go/view"
)

var layoutViews []view.HtmlView
var errorView *view.HtmlView
var indexView *view.HtmlView

func renderStatusCode(res http.ResponseWriter, req *http.Request, next router.RouteNext, statusCode StatusCode) {
    res.Header().Set("Content-Type", errorView.ContentType )
    code, _ := strconv.Atoi( statusCode.Code )
    res.WriteHeader( code )
    errorView.Render( res, req, next, struct {
        StatusCode StatusCode
    }{
        StatusCode: statusCode,
    })
}

func defaultHandler(res http.ResponseWriter, req *http.Request, next router.RouteNext) {
    if xCode := req.Header.Get("X-Code"); xCode != "" {
        statusCode := GetStatusCode( xCode )
        code, _ := strconv.Atoi( statusCode.Code )
        if xFormat := req.Header.Get("X-Format"); xFormat != "" {
            if strings.Contains( xFormat, "text/html" ) {
                renderStatusCode( res, req, next, statusCode )
                return
            } else if strings.Contains( xFormat, "application/json" ) {
                res.Header().Set("Content-Type", "application/json")
                res.WriteHeader( code )
                fmt.Fprintf( res, "{\"err\":\"%s %s\"}", statusCode.Code, statusCode.MessageEN )
                return
            }
        }
        res.Header().Set("Content-Type", "text/plain")
        res.WriteHeader( code )
        fmt.Fprintf( res, statusCode.String() )
        return
    }

    res.Header().Set("Content-Type", indexView.ContentType )
    indexView.Render( res, req, next, nil )
}

func main() {
    // create new kern instance on port 5000
    app := kern.New(":5000", []string{ "content" } )

    layoutViews, _ := app.Hierarchy.LookupDirectory( "views/layout" )

    var err error
    errorView, err = view.NewHtml( append( layoutViews, app.Hierarchy.LookupFatal( "views", "error.gohtml" ) )... )
    if err != nil {
        log.Fatal( "Cannot load error.gohtml", err )
    }
    indexView, err = view.NewHtml( append( layoutViews, app.Hierarchy.LookupFatal( "views", "index.gohtml" ) )... )

    app.Router.Get("/code/", func (res http.ResponseWriter, req *http.Request, next router.RouteNext ) {
        code := strings.TrimPrefix( req.URL.Path, "/code/" )
        statusCode := GetStatusCode( code )
        renderStatusCode( res, req, next, statusCode )
    })

    // mount index.gohtml on "/"
    app.Router.Get("/css/theme.css", view.NewCssHandler( app.Hierarchy.LookupFatal( "css", "theme.gocss" ) ) )
    app.Router.Get("/", defaultHandler )

    // start server
    app.Run()
}
