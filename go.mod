module nginx-error-server

go 1.23

require github.com/GeraldWodni/kern.go v0.0.0-20220812153137-f58b263924d6

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gomodule/redigo v1.9.2 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

replace github.com/GeraldWodni/kern.go => ../kern
