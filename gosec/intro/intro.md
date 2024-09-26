# Welcome! <img src="../assets/figure.png" style="width: 50px">

Welcome to this DevSecOps tutorial on *static application security testing* ([SAST](https://en.wikipedia.org/wiki/Static_application_security_testing)) with [gosec](https://github.com/securego/gosec).

# Intended Learning Outcomes

In this tutorial you will learn what SAST is and why it is important in a DevOps workflow to identify security vulnerabilities. You will also learn how to use a SAST tool in order to scan the source code of go projects for security vulnerabilities. We will cover the installation of *gosec* and guide you through different use cases. Finally, you will learn how *gosec* can be integrated in a CI workflow using GitHub Actions to automatically scan a project every time a commit is made to the repository.

# Static Application Security Testing (SAST)

The idea of SAST is to statically analyze the source code in order to find security vulnerabilities and insecure code patterns. SAST is done without having to build and execute the program so it is generally integrated in the code phase of DevOps (see figure below), allowing for quick feedback to the developers. There are many different SAST tools out there for different programming languages, and *gosec* is one such tool for the Go programming language. SAST tools are often integrated in DevSecOps pipelines as it is one of the cheapest and fastest ways to detect security vulnerabilities automatically. 

<img src="../assets/flowchart.png">
