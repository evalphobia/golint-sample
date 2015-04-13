# errors on lintUnexportedReturn()

## golint results

```
https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L1311
main.go:6:12: exported func Get returns unexported type unexportedreturn.privateStruct, which can be annoying to use
```
