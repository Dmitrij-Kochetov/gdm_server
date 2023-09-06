build:
	go build -o bin/gdm_server -v -race cmd/gdm_server/main.go

clear:
	rm -rf bin/*

rebuild:
	go build -o bin/gdm_server -a -v -race cmd/gdm_server/main.go 

migrate:
	go-bindata -prefix "migrations" -pkg migrations -o /home/dmitrij/projects/gdm_server/internal/gdm_server/adapter/database/migrations/migrations.go /home/dmitrij/projects/gdm_server/migrations

.PHONY:
	build-server rebuild	