OW_GOLANG=openwhisk/actionloop-golang-v1.11

golang.done:
	wsk package update golang \
	&& touch golang.done

utils2.done:
	wsk package bind /whisk.system/utils utils2 \
	&& touch utils2.done

trigger.done: utils2.done
	wsk trigger create golang-trigger
	wsk rule update golang-trigger-sort golang-trigger utils2/sort
	wsk rule enable golang-trigger-sort
	touch trigger.done

invokesort: golang.done utils2.done
	wsk action update golang/invokesort invokesort.go \
	-p action "utils2/sort" --main invoke --docker $(OW_GOLANG)

invokesort-test: invokesort
	wsk action invoke golang/invokesort -p message "Sorted: " -p text beta,alpha,gamma -r

triggersort.zip: invokesort.go triggersort.go
	zip triggersort.zip invokesort.go triggersort.go

triggersort: golang.done utils2.done trigger.done triggersort.zip 
	wsk action update golang/triggersort triggersort.zip  --main fire \
	-p trigger "golang-trigger"  \
	-p retrieve "golang/triggersort" \
	--docker $(OW_GOLANG)

triggersort-test: triggersort
	wsk action invoke golang/triggersort -p text "sort,me,out,please" -r

test: golang.done utils2.done trigger.done
	#go test

clean:
	-wsk action delete golang/invokesort
	-wsk action delete golang/triggersort
	-wsk package delete utils2
	-wsk trigger delete golang-trigger
	-wsk rule delete golang-trigger-sort
	-rm utils2.done trigger.done
	-rm *.zip
	