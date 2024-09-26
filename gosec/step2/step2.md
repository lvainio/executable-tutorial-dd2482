# Using `gosec`

> Ensure that you have Go installed on your machine.

**Run `gosec`**: Execute the following command to run `gosec` on all Go files in the current directory:

```bash
gosec ./... {{exec}}
```

This command scans all Go files in the directory and its subdirectories.

---

**Review the Output**: After running the command, `gosec` will provide output that lists any potential security issues found in your code.
