# Go-Tree

A basic ```tree``` implementation in golang.

Tested on OS X.

build: ```go build tree.go```

usage:
```
$./tree .

|-- ./gowiki/
|   |-- .tree.go.swp
|   |-- MyPage.txt
|   |-- MyPages.txt
|   |-- TestPage.txt
|   |-- TestPages.txt
|   |-- edit.html
|   |-- gowiki/foo
|   |   |-- hello
|   |   |-- s_hello -> ../../hello/hello.go
|   |   |-- s_link -> ../../hello/
|   |-- tree
|   |-- view.html
|   |-- web_server
|   |-- web_server.go
|   |-- wiki
|   |-- wiki.go
```
