GO ?= go

build:
	@echo "Building dist/ooni"
	@$(GO) build -i -o dist/ooni cmd/ooni/main.go
.PHONY: build

build-windows:
	@echo "Building dist/ooni.exe"
	CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o dist/ooni.exe -x cmd/ooni/main.go

update-mk-libs:
	@echo "updating mk-libs"
	@cd vendor/github.com/measurement-kit/go-measurement-kit && curl -L -o master.zip https://github.com/measurement-kit/golang-prebuilt/archive/master.zip && unzip master.zip && mv golang-prebuilt-master libs && rm master.zip # This is a hack to workaround: https://github.com/golang/dep/issues/1240
.PHONY: update-mk-libs

bindata:
	@$(GO) run vendor/github.com/shuLhan/go-bindata/go-bindata/*.go \
		-nometadata	\
		-o internal/bindata/bindata.go -pkg bindata \
	    data/...;
.PHONY: bindata
