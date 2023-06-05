# Contribution Guide

## Table of Contents

- [Introduction](#introduction)
- [Expectations](#expectations)
- [What You Can Do](#what-you-can-do)
  * [Open Issues](#open-issues)
  * [Help The Development](#help-the-development)
    + [Example Contribution: Implement Operators](#example-contribution--implement-operators)
    + [Git Flow](#git-flow)
  * [Improve Documentation](#improve-documentation)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

## Introduction

Thank you for your desire to contribute. Although this is a toy programming language from a book called [Writing An Interpreter in Go](https://www.goodreads.com/book/show/32681092-writing-an-interpreter-in-go), you can always contribute to the future to the repository.

To set our expectations together, continue reading to `Expectations` section below.

## Expectations

This language is a toy language, and it shouldn't be used for any production system, especially for commercial products. 

That being said, this language has so many learning opportunities. Either you're learning on how to create a programming language, or you just want to understand on how to add new operator (either prefix or infix operators), this repository is a good place to do so.

## What You Can Do

### Open Issues

If you find a bug or have a suggestion to improve this repository, feel free to raise an issue. For now, I haven't set up a proper labelling system, but that shouldn't deter you from reporting bugs or suggesting new features.

### Help The Development

You can also help the development process by raising a pull request. There are many ways you can improve this toy language, e.g. by implementing new operators or implementing loop mechanism (either [sentinel-controlled loop or counter-controlled loop](https://www.geeksforgeeks.org/difference-between-sentinel-and-counter-controlled-loop-in-c/)).

#### Example Contribution: Implement Operators

If you read the book, Monkey Language doesn't have operators that we usually use, e.g. a built-in boolean operators (`&&` and `||`). For the purpose of learning how the interpreter works, that's not needed, However, it's a good start to "go out of the book" and trying things on your own.

For boolean operators, I implemented those operators as example on how to implement infix operators. Example PR can be found here: https://github.com/iamdejan/monkey-lang/pull/38

In order to implement new operators, you need to modify 3 modules (in sequence):
1) `lexer`: you need to create the token;
2) `parser`: you need to parse the token to abstract syntax tree (AST), as well as either defining the precedence of the operator(s), or re-use. In theory, if the operator is not an arithmetic operator, you can put it anywhere. However, please think twice about the precedence; and
3) `evaluator`: this module is where the AST is evaluated. Depending on the need, you can either create new function(s), or reuse existing functions;

#### Git Flow

You need to fork this repository, then submit a PR. You can read [this article](https://www.atlassian.com/git/tutorials/git-forks-and-upstreams) for guide about Git remote(s) and upstream branch(es).

### Improve Documentation

Documentation is always helpful to people stopping by and looking at this repository.

You can improve documentation by:
- improving `README.md` file;
- create and/or improving wiki; and/or
- raise a PR to create `godoc`;
