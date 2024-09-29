# Advanced *Gosec* features

In this part of the tutorial you will learn about some other useful functionalities of *Gosec*.

### Severity and Confidence

By default, Gosec reports all potential vulnerabilities it detects, regardless of their severity level or the certainty of whether they are true positives (i.e., actual vulnerabilities). Depending on the security requirements of the project, you may want to filter out vulnerabilities based on a minimum severity and confidence level. This approach means that the found vulnerabilities are more likely to be actual security issues and since these have a higher severity it is most likely these issues that you would want to focus on first. 

In *Gosec* there are three severity levels and three confidence levels: low, medium and high. The following command tells *Gosec* to find all security vulerabilities of a minimum severity level of medium and a minimum confidence level of medium:

`gosec -severity medium -confidence medium ./...`{{exec}}

You can see several issues were found, but if we instead change the minimum severity level to high we find only one issue that is of high severity:

`gosec -severity high -confidence medium ./...`{{exec}}

### Handling false positives

With all SAST tools there is the issue of false positives. A false positive is a reported vulnerability that does not actually pose a security risk.

TODO: show example of running gosec that finds a false positive.

TODO: insert nosec things in the code.

TODO: run gosec again to show that it no longer reports the false positive.

> Press *NEXT* when you are ready to continue

