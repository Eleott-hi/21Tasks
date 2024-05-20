
all:
ifndef ex
	$(error ex is not set. Usage: make ex=00)
endif
	cd ex$(ex); make

init:
	cd vagrant; make vagrant_up