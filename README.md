### Update

A patch for this issue has been merged https://github.com/ent/ent/pull/3344

Thanks @a8m!

### Introduction

Example made to demostrate that the issue https://github.com/ent/ent/issues/3208 is still present in the latest version (`v0.11.8`) of ent

The issue is that the generation does not seem to be deterministic when a named import is created

To reproduce

```go
go generate ./...
```

```sh
git status --porcelain
 M ent/user.go
```

```diff
iff --git a/ent/user.go b/ent/user.go
index aedb17f..2345b5a 100644
--- a/ent/user.go
+++ b/ent/user.go
@@ -7,9 +7,8 @@ import (
        "strings"

        "entgo.io/ent/dialect/sql"
-       "github.com/migmartri/ent-gen-troubleshoot-example/ent/user"
-
        v1 "github.com/migmartri/ent-gen-troubleshoot-example/api/v1"
+       "github.com/migmartri/ent-gen-troubleshoot-example/ent/user"
 )
```

If you run it again another file is now affected instead

```
diff --git a/ent/user_create.go b/ent/user_create.go
index 0947122..7e4710c 100644
--- a/ent/user_create.go
+++ b/ent/user_create.go
@@ -9,9 +9,8 @@ import (

        "entgo.io/ent/dialect/sql/sqlgraph"
        "entgo.io/ent/schema/field"
-       "github.com/migmartri/ent-gen-troubleshoot-example/ent/user"
-
        v1 "github.com/migmartri/ent-gen-troubleshoot-example/api/v1"
+       "github.com/migmartri/ent-gen-troubleshoot-example/ent/user"
 )
```
