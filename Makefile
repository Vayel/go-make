target = all
output = outputfiles/
masteraddr = localhost
masterport = 10000
slaveaddr = localhost
slaveport = 40000

all: master slave

.PHONY: master
master:
	@mkdir -p bin
	@cp common.go master/ # All .go files must be in the same folder
	go build -o bin/master master/*.go
	@rm master/common.go

.PHONY: slave
slave:
	@mkdir -p bin
	@cp common.go slave/ # All .go files must be in the same folder
	go build -o bin/slave slave/*.go
	@rm slave/common.go

.PHONY: run_slave
run_slave:
	@./bin/slave $(masteraddr) $(masterport) $(slaveaddr) $(slaveport) $(output)

.PHONY: run_master
run_master:
	./bin/master makefiles/$(filter-out run_master, $(MAKECMDGOALS)) $(target) $(masterport) $(output)

clean:
	@rm -r bin

%:
	@:
