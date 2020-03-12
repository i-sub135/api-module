run:
ifdef url
	curl $(url) > .env && go run httpd/main.go
else
	@echo "masukan URL .env Anda \nContoh : \n - make run url=http://bla/bla/bla.env \n\n" && go run httpd/main.go
endif

build:
	go build httpd/*.go


autostart:
ifdef url
	curl $(url) > .env && reflex -r '\.go' -s -- sh -c 'go run httpd/main.go'
else
	@echo "masukan URL .env Anda \nContoh : \n - make run url=http://bla/bla/bla.env\n\n" && reflex -r '\.go' -s -- sh -c 'go run httpd/main.go'
endif

init:
	go get -v  -d ./...