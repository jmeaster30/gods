go test -v -coverprofile="coverage.text" .
go tool cover -html="coverage.text" -o "coverage.html"
firefox coverage.html