## Usage

Sources:

- [pgext/README.md](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/README.md)
- [pgext/Cargo.toml](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/Cargo.toml)
- [pgext/tqvector.control](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/tqvector.control)
- [pgext/sql/tqvector--0.1.0.sql](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/sql/tqvector--0.1.0.sql)
- [pgext/src/lib.rs](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/src/lib.rs)

`tqvector` 0.1.0 is a Rust component of TurboQuant Pro that stores scalar-quantized embeddings. It declares PostgreSQL 14–17 build features and uses no Python runtime. The control declares superuser=false, but the install script creates native C types/functions requiring elevated privileges. No shared preload is declared.

### Core workflow

```sql
CREATE EXTENSION tqvector;
SELECT tq_dim(tq_compress(ARRAY[0.1,0.2,0.3,0.4]::float4[], 3));
SELECT tq_decompress(tq_compress(ARRAY[0.1,0.2,0.3,0.4]::float4[], 3));
```

### Operations and limits

`tq_compress` accepts 2, 3 or 4 bits; `tq_decompress` reconstructs an approximate float array. Inspect values with `tq_dim`, `tq_bits`, `tq_norm`, `tq_ratio` and `tq_size_bytes`. `tq_cosine_sim` and `tq_cosine_dist` compare equal-dimensional vectors; `<=>` is cosine distance and `<->` is Euclidean distance. Quantization is lossy, so validate retrieval quality on your own data.

The installed SQL defines distance operators, not an index access method. Ordering by distance therefore does not itself create an ANN index. Treat table/column names passed to `tq_bulk_compress` as trusted administrative input. The optional GPU build feature is separate from the basic SQL workflow.
