
all:
ifndef ex
	$(error ex is not set. Usage: make ex=00)
endif
	cd ex$(ex); make

init:
	cd vagrant; docker-compose up

init_vagrant:
	cd vagrant; make vagrant_up