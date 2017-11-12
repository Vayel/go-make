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

.PHONY: run_slave1
run_slave1:
	@./bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40000 $(output)

.PHONY: run_slave2
run_slave2:
	@./bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40001 $(output)

.PHONY: run_master
run_master:
	./bin/master makefiles/$(filter-out run_master, $(MAKECMDGOALS)) $(target) $(masterport) $(output)

clean:
	@rm -r bin

%:
	@:
