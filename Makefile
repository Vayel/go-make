target = all
masteraddr = localhost
masterport = 10000
slaveaddr = localhost
slaveport = 40000
WORKING_DIR = tmp
SRC_DIR = src

all: master slave

.PHONY: prepare_compile
prepare_compile:
	@mkdir -p bin

.PHONY: sequential
sequential: prepare_compile
	@cp $(SRC_DIR)/common.go $(SRC_DIR)/sequential/ # All .go files must be in the same folder
	@cp $(SRC_DIR)/parser.go $(SRC_DIR)/sequential/
	go build -o bin/sequential $(SRC_DIR)/sequential/*.go
	@rm $(SRC_DIR)/sequential/common.go
	@rm $(SRC_DIR)/sequential/parser.go

.PHONY: master
master: prepare_compile
	@cp $(SRC_DIR)/common.go $(SRC_DIR)/master/ # All .go files must be in the same folder
	@cp $(SRC_DIR)/parser.go $(SRC_DIR)/master/
	go build -o bin/master $(SRC_DIR)/master/*.go
	@rm $(SRC_DIR)/master/common.go
	@rm $(SRC_DIR)/master/parser.go

.PHONY: slave
slave: prepare_compile
	@cp $(SRC_DIR)/common.go $(SRC_DIR)/slave/ # All .go files must be in the same folder
	go build -o bin/slave $(SRC_DIR)/slave/*.go
	@rm $(SRC_DIR)/slave/common.go

.PHONY: prepare_run
prepare_run:
	@mkdir -p $(WORKING_DIR)

.PHONY: run_seq
run_seq: prepare_run
	@cd $(WORKING_DIR); \
	../bin/sequential $(filter-out run_seq, $(MAKECMDGOALS)) $(target) sequential.json

.PHONY: run_slave1
run_slave1: prepare_run
	@cd $(WORKING_DIR); \
	../bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40000 .

.PHONY: run_slave2
run_slave2: prepare_run
	@cd $(WORKING_DIR); \
	../bin/slave $(masteraddr) $(masterport) $(slaveaddr) 40001 .

.PHONY: run_master
run_master: prepare_run
	@cd $(WORKING_DIR); \
	../bin/master $(filter-out run_master, $(MAKECMDGOALS)) $(target) $(masterport) master.json

clean:
	@rm -r bin
	@rm -r $(WORKING_DIR)

%:
	@:
