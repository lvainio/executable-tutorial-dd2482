# Welcome! <img src="../assets/figure.png" style="width: 50px">

Welcome to this DevSecOps tutorial on *static application security testing* ([SAST](https://en.wikipedia.org/wiki/Static_application_security_testing)) with [Gosec](https://github.com/securego/gosec).

# Intended Learning Outcomes

In this tutorial you will learn:

- What SAST is, why it is used, and how it can be used to automatically identify security vulnerabilities.

- How to use a SAST tool designed for the Go programming language, called *Gosec*.

- How to integrate *Gosec* with GitHub Actions.

We will first give some background information on SAST and *Gosec*. The first part of the tutorial will then cover installation of *Gosec. You will then be guided through various use cases to understand the capabilities of the tool. Finally, you will learn how *Gosec* can be integrated into a GitHub Actions CI workflow to automatically scan the source code of a Go project every time new code is commited to the GitHub repository.

# Static Application Security Testing (SAST)

The idea of SAST is to statically analyze source code in order to find security vulnerabilities. SAST is performed without having to build and execute the project, so it is generally integrated in the code phase of DevOps (see figure below), allowing for quick feedback to the developers. 

There are many different SAST tools out there for different programming languages. *Gosec* is one such tool for the Go programming language. SAST tools are often integrated in DevSecOps pipelines as it is one of the cheapest and fastest ways to detect security vulnerabilities automatically. 

<img src="../assets/flowchart.png">

# Gosec

You will also learn how to use a SAST tool called *Gosec*. *Gosec* is designed for the Go programming language and can identify common vulnerability patterns in source code, such as weak encryption and potential SQL injection vectors.

> Press *start* when you are ready to continue!