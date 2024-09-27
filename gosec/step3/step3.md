# Integration with GitHub Actions

[GitHub Actions](https://docs.github.com/en/actions) is commonly used to implement CI/CD workflows. *Gosec* integrates easily with GitHub actions and they even have a published Action on the GitHub Actions Marketplace. 

In this section you will learn how to set up automatic security scans with *Gosec* in in a GitHub repository.

### Downloading a Go project

First we will need a repository that contains a Go project. For demonstration purposes we have created a simple example project that can be downloaded by running the following command.

```
git clone https://github.com/lvainio/go-demo.git
cd go-demo/
ls -la
```{{exec}}

As you can see in the terminal, the repository contains a *README*, the Go program, 

### Creating the GitHub Action workflow

Now that we have downloaded a repository and cd:d into it we have to add a workflow so that GitHub automatically runs gosec on our repo when new code is commited. On GitHub, workflows need to be stored in the *.github/workflows* folder so that GitHub can automatically detect and run them, so we create a workflow file called *main.yml* in this folder with the following command: 

`touch .github/workflows/main.yml`{{exec}}

Finally, we need to define what this workflow should do. In this simple example we want the workflow to run when code is pushed to the main branch so we use the *on* keyword to tell GitHub to trigger this workflow on this event.

```
echo "name: Run Gosec
on:
  push:
    branches:
      - main"
```{{exec}}

Then we need to add the job to the workflow. It will contain a step that downloads the repository containing the go project and one step that runs gosec on the project.

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
          args: ./..." >> .github/workflows/main.yml
```{{exec}}

If you want to view the entire file you can run:

`cat .github/workflows/main.yml`

This workflow would then have to be commited and pushed to the remote repository but we are a bit limited in this tutorial since we don't want to give write access to our repo in this open tutorial. But for anyone that is interested to see it in action it would be possible to do the same steps as here and add the workflow to see how it works. 