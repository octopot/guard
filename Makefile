PACKAGE := github.com/kamilsk/guard
SECRET  := 10000000-2000-4000-8000-160000000000
VERSION := latest

include env/make/cmd.mk
include env/make/common.mk
include env/make/docker.mk
include env/make/docker-compose.mk
include env/make/tools.mk
