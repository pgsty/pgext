## Usage

Sources:

- [Upstream README and current-status statement](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/README.md)
- [Extension control file](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/pglens.control)
- [SQL API](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/sql/pglens--0.1.0.sql)
- [Recompute implementation](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/src/recompute.c)
- [Roadmap](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/docs/ROADMAP.md)

`pglens` is an experimental projection/read-model extension. A lens ties a plain SQL view to a physical target table and a stable key. Registration copies the view's current shape into an empty table with a primary key; explicit recomputation selects one key from the view and upserts or deletes the corresponding target row.

### Register and recompute a lens

```sql
CREATE EXTENSION pglens;

CREATE TABLE accounts (
  id bigint PRIMARY KEY,
  email text NOT NULL,
  status text NOT NULL
);

CREATE VIEW account_projection_view AS
SELECT id, email, status FROM accounts;

SELECT pglens.register_lens(
  'account_projection',
  'account_projection_view'::regclass,
  'public',
  'account_projection',
  'id'
);

INSERT INTO accounts VALUES (1, 'a@example.com', 'active');
SELECT pglens.recompute_lens('account_projection', '1');

TABLE account_projection;
```

Despite the intended asynchronous architecture, version `0.1.0` does not implement WAL decoding, committed-change capture, affected-key discovery, a background worker, checkpointing, retry, or restart recovery. Recompute is manual.

Only plain views are supported; the key must identify at most one projection row; the target shape is fixed at registration; later view changes are not propagated; and negative-path coverage is limited. Upstream explicitly marks the project experimental and unsuitable for production. Treat it as a design prototype until consistency, crash recovery, privilege boundaries, schema migration, and operational observability are implemented and tested.
