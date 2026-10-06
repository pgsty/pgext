## Usage

Sources:

- [DBMS_Course_Project/JsonDiffAuditLog/README.md](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/README.md)
- [DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff--1.0.sql](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff--1.0.sql)
- [DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff.control](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff.control)

`jsondiff` is the small C component of a course-project audit application. It exposes one shallow JSONB difference function; the audit tables, triggers and snapshot reconstruction belong to separate application scripts.

### Core Workflow

```sql
CREATE EXTENSION jsondiff;
SELECT public.compute_json_diff('{"a":1,"b":2}'::jsonb, '{"a":3,"c":4}'::jsonb);
```

### Operational Boundaries

`public.compute_json_diff(before_json, after_json)` returns an object containing `set` for added/changed keys and `unset` for removed keys. Nested values are replaced as whole values rather than recursively diffed. Use object-shaped input and validate edge cases before applying diffs to application records.

The versioned SQL fixes the function in public despite a relocatable control flag; do not rely on schema relocation to move this function. Installation requires a superuser and the compiled C library. No preload is specified. Upstream’s build example uses PostgreSQL 16; no broader support matrix or license is declared.
