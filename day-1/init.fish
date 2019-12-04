#! /usr/bin/env fish

pwd | xargs basename | xargs go mod init
