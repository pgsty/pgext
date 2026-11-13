## Usage

Sources:

- [README](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/README.md)
- [Control](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/pg_crdt.control)
- [Install SQL](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/pg_crdt--0.0.1.sql)
- [Installation and backup](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/install.md)
- [Untrusted writers](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/guides/security.md)
- [Collaborative text](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/types/collaborative-text.md)

`pg_crdt` 0.0.1 provides mergeable counters and collaborative text in SQL/PL/pgSQL. It supports PostgreSQL 16–18, requires no other extension or shared preload, and installs into the fixed `crdt` schema. The experimental text-column representation rewrites the entire document on each edit.

### Install

With its control and versioned SQL files installed on the server, a role with `CREATE` on the database can create the extension; a superuser is not required. Do not specify another schema. Installation rejects an existing schema writable by another role. The official guide also documents pg_tle registration and direct SQL installation; those are alternative installation paths, not additional dependencies.

```sql
CREATE EXTENSION pg_crdt;
```

### Counter workflow

`crdt.pn` stores each replica's positive and negative totals. `crdt.pn_increment`, `crdt.pn_decrement` and `crdt.pn_value` manipulate or read a counter; `crdt.pn_merge` combines states idempotently. Update the stored value in one SQL statement so row locking protects concurrent changes. Do not read, merge client-side and overwrite.

```sql
CREATE TABLE posts (
  id bigint PRIMARY KEY,
  likes crdt.pn NOT NULL DEFAULT crdt.pn_empty()
);
INSERT INTO posts (id) VALUES (1);
UPDATE posts SET likes = crdt.pn_increment(likes, 'alice', 3) WHERE id = 1;
UPDATE posts SET likes = crdt.pn_decrement(likes, 'alice', 1) WHERE id = 1;
SELECT crdt.pn_value(likes) FROM posts WHERE id = 1;
```

### Collaborative text

`crdt.yseq` stores a document in a JSONB-backed column; `crdt.yseq_insert`, `crdt.yseq_merge` and `crdt.yseq_value` edit, merge and read it. The `yseq_` functions require UTF8 for operations that transmit deleted characters. Column state retains tombstones and grows with edit history. The `ydoc_` API instead keeps documents in extension tables and requires table/sequence grants for application wrappers; edits there do not trigger policies or change feeds on a referencing application row.

### Trust boundaries and backup

Untrusted clients must not update raw CRDT state: they could forge another replica's totals or edits. Use restricted application wrappers around `crdt.pn_apply_delta` and the text apply APIs, bind a trusted replica identity, check row authorization and pin the wrapper search path. Client-controlled settings alone do not authenticate identity. The extension grants function execution publicly but no access to document tables. Extension configuration dumps preserve document tables and sequence positions; include the `crdt` schema. Dropping the extension deletes its document tables; using cascade also deletes application columns that use its types. Export data and remove dependencies deliberately before uninstalling.
