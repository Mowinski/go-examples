GOPLUGIN = go build -buildmode=plugin

all: json bytes print writer
	go build
	
build:
	mkdir build
	
bytes: build bytes/bytes.go
	${GOPLUGIN} -o build/bytes.so bytes/bytes.go

json: build json/json.go
	${GOPLUGIN} -o build/json.so json/json.go

print: build print/print.go
	${GOPLUGIN} -o build/print.so print/print.go

writer: build writer/writer.go
	${GOPLUGIN} -o build/writer.so writer/writer.go

clean:
	rm build -rf
	rm golang-workshop-presentation
