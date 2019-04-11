# btree-util

Ever came accross a scenario where you want to write an alogorithm on Binary Tree but had to write code to create a Tree first? Ever wondered how your binary tree looks after some manipulations? You would have traversed the tree and printed to see it's current state but wouldn't it be cool to see the Tree in a Tree shape! 

Btree-util does the job for you for all you Go programm Binary tree needs. Btree-util is a tiny package that helps you initialize a Binary tree with choice of data so that you can get started with applying your Btree algorithms right away. What's cool is you can graphically see how your tree looks like at a given point in your flow

## Getting started

Follow the instructions here [https://golang.org/doc/install](https://golang.org/doc/install) if you don't have Go installed on your machine already

In your choice of terminal, execute `$ go get github.com/anandkilli/btree-util/btree` to install btree-util Go package

## How it works

Check the included example program `examples/deleteNode.go` to jump start

The package includes the following 4 functions

1. [Init](#init)
2. BtreeToHtml
3. HPrintln
4. DrawBtree

## Init

This method accepts a string representation of a Binary tree and create the tree for you
