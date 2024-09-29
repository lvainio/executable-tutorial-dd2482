# Advanced *Gosec* features

In this part of the tutorial you will learn about some other useful functionalities of *Gosec*.

### Severity and Confidence

By default, Gosec reports all potential vulnerabilities it detects, regardless of their severity level or the certainty of whether they are true positives (i.e., actual vulnerabilities). Depending on the security requirements of the project, you may want to filter out vulnerabilities based on a minimum severity and confidence level. This approach means that the found vulnerabilities are more likely to be actual security issues and since these have a higher severity it is most likely these issues that you would want to focus on first. 

In *Gosec* there are three severity levels and three confidence levels: low, medium and high. The following command tells *Gosec* to find all security vulerabilities of a minimum severity level of medium and a minimum confidence level of medium:

`gosec -severity medium -confidence medium ./...`{{exec}}

You can see several issues were found, but if we instead change the minimum severity level to high we find only one issue that is of high severity:

`gosec -severity high -confidence medium ./...`{{exec}}

### Handling false positives

With all SAST tools there is the issue of false positives. A false positive is a reported vulnerability that does not actually pose a security risk. If a reported security issue is manually confirmed to be safe we may want to mark it to prevent *Gosec* from reporting it as a vulnerabilitiy every time the scan runs in our CI workflow. Now you will learn how to do this in *Gosec*.

Run the following scan:

`gosec -r false_positive/`{{exec}}

As you can see, *Gosec* prints out one potential vulnerability related to the usage of an insecure random number generator on line 9. This is only insecure when used for security purposes, but in this case we are just using it to print out a random number. So, we would like to mark this vulnerability as being a false positive. To do this we add a comment on the line where the vulnerability is reported, with the format `#nosec <Rules>` like follows:

```
8 func main() {
9 	 randomNumber := rand.Intn(100) // #nosec G404
10 	 fmt.Println("Random number:", randomNumber)
11 }
```

You can add this comment to the file by running the following Linux command:

`sed -i '9s|$| //#nosec G404|' false_positive/main.go`{{exec}}

Now if you run the same scan again you can see that *Gosec* ignores the vulnerability and reports the program as having zero vulnerabilities.

`gosec -r false_positive/`{{exec}}

> Press *NEXT* when you are ready to continue

