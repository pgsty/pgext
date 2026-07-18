## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/alipay/PASE/blob/eec38006c17af8b00ab65d03132bc2ae1264e947/README.md)
- [Version 0.0.1 SQL objects](https://github.com/alipay/PASE/blob/eec38006c17af8b00ab65d03132bc2ae1264e947/pase--0.0.1.sql)
- [HNSW implementation](https://github.com/alipay/PASE/tree/eec38006c17af8b00ab65d03132bc2ae1264e947/hnsw)

`pase` is an early approximate-nearest-neighbor extension for ultra-high-dimensional vectors. Version 0.0.1 defines a variable-length `pase` query type, distance operators, and `pase_hnsw` and `pase_ivfflat` index access methods for `float4[]` and textual vectors.

```sql
CREATE EXTENSION pase;
CREATE TABLE item (id text PRIMARY KEY, vector float4[] NOT NULL);
CREATE INDEX item_hnsw ON item USING pase_hnsw (vector);
SELECT id, vector <?> pase(ARRAY[0.1, 0.2]::float4[]) AS distance
FROM item
ORDER BY distance
LIMIT 10;
```

Approximate indexes trade recall for latency; measure recall against exact search across realistic dimensions, data distributions, filters, inserts, deletes, vacuum, rebuild, and concurrency. Verify vector dimension and L2 versus inner-product mode consistently.

Do not expose the bundled `ivfflat_search` and `hnsw_search` helpers to untrusted users. The reviewed SQL interpolates query vectors, predicates, limits, and table names into dynamic SQL without safe identifier/value binding, creating SQL-injection risk. It also changes planner settings during installation and contains apparent unfinished overloads. Revoke helper execution, call parameterized operator queries directly, and treat this 0.0.1 code as research-only until fully audited and tested.
