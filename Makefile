
MAKE_ARGS:=

test:
	go test ./... -cover

run:
	grnc-yaml-bind && go build . && ./todos -c config,env/config.yaml
