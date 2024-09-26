# Integration with GitHub Actions

[GitHub Actions](https://docs.github.com/en/actions) is commonly used to implement CI/CD workflows. *Gosec* integrates with GitHub actions easily since they have a published Action on the GitHub Actions Marketplace. You will now learn how to set it up in a repository so that your codebase is scanned for security vulnerabilities every time something is commited to the repository.  

First we will need a repository that contains a go project, an example project we have set up can be downloaded with this command:

```
git clone https://github.com/lvainio/go-demo.git
cd go-demo/
```{{exec}}

Then we need to create the .github/workflows folder and a file that will contain our workflow:

`touch .github/workflows/main.yml`{{exec}}

Finally we need to insert the contents defining our workflow

```
echo "name: Run Gosec
on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
jobs:
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
          args: ./..." > .github/workflows/main.yml
```{{exec}}

This would then have to be commited and pushed to the remote repository but we are a bit limited in this tutorial since we don't want to give write access to our repo in this open tutorial. But for anyone that is interested to see it in action it would be possible to do the same steps as here and add the workflow to see how it works. 