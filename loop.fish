#! /usr/bin/env fish

# Funcs
# ----------

# Echo & Do
function echo-do
    echo $argv
    eval $argv
    set rc $status
    return $rc
end

# Checks
function check
    echo-do goimports -w .
    and echo-do gofumpt -w .
    and echo-do golangci-lint run --no-config --disable-all --enable typecheck ./...
    and echo-do golangci-lint run --disable typecheck ./...
    and echo-do go test -failfast -timeout 1s -v ./...
    and echo-do go run .
    set rc $status
    echo "(last rc: $status)"
    return $status
end

# Arg Validation
# ----------

if test (count $argv) -ne 1;
    echo "must specify one argument: the path to the part to run"
    exit 1
end

set PART $argv[1]

if test ! -d $PART;
    echo "$PART must be a directory"
    exit 1
end

# Run
# ----------
set CALLING_DIR (pwd)
echo-do cd $PART
fswatch . | terror | debounce | restart $CALLING_DIR/check.fish
