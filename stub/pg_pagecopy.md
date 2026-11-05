## Usage

Sources:

- [Upstream source 1](https://github.com/pgtoolz/pg_pagecopy/blob/203a9361d598a7433c086c621f965a2d909e5fb4/README.md)
- [Upstream source 2](https://github.com/pgtoolz/pg_pagecopy/blob/203a9361d598a7433c086c621f965a2d909e5fb4/pg_pagecopy.control)

`pg_pagecopy` 0.0.1 transfers table pages between compatible PostgreSQL instances or clones a table locally.

### Core Workflow

`pagecopy.pull` connects to a source database and transfers a specified relation. `pagecopy.clone` stays within one database; `pagecopy.check` reports eligibility. Missing destinations can be created from the source schema, while existing destinations are overwritten according to the transfer contract.

```sql
CREATE EXTENSION pg_pagecopy;
CREATE TABLE pagecopy_source (id integer, payload text);
INSERT INTO pagecopy_source VALUES (1, 'example');
SELECT * FROM pagecopy.check('public.pagecopy_source');
SELECT pagecopy.clone('public.pagecopy_source', 'public.pagecopy_copy');
VACUUM (ANALYZE) public.pagecopy_copy;
```

### Operational Boundaries

Both ends need the extension and superuser authority. Upstream supports PostgreSQL 13–18 for same-major and upward transfers, with compatibility checks; downgrades are rejected. Source freezing and an ACCESS EXCLUSIVE lock can block application work. Only supported relation/index shapes qualify. Protect connection secrets, retain a recoverable copy of the destination and run post-transfer vacuum/analyze.
