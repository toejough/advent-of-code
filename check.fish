#! /usr/bin/env fish

function ecod
    echo $argv
    fish -c "eval $argv"
    set rc $status
    return $rc
end

ecod goimports -w .
and ecod gofumpt -w .
and ecod golangci-lint run --no-config --disable-all --enable typecheck ./...
and ecod golangci-lint run --disable typecheck ./...
and ecod go test -failfast -timeout 1s -v ./...
and ecod 'cd day-1/part-1 && go run .'
and ecod 'cd day-1/part-2 && go run .'
and ecod 'cd day-2/part-1 && go run .'
and ecod 'cd day-2/part-2 && go run .'
and ecod 'cd day-3/part-1 && go run .'
and ecod 'cd day-3/part-2 && go run .'
set rc $status
echo "(last rc: $status)"
exit $status
