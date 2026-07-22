## Usage

Sources:

- [Official README](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/README.md)
- [Official encoding implementation](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/encode.rs)
- [Official SQL function implementation](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/datatype/function.rs)
- [Official extension entrypoint](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/lib.rs)
- [Official extension control file](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/pg_splade.control)

`pg_splade` is a work-in-progress neural sparse-embedding extension. It runs SPLADE models inside PostgreSQL and returns pgvector `sparsevec` values for document and query text, but model files, backend memory, and server preload become part of database operations.

### Setup and Core Workflow

The extension requires pgvector's `vector` extension and must be loaded by the postmaster:

```conf
shared_preload_libraries = 'pg_splade'
splade.preload_models = 'distill'
```

Restart PostgreSQL before installing and using the extension:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_splade;

SELECT list_model();
SELECT encode_document('distill', 'PostgreSQL full text document');
SELECT encode_query('distill', 'database search');
```

`encode_document(text, text)` and `encode_query(text, text)` produce sparse vectors. `truncate_sparsevec(sparsevec, integer)` keeps a smaller representation. Models load per backend and can also load lazily.

### Model Management and Caveats

`download_model(text, text)` downloads a Hugging Face repository into the server's shared data directory under `splade`, `delete_model(text)` recursively removes a named model directory, and `list_model()` reports installed models. These functions perform outbound network and filesystem operations; grant them only to trusted administrators and maintain an independent model provenance and backup policy.

The reviewed code targets little-endian 64-bit builds and PostgreSQL 14 through 17. Capacity and latency depend on the selected model and CPU, CUDA, or Metal execution path. Although encoding functions are declared immutable and parallel-safe, replacing or deleting model files can change or break results, so pin model content and coordinate updates with sessions and indexes that store embeddings.
