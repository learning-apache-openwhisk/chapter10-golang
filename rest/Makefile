OW_GOLANG=msciab/actionloop-golang-v1.11:1540726346

golang.done:
	wsk package update golang \
	&& touch golang.done

utils2.done:
	wsk package bind /whisk.system/utils utils2 \
	&& touch utils2.done

trigger.done:
	-wsk trigger create golang-trigger
	wsk rule update golang-trigger-date golang-trigger utils2/date
	touch trigger.done

invokesort: golang.done utils2.done
	wsk action update golang/invokesort invokesort.go \
	-p action "utils2/sort" --main invoke --docker $(OW_GOLANG)

invokesort-test: invokesort
	wsk action invoke golang/invokesort -p message "Sorted: " -p text beta,alpha,gamma -r

triggersort.zip: invokesort.go triggersort.go
	zip firetrigger.zip invokesort.go triggersort.go

triggersort: golang.done utils2.done trigger.done triggersort.zip 
	wsk action update golang/triggersort firetrigger.zip  --main fire -p trigger "golang-trigger"  --docker $(OW_GOLANG)

#triggersort-test: firetrigger
#	wsk action invoke golang/firetrigger -p message "Date is" -r

test: golang.done utils2.done trigger.done
	#go test