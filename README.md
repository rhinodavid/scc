# SCC

## Overview

This application computes the [strongly connected components](https://en.wikipedia.org/wiki/Strongly_connected_component) of a directed graph using [Kosaraju's two-pass algorithm](https://en.wikipedia.org/wiki/Kosaraju's_algorithm).

## Input format

See `test_graph.txt` for a sample input. Each line on the input file represents one edge of the graph.
Nodes are labeled by integers, and each edge consists of the label of the tail node followed by the label
of the head node.

## Usage

Run `go run main.go <GRAPH_FILE.txt>`. The output will be the population of the top five most populous
strongly connected components.
