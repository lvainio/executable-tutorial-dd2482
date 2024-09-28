# *Gosec* basics

Now that *Gosec* is installed we will begin with learning some basic commands. 

### Run *Gosec*

Often we want to scan an entire repository for vulnerabilities. *Gosec* has an easy way of doing this. Execute the following command to run *Gosec* on all Go files in the current directory and recursively in all subdirectories.

`gosec ./...`{{exec}}

After running the command, you can see that *Gosec* provides output that lists any potential security issues found. Since we are scanning recursively you can see that there are many potential issues listed. If you look at the output in the terminal you also see that *Gosec* lists in which file and where in the file the security issue is located. This is very helpful for developers that will work on fixing the issues. Additionally you can see that the output states which CWE ([Common Weakness Enumeration](https://cwe.mitre.org/)) that the found vulnerability maps to which can also be helpful in fixing the issue.

### Scan Specific Directories

Sometimes we only need to scan certain parts of our code base, such as a single directory. We can use the `-r` flag for this purpose. Run the following command to scan the directory *insecure_rand/*:

`gosec -r insecure_rand/`{{exec}}

You can see that there is less output now which can be easier to filter through if you are working on a specific part of the code base. This can also save time, since a full scan could take some time on a really large code base.

### Scan for Specific vulnerabilities

It is also possible to instruct *Gosec* to only search for specific vulnerabilities. By running the following command we can see which rules (vulnerabilities) that *Gosec* supports:

`gosec`{{exec}}

We can pick any rule from the list we see in the terminal but as an example we can run the following to only look for rule *G404: Insecure random number source (rand)*:

`gosec -include=G404, ./...`{{exec}}

As you can see in the terminal we only see one vulnerability and it matches rule G404. Feel free to experiment with other rules in the terminal.

> Press *NEXT* when you are ready to continue