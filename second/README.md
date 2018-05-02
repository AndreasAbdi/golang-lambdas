<3

how to build aws lambda compatible go executables. 

1. Gotta make sure your make's got 
`
env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/service service/main.go
`
or
`
GOOS=linux GOARCH=amd64 go build -o main main.go
`
This is cuz the flags are needed to have nonstatic links and other shit. 
That makes bigly sense doe, cuz like its a lambda ffs.

2. Zip and upload
`
go build main.go
zip main.zip main
`

