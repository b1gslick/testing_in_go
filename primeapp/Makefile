default:
	go run .

test:
	go test -v -cover -coverprofile=coverage.out . || true

testc:
		go test -v -cover -coverprofile=coverage.out . || true
		go tool cover -html=coverage.out
	

