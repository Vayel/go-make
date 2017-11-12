# go-make

A parallel version of the `make` tool in Go which uses RPC. Works with
basic Makefiles only.

# Commands

## Install

```bash
make
```

## Run master

```bash
make run_master fname-in-makefiles-dir
```

For instance:

```bash
make run_master 8
```

## Run slave

```bash
make run_slave1
# Or: make run_slave2
```

# Makefiles
This tool handles only basic Makefiles. All the rules must look like:
```
target: dep1 dep2 dep3
	instruction1
	instruction2
```

And all the rules must create a file of the same name.
