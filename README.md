<div align="center">
  <h2>mailcheck</h2>
  <img src="demo.gif" width="720">
</div>

a lil tool to check domain email config (MX, SPF, DMARC). nothing fancy.

pipe domains in, get CSV out:

```
echo "gmail.com" | go run main.go
```

or pass as CLI args:

```
go run main.go gmail.com protonmail.com
```

sample output:

```
domain          hasMX  hasSPF  spfRecord                               hasDMARC  dmarcRecord
gmail.com       true   true    v=spf1 redirect=_spf.google.com         true      v=DMARC1; p=none; sp=quarantine; rua=mailto:mailauth-reports@google.com
protonmail.com  true   true    v=spf1 include:_spf.protonmail.ch ~all  true      v=DMARC1; p=quarantine; fo=1; aspf=s; adkim=s;
```

just the stdlib.
