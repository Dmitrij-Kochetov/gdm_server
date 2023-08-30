build:
	go build -o bin/gdm_server -v -race cmd/gdm_server/main.go

clear:
	rm -rf bin/*

rebuild:
	go build -o bin/gdm_server -a -v -race cmd/gdm_server/main.go 


.PHONY:
	build-server rebuild	