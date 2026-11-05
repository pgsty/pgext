## Usage

Sources:

- [Upstream source 1](https://github.com/postgamma/gammabranch/blob/3aeda5c228c9ab72373a1774b981f58cceb871b2/README.md)
- [Upstream source 2](https://github.com/postgamma/gammabranch/blob/3aeda5c228c9ab72373a1774b981f58cceb871b2/overlay_branch.control)

`overlay_branch` 0.1, distributed as Gamma Branch, stores experimental draft changes in MVCC delta overlays.

### Core Workflow

Restart after adding the preload, preserving any other preloaded libraries. `create_branch` creates a draft, `use_branch` selects it for the session, and `apply_branch` merges after conflict checks. The example returns to MAIN without applying the draft.

```ini
shared_preload_libraries = 'overlay_branch'
```

```sql
CREATE EXTENSION overlay_branch;
CREATE TABLE branch_products (id integer PRIMARY KEY, price integer);
INSERT INTO branch_products VALUES (1, 10);
SELECT create_branch('draft');
SELECT use_branch('draft');
UPDATE branch_products SET price = 20 WHERE id = 1;
SELECT use_branch(NULL);
SELECT * FROM branch_products;
```

### Operational Boundaries

PostgreSQL 17 only; disposable environments only. Without preload, write redirection silently fails and writes reach MAIN. The default supported shape is a plain heap table with one primary-key column. Partitioned/foreign tables, foreign keys and other unsupported constraints are rejected. Concurrent MAIN changes can make apply fail; snapshot mode is time-bounded and can force a session back to MAIN.
