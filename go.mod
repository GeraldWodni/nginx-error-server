module nginx-error-server

go 1.18

require github.com/GeraldWodni/kern.go v0.0.0-20220421165921-871eb65c3469

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gomodule/redigo v1.8.3 // indirect
	golang.org/x/sys v0.0.0-20210105210732-16f7687f5001 // indirect
)

replace github.com/GeraldWodni/kern.go => ../kern
