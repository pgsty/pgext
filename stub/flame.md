## Usage

Sources:

- [Upstream README and proof-of-concept status](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/README.md)
- [Extension control file](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/flame.control)
- [Cargo package metadata](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/Cargo.toml)
- [Model initialization](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/src/fast_embeddings.rs)

`flame` version `0.0.1` is a proof-of-concept Rust/pgrx text-embedding function. `fast_embed(text)` uses fastembed's `AllMiniLML6V2` model and returns the embedding as a PostgreSQL real array.

### Example

```sql
CREATE EXTENSION flame;
SELECT array_length(fast_embed('hello world'), 1);
SELECT fast_embed('PostgreSQL text embedding');
```

The model is initialized once per backend process and may need to download/cache model artifacts under the PostgreSQL OS account. This can add cold-start latency, network supply-chain exposure, memory use per connection, and nondeterministic operational failure. The README calls the release a proof of concept, lists missing tests, and disagrees with the selected model about expected dimensions; initialization failure can panic and embedding failure returns an empty array. Pin and pre-stage model files, restrict execution, cap input/concurrency, and validate output length/version before storing vectors.
