#! /usr/bin/env fish

# first check if there's anything to even commit
if git diff --cached --quiet
    echo "nothing to commit!"
    exit 1
end

# now choose from the options
set TYPE (gum choose "fix" "feat" "docs" "style" "refactor" "test" "dev" "revert")
set SCOPE (gum input --placeholder "scope")

# Since the scope is optional, wrap it in parentheses if it has a value.
test -n "$SCOPE" && set SCOPE "($SCOPE)"

# Pre-populate the input with the type(scope): so that the user may change it
set SUMMARY (gum input --prompt "$TYPE$SCOPE: " --placeholder "Imperative commit message" --width 72)
set DESCRIPTION (gum write --placeholder "Details of this change (CTRL+D to finish)" --width 80)

# Commit these changes
gum confirm "Commit changes?" && git commit -m "$TYPE$SCOPE: $SUMMARY" -m "$DESCRIPTION"
