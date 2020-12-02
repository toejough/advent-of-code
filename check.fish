#! /usr/bin/env fish

function echo-do
    echo $argv
    eval $argv
    set rc $status
    return $rc
end

echo-do goimports -w .
and go mod tidy
and echo-do gofumpt -w .
and echo-do golangci-lint run --no-config --disable-all --enable typecheck ./...
and echo-do golangci-lint run --disable typecheck ./...
and echo-do go test -failfast -timeout 1s -v ./...
and echo-do go run .
set rc $status
echo "(last rc: $status)"
exit $status
