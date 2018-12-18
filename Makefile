source = main.go liuli/*.go

LiuliGo.cgi : $(source)
	cmd.exe /c "build.bat"

.PHONY : deploy
deploy : LiuliGo.cgi
	scp LiuliGo.cgi SS:html_root/interface/LiuliGo.cgi

.PHONY : test
test : LiuliGo.cgi
	cp LiuliGo.cgi /var/www/cgi_bin

.PHONY : clean
clean :
	-rm *.cgi
	-rm build
