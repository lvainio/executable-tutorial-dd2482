# Motivation

Scanning the source code for security vulnerabilities every time new code is added helps us identify potential issues early in the development process. Security issues are often far easier and more cost-effective to address early on. Using a SAST tool such as *gosec* is one way to detect common security vulnerabilities early, before they become a bigger problem.

Security can become a bottleneck in DevOps workflows when separate security teams have to manually review each new update. This can potentially delay deployment which goes against the nature of fast and automatic deployments in DevOps. If we instead integrate security directly into the pipeline we can ensure that the deployments meet security standards while still being able to deploy quickly. While SAST tools like *gosec* will not detect every possible security vulnerability, they do provide an automated way of catching common vulnerabilites early on, which can save both time and money.

> Press *start* when you are ready to continue!