# Advanced *Gosec* features

### Severity and Confidence

Talk about how in large programs SAST tools might detect a lot of potential vulnerabilities. Depending on requirements and other things like time and money it might only be necessary or possible to focus on the most important and severy vulnerabilities. Gosec therefore have a way of filtering these out. Show example of gosec showing ALOT of vulnerabilities and then changes setting to only show high severity and such to show that it is now clearer.

Execute the following command to run `gosec` on all Go files in the current directory:

`gosec -severity medium ./...`{{exec}}

This command scans all Go files in the directory and its subdirectories based on the severity level of the vulnerabilities, and filter out the issues with a lower severity than the given value.

### Handling false positives

talk about how SAST tools will always have false positives. It might be that the tool detects something that isnt actually a vulnerability or it might be that the vulnerable behaviour is intended. Show example of a false positive. Fix it in the code. Then run gosec again to show that it now does not highlight it. 