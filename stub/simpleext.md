## Usage

Sources:

- [Official upstream README](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello-with-ffi/README.md)
- [Official extension control file (simpleext.control)](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello-with-ffi/simpleext.control)
- [Official extension SQL (simpleext--1.0.sql)](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello/simpleext--1.0.sql)

`simpleext` — 현재 프로젝트에서 바로 빌드 가능한 가장 간단한 C++ PostgreSQL 확장 템플릿입니다. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION simpleext;
   SELECT simpleext.hello_world(1); -- 2 (입력값 + 1)
   SELECT simpleext.add_ints(2, 3); -- 5
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `simpleext.add_ints(a int, b int)` is an extension function and returns `int`.
- `simpleext.hello_world()` is an extension function and returns `text`.
- `simpleext` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
