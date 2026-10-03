## Usage

Sources:

- [extensions/pg_glot/pg_glot.control](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/pg_glot.control)
- [README.md](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/README.md)
- [Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/Cargo.toml)
- [extensions/pg_glot/Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/Cargo.toml)
- [extensions/pg_glot/src/lib.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/src/lib.rs)

`pg_glot` provides Korean, Japanese and Chinese PostgreSQL text-search configurations with embedded Lindera dictionaries, plus the `glot.rrf` fusion primitive. The independently installed hybrid companion is optional.

### Core Workflow

```sql
CREATE EXTENSION pg_glot;
SELECT to_tsvector('japanese', '東京都に住む');
SELECT * FROM glot.rrf(ARRAY[10,20,30]::bigint[], ARRAY[20,40]::bigint[], 60);
```

### Operational Boundaries

Install as a superuser. The base extension has no extension dependency or preload requirement. The workspace defaults to PostgreSQL 17; Cargo feature switches alone are not evidence of a tested compatibility matrix. Korean is the most thoroughly validated language upstream. Text search uses standard PostgreSQL matching/ranking and can use application-created GIN indexes. `glot.rrf` merges ranked identifier arrays; no embedding model is supplied.
