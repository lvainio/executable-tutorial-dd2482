# Integration with GitHub Actions

[GitHub Actions](https://docs.github.com/en/actions) is commonly used to implement CI/CD workflows. *Gosec* integrates easily with GitHub actions and they even have a published Action on the GitHub Actions Marketplace. 

In this section you will learn how to set up automatic security scans with *Gosec* in in a GitHub repository.

### Downloading a Go project

First we will need a GitHub repository that contains a Go project. For demonstration purposes we have created a simple Go project that can be cloned by running the following command:

```
git clone https://github.com/lvainio/go-demo.git
cd go-demo/
ls
```{{exec}}

As you can see in the terminal, the repository contains a readme, a simple Go program, and the go.mod file. The program can be executed via the following command:

`go run main.go`{{exec}}

### Creating the GitHub Action workflow

Now that we have cloned the repository we need to create the GitHub Actions workflow. In GitHub repositories, all workflows are defined in YAML files stored in the `.github/workflows/` directory. GitHub will automatically recognize and run workflows in this directory depending on what event the workflow is configured to trigger on. Run the following command to create the YAML file in the workflows directory:

`mkdir -p .github/workflows && touch .github/workflows/gosec-security-scan.yml`{{exec}}

Then we need to define when this workflow should be executed. In this example we want the workflow to automatically execute when new code is pushed to the main branch of the repository. The `on` keyword is used to tell GitHub which event we want the workflow to trigger on. In our case it will be the `push` event on the `main` branch. Run the following command to add the event trigger in our workflow:

```
echo "name: Gosec Security Scan
on:
  push:
    branches:
      - main" > .github/workflows/gosec-security-scan.yml
```{{exec}}

Then we need to define what the workflow should do. We do this by adding a job to the workflow. As can be seen in the command right below, the job contains two steps. The first step just downloads the repository and the second step uses the *Gosec* Action (from GitHub Actions Marketplace) to scan the repository for security vulnerabilities. Remember that the argument *./...* tells *Gosec* to scan files in the current directory and all sub directories.

```
echo "jobs:
  tests:
    runs-on: ubuntu-latest
    env:
      GO111MODULE: on
    steps:
      - name: Checkout Source
        uses: actions/checkout@v3
      - name: Run Gosec Security Scanner
        uses: securego/gosec@master
        with:
          args: ./..." >> .github/workflows/gosec-security-scan.yml
```{{exec}}

> If you want to view the entire workflow file you can run:
>
> `cat .github/workflows/gosec-security-scan.yml`{{exec}}

### Pushing the workflow to github.com

The workflow is now completed and the final step would be to push the changes to the remote repository on `github.com`. We cannot really demonstrate that in this limited tutorial environment as it would require permission to push changes to our remote repository. If you are curious and want to see a successful run and an unsuccessful run of this workflow on `github.com` you can view these links:

- [successful]() 
- [unsuccessful]()





