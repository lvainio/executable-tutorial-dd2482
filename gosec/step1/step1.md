# Installation

In order for us to begin with the tutorial we first need a few things. A script is running in the console which updates the version of Go and downloads several Go programs which we will scan for security vulnerabilities.

We will of course need to install *gosec* as well. This can be done in many different ways but one of the easiest ways to do it is to run the script below:

`go install github.com/securego/gosec/v2/cmd/gosec@latest`{{exec}}

> Please wait for the *gosec* installation to finish, it might take a little while. 

`cat executable-tutorial-dd2482/gosec/step1/text.txt`{{exec}}