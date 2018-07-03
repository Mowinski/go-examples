GOPLUGIN = go build -buildmode=plugin
ODIR = plugins
OUT_NAME = go-examples

all: json bytes print writer helloworld variables zerovalues conditional \
	forloop shortstatement switch shortswitch functions returnmanyvalues \
	returnfunction defer
	go build -o $(OUT_NAME)
	
create-structure:
	mkdir -p $(ODIR)
	
bytes: create-structure bytes/bytes.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

json: create-structure json/json.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

print: create-structure print/print.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

writer: create-structure writer/writer.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

helloworld: create-structure helloworld/helloworld.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go
	
variables: create-structure variables/variables.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

zerovalues: create-structure zerovalues/zerovalues.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

conditional: create-structure conditional/conditional.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

forloop: create-structure forloop/forloop.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

shortstatement: create-structure shortstatement/shortstatement.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

switch: create-structure switch/switch.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

shortswitch: create-structure shortswitch/shortswitch.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

functions: create-structure functions/functions.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

returnmanyvalues: create-structure returnmanyvalues/returnmanyvalues.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

returnfunction: create-structure returnfunction/returnfunction.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

defer: create-structure defer/defer.go
	${GOPLUGIN} -o $(ODIR)/$@.so $@/$@.go

	
clean:
	rm $(ODIR) -rf
	rm $(OUT_NAME)
