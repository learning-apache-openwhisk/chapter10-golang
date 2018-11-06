OW_USER?=openwhisk
OW_RUNTIME?=$(OW_USER)/actionloop
OW_COMPILER?=$(OW_USER)/actionloop-golang-v1.11
WSK?=wsk
MAIN=main
NAME=calc
PACKAGE=golang
SRCS=$(wildcard src/*.go src/*/*.go)
VENDORS=$(wildcard src/*/Gopkg.toml)
BINZIP=$(MAIN)-bin.zip
SRCZIP=$(MAIN)-src.zip

deploy: package.done $(BINZIP)
	$(WSK) action update $(PACKAGE)/$(NAME) $(BINZIP) \
	--main $(MAIN) --docker $(OW_RUNTIME) 

devel: package.done $(SRCZIP)
	$(WSK) action update $(PACKAGE)/$(NAME) $(SRCZIP) \
	--main $(MAIN) --docker $(OW_COMPILER) 

unit-test:
	cd src ; GOPATH=$$PWD/.. go test *

src/%/vendor:
	cd $(@D) ; dep ensure

$(BINZIP): $(SRCS) $(VENDORS)
	cd src ; packr ; zip - -r * | docker run -i $(OW_COMPILER) -compile $(MAIN) >../$(BINZIP)

$(SRCZIP): $(SRCS) $(VENDORS)
	cd src ; zip ../$(SRCZIP) -r *

clean:
	wsk action delete $(PACKAGE)/$(NAME)
	rm  $(BINZIP) $(SRCZIP) package.done

test:
	$(WSK) action invoke $(PACKAGE)/$(NAME) -p expr "2 + 2" | awk  '{print $$6}' | tee last.id
	$(WSK) activation logs $$(cat last.id) 
	$(WSK) activation result $$(cat last.id) 
	$(WSK) action invoke $(PACKAGE)/$(NAME) -p expr "3 + 3" | awk  '{print $$6}' | tee last.id
	$(WSK) activation logs $$(cat last.id) 
	$(WSK) activation result $$(cat last.id) 

package.done:
	$(WSK) package update $(PACKAGE)
	touch package.done

.PHONY: deploy devel test clean
