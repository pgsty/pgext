## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/README.md)
- [Official extension control file (resetsequence.control)](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/resetsequence.control)
- [Official extension SQL (resetsequence--1.0.0.sql)](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/resetsequence--1.0.0.sql)

`resetsequence` — This module contains a single PostgreSQL extension, comprising simple utility functions that help with maintaining sequence values. These functions let you reset the sequences to the maximum values used in table columns that link to them via DEFAULT clauses. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION resetsequence;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `resetseq_reset_sequences_in_database()` is an extension function and returns `setof`.
- `resetseq_reset_sequences_in_schema(name)` is an extension function and returns `setof`.
- `resetseq_reset_sequences_in_table(regclass)` is an extension function and returns `setof`.
- `resetseq_sequence_max_value(oid)` is an extension function and returns `bigint`.
- `resetseq_report_type` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
