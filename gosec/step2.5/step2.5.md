# Fixing vulnerabilities

Run the following command:

`gosec -r broken_hash/`{{exec}}

we can see in the output that there are two listed security vulnerabilities related to the usage of the broken hash algorithm [MD5](https://en.wikipedia.org/wiki/MD5). 

So what should we do when *Gosec* detects a vulnerability? The first step would of course be to look at the output and see what information it gives us. In this case, *Gosec* points out that the problems exist in the file *main.go* on line 4 and 12 which is where the broken hash algorithm MD5 is imported and then used. *Gosec* also points out which CWE:s that the problems map to which can be helpful if you need to research what the security issue is about. 

In this example we can solve the problems by just using a more secure hash algorithm such as [SHA256](https://en.wikipedia.org/wiki/SHA-2). We can do this by replacing the occurences of the string "md5" with the string "sha256" in the file *main.go* with the following Linux command:

`sed -i 's/md5/sha256/g' main.go`{{exec}}

If we run the *Gosec* scan again we should see that the security issue is now gone!

`gosec -r broken/hash`{{exec}}

Hopefully it is clear that having a tool automatically detect vulnerabilities and pointing out exactly where they exist in the code can make it much more easy and faster to resolve the issues.

> Press *NEXT* to continue.