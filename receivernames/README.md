# errors on lintReceiverNames()

## golint results

```
// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L1185
main.go:7:1: receiver name should not be an underscore

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L1189
main.go:9:1: receiver name should be a reflection of its identity; don't use generic names such as "me", "this", or "self"
main.go:10:1: receiver name should be a reflection of its identity; don't use generic names such as "me", "this", or "self"
main.go:11:1: receiver name should be a reflection of its identity; don't use generic names such as "me", "this", or "self"

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L1194
main.go:14:1: receiver name s should be consistent with previous receiver name m for myStruct
```
