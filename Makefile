source = main.go liuli/*.go
os = linux
arch = amd64
buildargs = GOOS=$(os)
buildargs += GOARCH=$(arch)
buildargs += CGO_ENABLED=1

LiuliGo.cgi : $(source)
	go get -d -v ./...
	$(buildargs) go build -o LiuliGo.cgi

.PHONY : deploy
deploy : LiuliGo.cgi
	tar czf - LiuliGo.cgi | ssh SS "tar xzf - && mv LiuliGo.cgi /var/www/interface/test/"

.PHONY : test
test : LiuliGo.cgi
	REQUEST_URI=?req=articles ./LiuliGo.cgi

.PHONY : clean
clean :
	-rm *.cgi
	-rm -rf caches/
	-rm *.log
	-rm index.db
