#! /usr/bin/env fish

fswatch . | terror | debounce | restart ./check.fish
