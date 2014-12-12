Simple pattern for using a json file as a database assuming you don't have too many users or too much data.

## Build

	go get
	go build

## Run

	./jsondb-go

## Usage
	$ echo '{}' > db.json 

	$ curl -i -X POST localhost:8080/todo -d '{"value":"hello world"}'
	HTTP/1.1 201 Created
	Content-Type: application/json
	Location: todo/fc366812-3440-4767-a4cd-1130f0f61219
	Date: Fri, 12 Dec 2014 21:32:53 GMT
	Content-Length: 67

	{"id":"fc366812-3440-4767-a4cd-1130f0f61219","value":"hello world"}

	$ curl -i -X POST localhost:8080/todo -d '{"value":"hello world2"}'
	HTTP/1.1 201 Created
	Content-Type: application/json
	Location: todo/657859f9-c591-4a47-9c98-ebd9f1f7afa4
	Date: Fri, 12 Dec 2014 21:33:02 GMT
	Content-Length: 68

	{"id":"657859f9-c591-4a47-9c98-ebd9f1f7afa4","value":"hello world2"}

	$ curl -i localhost:8080/todo
	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Fri, 12 Dec 2014 21:33:07 GMT
	Content-Length: 138

	[{"id":"657859f9-c591-4a47-9c98-ebd9f1f7afa4","value":"hello world2"},{"id":"fc366812-3440-4767-a4cd-1130f0f61219","value":"hello world"}]

	$ curl -i localhost:8080/todo/fc366812-3440-4767-a4cd-1130f0f61219
	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Fri, 12 Dec 2014 21:33:15 GMT
	Content-Length: 67

	{"id":"fc366812-3440-4767-a4cd-1130f0f61219","value":"hello world"}