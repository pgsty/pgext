## Usage

Sources:

- [1.1.1 README](https://github.com/supervc-stack/VectorChord/blob/1.1.1/README.md)
- [Control and dependency](https://github.com/supervc-stack/VectorChord/blob/1.1.1/vchord.control)
- [Preload requirement](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/lib.rs)
- [1.1.1 SQL objects](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/install/vchord--1.1.1.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/index/gucs.rs)
- [1.1.1 migration](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/upgrade/vchord--1.1.0--1.1.1.sql)
- [1.1.1 release notes](https://github.com/supervc-stack/VectorChord/releases/tag/1.1.1)

`vchord` adds approximate vector indexes to PostgreSQL using pgvector's types. It provides the partition-based `vchordrq` and graph-based `vchordg` access methods. The extension requires `vector`, shared preloading, and superuser privileges to create.

### Create and Query an Index

Add the library to the existing preload list, preserving other entries, and restart PostgreSQL:

```conf
shared_preload_libraries = 'vchord'
```

```sql
CREATE EXTENSION vchord CASCADE;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
INSERT INTO items(embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
CREATE INDEX items_embedding_idx ON items
USING vchordrq (embedding vector_l2_ops);

SELECT id FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
SELECT vchordrq_prewarm('items_embedding_idx'::regclass);
```

Use `vector_l2_ops` with `<->`, `vector_ip_ops` with `<#>`, and `vector_cosine_ops` with `<=>`. The inner-product operator returns a negative value for ascending index ordering. The same operator classes can be used with the graph access method; choose one index design for the workload:

```sql
CREATE INDEX items_embedding_graph_idx ON items
USING vchordg (embedding vector_l2_ops);
```

### Range Queries and Tuning

The extension supplies explicit sphere predicates for range search:

```sql
SELECT id FROM items
WHERE embedding <<->> sphere('[1,2,3]'::vector, 0.5);

SET vchordrq.probes = '100';
SET vchordrq.epsilon = 1.9;
SET vchordg.ef_search = 64;
```

`<<->>`, `<<#>>`, and `<<=>>` are sphere predicates for L2, inner product, and cosine metrics. Probe counts depend on the partition layout; tune them with representative data. The epsilon setting controls the reranking tradeoff. The graph search setting controls its candidate search breadth. Both index methods are approximate: check recall, filters, and query plans before choosing settings.

### Quantization in 1.1.1

`rabitq8` and `rabitq4` store quantized vectors. `quantize_to_rabitq8` and `quantize_to_rabitq4` accept `vector` or `halfvec`. Version 1.1.1 adds `dequantize_to_vector` and `dequantize_to_halfvec` overloads for both quantized types:

```sql
SELECT dequantize_to_vector(quantize_to_rabitq8('[1,2,3]'::vector));
SELECT dequantize_to_halfvec(quantize_to_rabitq4('[1,2,3]'::halfvec));
```

Quantization loses precision; dequantization returns an approximation. The release also replaces the quantization implementation. Install matching library and SQL files, restart for the preloaded library, then update each database:

```sql
ALTER EXTENSION vchord UPDATE TO '1.1.1';
```

The 1.1.0-to-1.1.1 script adds these four conversion overloads and declares no index-format migration. Earlier-version upgrade requirements depend on the starting version. Index construction and prewarming consume resources; schedule them for the dataset size. `vchordg_prewarm` is the corresponding graph-index helper. The control is relocatable, so qualify extension objects or include their installation schema in the search path when installed outside the usual schema.
