# Welcome!

Welcome to this DevSecOps tutorial on *static application security testing* ([SAST](https://en.wikipedia.org/wiki/Static_application_security_testing)) with [gosec](https://github.com/securego/gosec). <br>

# Intended Learning Outcomes

In this tutorial you will learn what SAST is and why it is important in a DevOps workflow to identify security vulnerabilities. You will also learn how to use a SAST tool in order to scan the source code of go projects for security vulnerabilities. We will cover the installation of *gosec* and guide you through different use cases. Finally, you will learn how *gosec* can be integrated in a CI workflow using GitHub Actions to automatically scan a project every time a commit is made to the repository.

# Static Application Security Testing (SAST)



# Motivation

Scanning the source code for security vulnerabilities every time new code is added helps us identify potential issues early in the development process. Security issues are often far easier and more cost-effective to address early on. Using a SAST tool such as *gosec* is one way to detect common security vulnerabilities early, before they become a bigger problem.

Security can become a bottleneck in DevOps workflows when separate security teams have to manually review each new update. This can potentially delay deployment which goes against the nature of fast and automatic deployments in DevOps. If we instead integrate security directly into the pipeline we can ensure that the deployments meet security standards. While SAST tools like *gosec* will not detect every possible security vulnerability, they do provide an automated way of catching common vulnerabilites early on, which can save both time and money.