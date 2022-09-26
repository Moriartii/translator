export $(cat .env | egrep -v "(^#.*|^$)" | xargs)
mkdir ~/logs
go run ./cmd/main.go 2>&1 | tee -a ~/logs/log.txt

