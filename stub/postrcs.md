## Usage

Sources:

- [Official source](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/USAGE.md)
- [Official source](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/postrcs.control)
- [Official source](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/postrcs--1.2.1.sql)
- [Official source](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/README.md)

postrcs stores an RCS text archive in an rcs_archive domain. It does not automatically track changes to other columns: the caller must save each returned archive. The reviewed source demonstrates PostgreSQL 18 on Windows; broader compatibility and licensing are not established.

### Save and retrieve revisions

```sql
CREATE EXTENSION postrcs;
CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  history rcs_archive NOT NULL DEFAULT ''::rcs_archive
);
INSERT INTO documents(history)
VALUES (rcs_checkin(''::rcs_archive, E'First line\n', 'initial'));
UPDATE documents
SET history = rcs_checkin(history, E'First line\nSecond line\n', 'append')
WHERE id = 1;
SELECT rcs_head(history), rcs_checkout(history, '1.1'),
       rcs_diff(history, '1.1', '1.2')
FROM documents WHERE id = 1;
```

rcs_revisions lists metadata; rcs_create_branch, rcs_branch_checkin and rcs_tag return replacement archives. rcs_merge returns content and has_conflicts without persisting a merge. Inspect conflict markers and check in the resolved text. rcs_import validates archive text; rcs_export returns it.

### Limits and upgrade

Nonempty content must end in a newline. Binary content and nested branches are unsupported. Limits are 16 MiB per archive, 8 MiB per content, 1 MiB per log message and four million line pairs per LCS diff/merge. No preload is needed. Install matching 1.2.1 files and run ALTER EXTENSION postrcs UPDATE for existing 1.0, 1.1 or 1.2 installations.
