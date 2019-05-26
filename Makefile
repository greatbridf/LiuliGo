source = main.go liuli/*.go
os = linux
arch = amd64
buildargs = GOOS=$(os)
buildargs += GOARCH=$(arch)
buildargs += CGO_ENABLED=1
targetname = LiuliGo

LiuliGo : $(source)
	go get -d -v ./...
	$(buildargs) go build -o $(targetname)

.PHONY : clean
clean :
	-rm $(targetname)
	-rm -rf caches/
	-rm *.log
	-rm index.db
