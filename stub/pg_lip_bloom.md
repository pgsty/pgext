## Usage

Sources:

- [README.md](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/README.md)
- [pg_lip_bloom/pg_lip_bloom--1.0.sql](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_bloom--1.0.sql)
- [pg_lip_bloom/pg_lip_validation.sql](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_validation.sql)
- [LICENSE](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/LICENSE)
- [pg_lip_bloom/pg_lip_bloom.control](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_bloom.control)

`pg_lip_bloom` 1.0 supplies Bloom-filter functions for rewritten join queries in the upstream adaptive-query-processing experiment. It does not automatically optimize arbitrary SQL or install a Bloom index access method despite the control-file description.

### Core Workflow

```sql
CREATE EXTENSION pg_lip_bloom;
SELECT pg_lip_bloom_init(1);
SELECT pg_lip_bloom_add(0, 42);
SELECT pg_lip_bloom_probe(0, 42);
SELECT pg_lip_bloom_info();
```

### Operational Boundaries

`pg_lip_bloom_init(integer)` allocates filter slots; `pg_lip_bloom_add(integer, integer)` populates a numbered filter and `pg_lip_bloom_probe(integer, integer)` tests membership. `pg_lip_bloom_info()` reports filter state and `pg_lip_bloom_set_dynamic(integer)` selects the experimental dynamic mode. Preserve the original joins and predicates after this probabilistic prefilter. The provided environment is PostgreSQL 12.5 and the experiments use rewritten queries and session state. Installation defaults to superuser-only; this is research code, not a transparent production optimizer.
