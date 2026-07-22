## Usage

Sources:

- [Official README at the catalog revision](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/README.md)
- [Extension SQL at the catalog revision](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/jsonb_extend--1.0.sql)

`jsonb_extend` 1.0 provides shallow and recursive JSONB merging. It is useful when merge order must be explicit: later object values replace earlier values, while arrays are concatenated. The source targets old PostgreSQL internals, so build and test it against the exact server major you operate.

### Core Workflow

```sql
CREATE EXTENSION jsonb_extend;

-- Shallow object merge: the rightmost value wins.
SELECT jsonb_extend('{"a":1,"nested":{"x":1}}'::jsonb,
                    '{"a":2,"nested":{"y":2}}'::jsonb);

-- Recursive object merge when the first flag is false.
SELECT jsonb_deep_extend(false,
                         '{"nested":{"x":1}}'::jsonb,
                         '{"nested":{"y":2}}'::jsonb);

-- Arrays are concatenated; a later scalar becomes another element.
SELECT jsonb_extend('[1,2]'::jsonb, '3'::jsonb);
```

### Function Index

- `jsonb_extend(variadic jsonb[])` performs a shallow merge. An object must be followed only by objects; an initial array accepts later arrays, objects, or scalars.
- `jsonb_deep_extend(boolean, variadic jsonb[])` accepts objects and recursively merges nested objects when the flag is false. A true flag treats nested values atomically instead of descending into them.
- Both functions are declared immutable and strict. A single input is returned unchanged; incompatible shapes raise an error.

### Caveats

- Merge order is significant, and replacement is not the same as RFC JSON Merge Patch semantics.
- The implementation was written against PostgreSQL 9.4-era JSONB internals. A successful source build on one major is not compatibility evidence for another.
- Review output sizes before merging large arrays or deeply nested objects; the operation constructs a new JSONB value.
