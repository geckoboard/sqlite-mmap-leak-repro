.PHONY: debug pmap

cities.db:
	wget https://s3.amazonaws.com/gecko.static/geo_cities-201503.db -O cities.db

debug: $(GOPATH)/bin/dlv ./cities.db
	dlv debug ./ --output __sqlite_debug

pmap:
	watch 'pmap -X $$(pgrep -n __sqlite_debug) | grep deleted | wc -l'

$(GOPATH)/bin/dlv:
	cd $GOPATH && go get github.com/go-delve/delve/cmd/dlv
