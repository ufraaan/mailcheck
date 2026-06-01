# email-check

a lil tool to check domain email config (MX, SPF, DMARC). nothing fancy.

pipe domains in, get CSV out:

```
echo "gmail.com" | go run main.go
```

just the stdlib.
