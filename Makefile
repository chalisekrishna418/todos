
MAKE_ARGS:=

test:
	go test ./... -cover

run:
	grnc-yaml-bind && go build . && ./main -c config,env/config.yaml
