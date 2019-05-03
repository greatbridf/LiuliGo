source = main.go liuli/*.go
os = linux
arch = amd64
buildargs = CGO_ENABLED=0
buildargs += GOOS=$(os)
buildargs += GOARCH=$(arch)

LiuliGo.cgi : $(source)
	go get -d -v ./...
	$(buildargs) go build -o LiuliGo.cgi
  
.PHONY : deploy
deploy : LiuliGo.cgi
	tar czf - LiuliGo.cgi | ssh SS "tar xzf - && mv LiuliGo.cgi /var/www/interface/test/"

.PHONY : test
test : LiuliGo.cgi
	cp LiuliGo.cgi /var/www/cgi_bin

.PHONY : clean
clean :
	-rm *.cgi
	-rm -rf caches/
	-rm *.log
