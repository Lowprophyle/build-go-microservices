check_install: 
	which swagger || go get -u github.com/go-swagger/go-swagger

swagger: check_install 
	swagger generate spec -o ./swagger.yaml --scan-models