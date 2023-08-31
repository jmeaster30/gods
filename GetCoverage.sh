go test -v -coverprofile="coverage.text" ./linear ./composite
go tool cover -html="coverage.text" -o "coverage.html"
firefox coverage.html