#! /usr/bin/env fish

function ecod
    echo $argv
    eval $argv
    set rc $status
    return $rc
end

ecod goimports -w .
and ecod gofumpt -w .
and ecod golangci-lint run --no-config --disable-all --enable typecheck .
and ecod golangci-lint run --disable typecheck .
and ecod go test -failfast -timeout 1s -v
and ecod go run .
set rc $status
echo "(last rc: $status)"
exit $status
