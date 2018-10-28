OW_GOLANG=msciab/actionloop-golang-v1.11:1540726346

golang.done:
	wsk package update golang \
	&& touch golang.done

utils2.done:
	wsk package bind /whisk.system/utils utils2 \
	&& touch utils2.done

trigger.done:
	wsk trigger create golang-trigger
	touch trigger.done

invokedate: golang.done utils2.done
	wsk action update golang/invokedate invokedate.go -p action "utils2/date" --docker $(OW_GOLANG)

invokedate-test: invokedate
	wsk action invoke golang/invokedate -p message "Date is" -r

firetrigger: golang.done utils2.done trigger.done
	wsk action update golang/firetrigger firetrigger.go -p trigger "golang-trigger" --docker $(OW_GOLANG)
	wsk rule update golang-trigger-date  golang-trigger utils2/date

firetrigger-test: firetrigger
	wsk action invoke golang/firetrigger -p message "Date is" -r
