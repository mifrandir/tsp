# Travelling Salesman Problem

This program written in Golang finds the shortest Hamiltonian-Path on a given graph.

## Usage

There is no actual CLI-Parser in place yet so the syntax may vary.
Calling `$ tsp` does something, depending on the configuration I am using at the moment. It is not unlikely though, that it will cause an error because I am using certain graphs to benchmark the program that are not available on your machine. Maybe I'll publish some of them.

Eventually the command is supposed to take a file path as an argument and have certain flags like the number of goroutines and the number of used OS threads as flags. Further there will be the option to chose between the concurrent and the parallel solution. So something along the lines of:
```
$ tsp <path> [--concurrent] [--max-progs <n>] [--go-routines <n>]
```

## Implementation

As of 2019-01-22 the program uses a branch-and-bound approach to solve the problem. That means every instance contains a path, a count of edges already passed and a minimum cost of a roundtrip with this path at the start. Those instances are managed in a Priority Queue which is currently implemented as an array based binary heap. An instance is extended by generating all possible extensions of the given path by adding all possible available edges, recalculating the boundaries and finally inserting those instances back into the heap. Once the highest element on the heap is a roundtrip, we know it's the solution.

## Installation

It's as easy as:
```
$ go get github.com/miltfra/tsp
```
This should - in theory - install my [tools](https://github.com/miltfra/tools) package as well.

## Limitations

Even though the program has solved every problem so far, I cannot gurantee that it will continue to do so. This is because in its current state one goroutine might find a roundtrip on top of the heap while another is still calculting an extension of another, better path. This should not be a problem in most cases though.

Another quite big limitation is memory. It is not unusual that the program uses 20 Gigabytes of RAM. I've yet to find a solution for that problem because I've already used the smalles datatypes that are reasonable. That's why only 2^7 nodes are technically possible, because the referencing system uses int8.
