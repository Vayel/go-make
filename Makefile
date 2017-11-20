target = all
output = outputfiles/
masteraddr = localhost
masterport = 10000
slaveaddr = localhost
slaveport = 40000
logfile = log.json

all: master slave

.PHONY: sequential
sequential:
	@mkdir -p bin
	@cp common.go sequential/ # All .go files must be in the same folder
	@cp parser.go sequential/
	go build -o bin/sequential sequential/*.go
	@rm sequential/common.go
	@rm sequential/parser.go

.PHONY: master
master:
	@mkdir -p bin
	@cp common.go master/ # All .go files must be in the same folder
	@cp parser.go master/
	go build -o bin/master master/*.go
	@rm master/common.go
	@rm master/parser.go

.PHONY: slave
slave:
	@mkdir -p bin
	@cp common.go slave/ # All .go files must be in the same folder
	go build -o bin/slave slave/*.go
	@rm slave/common.go

.PHONY: run_seq
run_seq:
	@./bin/sequential $(filter-out run_seq, $(MAKECMDGOALS)) $(target) $(logfile)

.PHONY: run_slave1
run_slave1:
	@./bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40000 $(output)

.PHONY: run_slave2
run_slave2:
	@./bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40001 $(output)

.PHONY: run_master
run_master:
	./bin/master $(filter-out run_master, $(MAKECMDGOALS)) $(target) $(masterport) $(output) $(logfile)

clean:
	@rm -r bin

%:
	@:
