#!/bin/sh

# FUNCTION
function sayHello() {
	echo "Hello world"
}
sayHello

# FUNCTION WITH PARAMS
function greet() {
	echo "Hello, I'm $1 and I'm $2 "
}

greet "Noureddine" "22"
