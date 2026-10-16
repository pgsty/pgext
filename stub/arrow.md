## Usage

Sources:

- [README.md](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/README.md)
- [arrow--0.1.sql](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrow--0.1.sql)
- [sql/basic.sql](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/sql/basic.sql)
- [arrowam_handler.c](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrowam_handler.c)
- [LICENSE](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/LICENSE)
- [arrow.control](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrow.control)

`arrow` 0.1 registers a table access method backed by in-memory Arrow column buffers. The upstream regression example inserts and scans integer and floating-point columns. It is an experimental storage implementation.

### Core Workflow

```sql
CREATE EXTENSION arrow;
CREATE TABLE arrow_example (a integer, b double precision) USING arrow;
INSERT INTO arrow_example VALUES (1, 2.5);
SELECT * FROM arrow_example;
```

### Operational Boundaries

The SQL objects are `arrowam_handler(internal)` and the `arrow` table access method. Installation requires superuser privileges under the control defaults. The implementation has no completed MVCC/snapshot handling and several index-fetch and tuple operations return empty results. Use disposable experiments; do not treat it as a durable, transactional heap replacement. The upstream does not declare a supported PostgreSQL-major range.
