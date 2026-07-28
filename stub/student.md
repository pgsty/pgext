## Usage

Sources:

- [Official upstream README](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/README.md)
- [Official extension control file (student.control)](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/student.control)
- [Official extension SQL (student--1.0.sql)](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/student--1.0.sql)

`student` — In order to develop your base type in postgres, write:. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION student;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `student(text)` is an extension function and returns `student`.
- `student_in(cstring)` is an extension function and returns `student`.
- `student_out(student)` is an extension function and returns `cstring`.
- `text(student)` is an extension function and returns `text`.
- `student` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
