# Using `gosec`

### Run `gosec`

Execute the following command to run `gosec` on all Go files in the current directory:

`gosec ./... {{exec}}`

This command scans all Go files in the directory and its subdirectories.

#### Run `gosec` with Advanced Options

Execute the following command to run `gosec` on all Go files in the current directory:

`gosec ./... -severity medium {{exec}}`

This command scans all Go files in the directory and its subdirectories based on the severity level of the vulnerabilities, and filter out the issues with a lower severity than the given value.

### Review the Output

After running the command, `gosec` will provide output that lists any potential security issues found in your code.
