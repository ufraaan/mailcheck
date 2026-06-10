# email-check

a lil tool to check domain email config (MX, SPF, DMARC). nothing fancy.

pipe domains in, get CSV out:

```
echo "gmail.com" | go run main.go
```

or pass as CLI args:

```
email-check gmail.com protonmail.com
```

sample output:

```
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
gmail.com, true, true, v=spf1 redirect=_spf.google.com, true, v=DMARC1; p=none; sp=quarantine; dis=...;
```

just the stdlib.
