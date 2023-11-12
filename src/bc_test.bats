#!/usr/bin/env bats

# Load the script to be tested
source ./src/bc


# Tests for the add function
@test "add function adds all changes to staging area" {
    run add
    [ "$status" -eq 0 ]
    [ "$output" = "" ]
}

# Tests for the is_main_branch function
@test "is_main_branch returns true for main branch" {
    local result=$(is_main_branch "main")
    [ "$result" = "true" ]
}

@test "is_main_branch returns false for non-main branch" {
    local result=$(is_main_branch "feature-branch")
    echo $result
    [ "$result" = "false" ]
}

# Tests for the get_current_branch_name function
@test "get_current_branch_name returns the correct branch name" {
    local result=$(get_current_branch_name)
    [ "$result" = "main" ]  # Change this to the expected current branch name
}