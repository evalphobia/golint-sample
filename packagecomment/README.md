# errors on lintPackageComment()

## golint results

```
// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L379
detached.go:2:1: package comment is detached; there should be no blank lines between it and the package statement

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L395
main.go:3:1: package comment should be of the form "Package packagecomment ..."

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L390
main.go:3:1: package comment should not have leading space
```
