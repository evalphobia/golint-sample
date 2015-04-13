# errors on lintVarDecls()

## golint results

```
// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L936
main.go:4:17: should drop = 0 from declaration of var zero; it is the zero value

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L969
main.go:5:10: should omit type int from declaration of var one; it will be inferred from the right-hand side
```
