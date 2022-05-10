if not exist pack md pack
if not exist pack\static md pack\static
if not exist pack\http md pack\http
if not exist pack\http\static md pack\http\static

go get -u github.com/intmian/mian_go_lib
go mod tidy
go build -o pack
copy setting_tep.json pack\setting.json
copy static pack\static
copy http\static pack\http\static
msg * "suc"