#! /usr/bin/make -f

# GOROOT may need to be set if Go's installed somewhere unusual,
#
# PATH needs to contain $GOROOT/bin.
#
# GOPATH needs to be set, e.g. export GOPATH=/<path to wnu_csv>
#
# To build from source for local architecture do:
#     make var check all
#
# To build from source for new executable suitable for deployment, e.g.
# for a different architecture, do:
#     make var check deploy
#
# where deploy makes for 64 bit linux (if on osx)

SHELL = /bin/bash

all:
	go install wnu_csv   # Native executable in ../bin/wnu_csv.

var:
	[[ -d $$GOPATH ]]
	[[ -f $$GOPATH/src/wnu_csv/wnu_csv.go ]]

check:
	go vet wnu_csv
	go test wnu_csv

fmt:
	go fmt wnu_csv

deploy:
	GOOS=linux GOARCH=amd64 go install -installsuffix deploy wnu_csv

clean:
	rm -f ../bin/wnu_csv
	rm -f ../bin/*/wnu_csv
