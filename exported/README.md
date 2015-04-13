# errors on lintExported()

## golint results

```
// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L796
main.go:3:1: exported function Main should have comment or be unexported

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L802
main.go:6:1: comment on exported function ExportedMain should be of the form "ExportedMain ..."

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L869
main.go:7:6: func name will be used as exported.ExportedMain by other packages, and that stutters; consider calling this Main

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L837
main.go:12:2: exported var Foo1 should have comment or be unexported

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L819
main.go:13:2: exported var Foo3 should have its own declaration

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L837
main.go:17:2: exported const Bar should have comment (or a comment on this block) or be unexported

// https://github.com/golang/lint/blob/3b3fdebcde972819016da7b8377e79651998f5fc/lint.go#L741
main.go:20:6: exported type Baz should have comment or be unexported
```
