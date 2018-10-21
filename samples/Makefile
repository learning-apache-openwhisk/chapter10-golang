PACKAGE=golang
OW_RUNTIME=msciab/actionloop
OW_COMPILER=msciab/actionloop-golang-v1.11
WSK?=wsk

echo: package.done
	$(WSK) action update $(PACKAGE)/echo echo.go --docker $(OW_COMPILER)

echo-test: echo
	$(WSK) action invoke $(PACKAGE)/echo -p name Mike -r

hello: package.done
	$(WSK) action update $(PACKAGE)/hello hello.go --main hello --docker $(OW_COMPILER)

hello-test: hello
	$(WSK) action invoke $(PACKAGE)/hello -r
	$(WSK) action invoke $(PACKAGE)/hello -p name Mike -r

datetime: package.done
	zip datetime.zip datetime.go util.go
	$(WSK) action update $(PACKAGE)/datetime datetime.zip --main datetime --docker $(OW_COMPILER)

datetime-test: datetime
	$(WSK) action invoke $(PACKAGE)/datetime -p format "2001-01-02 3:4:5" -r
	$(WSK) action invoke $(PACKAGE)/datetime -r

package.done:
	wsk package update $(PACKAGE)
	touch package.done