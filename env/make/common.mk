.PHONY: help
help: #| Shows available help information of Makefile.
	@fgrep -h "#|" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/#|//'
