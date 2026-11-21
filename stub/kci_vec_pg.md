## Usage

Sources:

- [pg/README.md](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/README.md)
- [pg/kci_vec_pg.control](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/kci_vec_pg.control)
- [pg/kci_vec_pg--0.3.sql](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/kci_vec_pg--0.3.sql)
- [LICENSE](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/LICENSE)

`kci_vec_pg` exposes KCI-Vec approximate nearest-neighbor search through SQL functions. SQL version 0.3 is a prototype tested on PostgreSQL 15–17 on Linux x86-64. Its source uses the PolyForm Noncommercial 1.0.0 license. Vector index images are external files, so backups and standby operation require explicit handling.

### Build and Search

Create the extension as a superuser, then build from a table whose ID casts to bigint and whose vector casts to real arrays:

```sql
CREATE EXTENSION kci_vec_pg;
CREATE TABLE docs (id bigint PRIMARY KEY, embedding real[]);
INSERT INTO docs
SELECT i, ARRAY(SELECT random()::real FROM generate_series(1, 128))
FROM generate_series(1, 1000) AS i;
SELECT kci_vec_build('docs_index', 'docs', 'id', 'embedding', nlist => 16);
SELECT kci_vec_track('docs_index');
SELECT * FROM kci_vec_search(
  'docs_index', (SELECT embedding FROM docs WHERE id = 1), 10);
SELECT * FROM kci_vec_status();
```

The build supports cosine, squared Euclidean distance and inner product. Search returns IDs and scores; cosine and inner-product scores increase with similarity, while squared distances decrease. Larger probe and rerank settings trade work for recall. An overload accepts an allowed-ID array for filtering. The planner does not automatically choose this search through a PostgreSQL index access method.

### Changes and Maintenance

`kci_vec_track()` records inserts, updates, deletes and truncation in ordinary tables; searches merge pending changes, including the caller's transaction. `kci_vec_refresh()` folds them into the image. Large backlogs slow searches. `kci_vec_untrack()` stops tracking; `kci_vec_drop()` removes an index.

Optional preload enables cheaper change tracking and the background worker:

```ini
shared_preload_libraries = 'kci_vec_pg'
kci_vec_pg.databases = 'mydb'
kci_vec_pg.refresh_pending = 10000
```

Changing preload requires a restart; retain existing library entries. The ordinary function interface does not require this setup. The separate `vector` extension is optional; install it before this extension when its vector-type overloads are wanted.

### Recovery and Limits

Images under `$PGDATA/kci_vec/` are outside WAL and `pg_dump`, and are not encrypted by the extension. A restored database or standby must rebuild its local image with `kci_vec_materialize()`; searches reject a missing or outdated image. `kci_vec_image_here()` checks local presence. Builds and refreshes run only on the primary.

Call `kci_vec_drop()` before removing the extension because dropping it does not remove images. The documented upgrade procedure installs new files, restarts PostgreSQL and runs `ALTER EXTENSION kci_vec_pg UPDATE`. NULL IDs or vectors are skipped. Very selective filters may need more probes or an exact SQL search. Interfaces and catalog tables may change while this remains a prototype.
