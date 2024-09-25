# Welcome!

Welcome to this DevSecOps tutorial on *static application security testing* ([SAST](https://en.wikipedia.org/wiki/Static_application_security_testing)) with [gosec](https://github.com/securego/gosec). <br>

# Indended Learning Outcomes

In this tutorial you will learn what SAST is and why it is important in a DevOps workflow to identify security vulnerabilities. You will also learn how to use a SAST tool in order to scan the source code of go projects for security vulnerabilities. We will cover the installation of *gosec* and guide you through different use cases. Finally, you will learn how *gosec* can be integrated in a CI workflow using GitHub Actions to automatically scan a project every time a commit is made to the repository.

# Motivation

Scanning the source code for vulnerabilities every time new code is added allows us to find security vulnerabilities early in the development process. Issues are often far easier and cheaper to fix early on, and using a SAST tool like *gosec* is one way to detect vulnerabilities early.

Security is often a bottleneck in DevOps workflows completely separate security teams have to manually review each new update. This can potentially delay deployment and goes against one of the core ideas of DevOps, which is to automate 

