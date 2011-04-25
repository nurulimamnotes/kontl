include $(GOROOT)/src/Make.inc

all: clean kontl

TARG=kontl
GOFILES=\
	kontl.go\

include $(GOROOT)/src/Make.cmd

