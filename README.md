# Travelling Salesman Problem

This program written in Golang finds the shortest Hamiltonian-Path on a given graph.

## Usage

There is no actual CLI-Parser in place yet so the syntax may vary.
Calling `$ tsp` does something, depending on the configuration I am using at the moment. It is not unlikely though, that it will cause an error because I am using certain graphs to benchmark the program that are not available on your machine. Maybe I'll publish some of them.

Eventually the command is supposed to take a file path as an argument and have certain flags like the number of goroutines and the number of used OS threads as flags. Further there will be the option to chose between the concurrent and the parallel solution. So something along the lines of:
```
$ tsp <path> [--concurrent] [--max-progs <n>] [--go-routines <n>]
```
The files that this program uses are of the following form:
```
1 line: Number of vertices V
V lines: Names of vertices
1 line: Number of edges E
E lines: A B C describing an edge from A to B with the weight C
```

Due to this configuration those graphs do not have to be metric. The weights assigned can be arbitrary but have to be non-negative. (Even though a weight of 0 does not make a lot of sense)

## Implementation

As of 2019-01-22 the program uses a branch-and-bound approach to solve the problem. That means every instance contains a path, a count of edges already passed and a minimum cost of a roundtrip with this path at the start. Those instances are managed in a Priority Queue which is currently implemented as an array based binary heap. An instance is extended by generating all possible extensions of the given path by adding all possible available edges, recalculating the boundaries and finally inserting those instances back into the heap. Once the highest element on the heap is a roundtrip, we know it's the solution.

## Installation

It's as easy as:
```
$ go get -u github.com/miltfra/tsp
```
## Limitations

Even though the program has solved every problem so far, I cannot gurantee that it will continue to do so. This is because in its current state one goroutine might find a roundtrip on top of the heap while another is still calculting an extension of another, better path. This should not be a problem in most cases though.

Another quite big limitation is memory. It is not unusual that the program uses 20 Gigabytes of RAM. I've yet to find a solution for that problem because I've already used the smalles datatypes that are reasonable. That's why only 2^7 nodes are technically possible, because the referencing system uses int8.

## Algorithm

So, how does this work?

### The Problem 

Before we discuss the solution, we need to clarify the problem. What we want to do is to find the shortest path in any given graph that visits every single vertex exactly once and is also a cycle (i.e. the first vertex and the last node are identical).

In this case a path is a sequence of vertices. If given a graph with more than one edge between two vertices, it is trivial to create a subgraph with the same solution to the TSP but where there is at most one edge between two vertices.

Now, if `n` is the number of vertices, we can calculate how many possible options there are. If we can step through all these sequences of unique vertices of length `n`, we still have to check `n!` possibilites. The problem is that practically speaking we cannot know at any point that our solution is correct. 

Due to the sheer amount of possibilities it becomes extremely hard to find optimal solutions on bigger and especially denser graphs.

### The Solution

The way we approach this problem is using lower boundaries. We calculate a value which represents a definitive minimum length for each possible path that has a certain sub path.

To do this, we start by calculating the lowest weight of an outgoing edge for every vertex. 
Now we can add all of these together. We can do the same for incoming edges. 
Since both of these values are definitve lower boundaries for the length of the shortest path, we only need to consider the higher value.

Now we have our first sub path. It contains exactly one vertex and no edge. This path may be extended by all the neighbours of the last visited node (the only one so far). Now we have created a maximum of `n-1` new candidates. For all of those we can calculate the definitve lower boundary again. The only difference is that, at this point in time, we have one node where we know the outgoing weight and one node where we know the incoming weight for sure. 

Since this decision might change the lower boundary, we may now take the candidate with the lowest minimum cost. If there are two candidates with the same boundary but different lengths, we take the shorter one, since more visited nodes means the boundary is more likely to be the actual final value.

At some point we will end up with a "candidate" that has exactly n edges and the cost of which is at most as much as another candidates lower boundary. Thus there is no other path which is shorter.

The problem with the implementation is not only the algorithm, but efficiency as well. You may have a look at the code and notice how I had to micro manage memory and efficiency.
