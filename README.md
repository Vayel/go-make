# go-make

A parallel version of the `make` tool in Go which uses RPC. Works with
basic Makefiles only.

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
