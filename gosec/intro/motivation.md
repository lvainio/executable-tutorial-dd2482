# Motivation

Security issues are often easier and less costly to fix when addressed early. Using a SAST tool such as *gosec* to automatically scan for vulnerabilities is one way to detect common security vulnerabilities before they become a larger problem. 

Furthermore, SAST tools are easily integrated in DevOps workflows as they can be configured to run automatically every time new code is commited to a repository. This allows for automatic detection of security vulnerabilities which are then reported back to the developers for fixing.

Security can become a bottleneck in DevOps environments when a separate security team has to manually review each new update. This can potentially delay deployment. If we instead integrate security directly into the pipeline we can ensure that deployments meet security standards while still being able to deploy quickly. 

While SAST tools like *gosec* will not detect every possible security vulnerability, they do provide an automated way of catching common vulnerabilities, which can save both time and money.

> Press *next* when you are ready to continue!