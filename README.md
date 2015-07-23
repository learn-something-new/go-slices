# Slices

[![logo](https://raw.githubusercontent.com/learn-something-new/learn-something-new.github.io/master/logo.png?v=2)](#)

A very short example of slices in [Go](https://golang.org/). Copied from my examples of [common data structures](https://github.com/obihann/GoDataStructures/stack), I am using the stack as a way to implement slices.


## Table of Contents
* [Arrays](/array/src)
* [Copy](/copy/src)
* [Append](/append/src)

## Arrays

A stack built with a traditional array. You will notice this as a fixed size.

## Copy

A stack built with slices implementing the copy method. This is more efficient than a traditional array as it allows you to easily grow the array by making larger arrays and copying your contents over.

## Append

A stack built with append. Append allows us to grow the array as we need, using as little memory as possible.

##License
This tool is protected by the [GNU General Public License v2](http://www.gnu.org/licenses/gpl-2.0.html).

Copyright [Jeffrey Hann](http://jeffreyhann.ca/) 2014
