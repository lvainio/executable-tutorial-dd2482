# Static Application Security Testing (SAST)

The idea of SAST is to statically analyze source code in order to find security vulnerabilities. SAST is performed without having to build and execute the project, so it is generally integrated in the code phase of DevOps (see figure below), allowing for quick feedback to the developers. 

There are many different SAST tools out there for different programming languages. SAST is often a part of DevSecOps pipelines as it is one of the cheapest and fastest ways to detect security vulnerabilities automatically. 

<img src="../assets/flowchart.png">

# Gosec

*Gosec* is a SAST tool designed for the Go programming language. It works by scanning the *abstract syntax tree* ([AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)) and the *static single assignment* ([SSA](https://en.wikipedia.org/wiki/Static_single-assignment_form)) representations of the source code. It can identify a list of common vulnerabilities such as the use of broken hash functions and potential SQL injection vectors.

> Press *NEXT* when you are ready to continue!