OW_USER?=msciab
OW_RUNTIME?=$(OW_USER)/actionloop
OW_COMPILER?=$(OW_USER)/actionloop-golang-v1.11
WSK?=wsk
MAIN=main
PACKAGE=golang
SRCS=src/$(MAIN).go src/app/assets.go src/app/app-packr.go
VENDORS=src/app/vendor
NAME=webapp
BINZIP=$(MAIN)-bin.zip
SRCZIP=$(MAIN)-src.zip

deploy: package.done $(BINZIP)
	$(WSK) action update $(PACKAGE)/$(NAME) $(BINZIP) --main $(MAIN) --docker $(OW_RUNTIME) --web true
	$(WSK) action get $(PACKAGE)/$(NAME) --url

src/%/vendor:
	cd $(@D) ; dep ensure

src/app/app-packr.go:
	cd src/app ; packr

$(BINZIP): $(SRCS) $(VENDORS)
	cd src ; packr ; zip - -r * | docker run -i $(OW_COMPILER) -compile $(MAIN) >../$(BINZIP)

$(SRCZIP): $(SRCS) $(VENDORS)
	cd src ; zip ../$(SRCZIP) -r *

clean:
	wsk action delete $(PACKAGE)/$(NAME)
	rm  $(BINZIP) $(SRCZIP) package.done

test: test.json
	$(WSK) action invoke test/$(NAME) -r
	$(WSK) action invoke test/$(NAME) -P test.json -r

test.json:
	echo '{ "name": "Mike" }' >test.json

package.done:
	$(WSK) package update $(PACKAGE)
	touch package.done

.PHONY: deploy devel test clean
