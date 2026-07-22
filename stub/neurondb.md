## Usage

Sources:

- [Official README](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/README.md)
- [Official control file](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/neurondb.control)
- [Official documentation](https://www.neurondb.ai/docs)

`neurondb` is a proprietary PostgreSQL extension for vector types and search, in-database ML and embedding inference, hybrid retrieval, and optional background workers. The reviewed source identifies itself as development version `4.0.0-devel`; verify licensing and the exact binary's implemented SQL surface before storing production data in its types or indexes.

### Core Workflow

Start by verifying the extension build rather than assuming every advertised module is present:

```sql
CREATE EXTENSION neurondb;

SELECT neurondb.version();

SELECT n.nspname AS schema_name,
       p.proname AS function_name,
       pg_get_function_identity_arguments(p.oid) AS arguments
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE n.nspname = 'neurondb'
ORDER BY p.proname, arguments;
```

The upstream project documents vector types including `vector`, `halfvec`, `binaryvec`, and `sparsevec`, distance operators including `<->`, `<=>`, and `<#>`, plus HNSW and IVF indexes. Follow the version-matched indexing documentation and test exact recall before adopting an operator class or tuning setting.

### Major Surfaces

- Vector storage, distance calculation, approximate HNSW/IVF search, and quantization.
- SQL-facing model training, prediction, embeddings, hybrid retrieval, reranking, and RAG helpers.
- Optional workers named `neuranq`, `neuranmon`, `neurandefrag`, and `neuranllm`.
- Optional platform-specific GPU paths for CUDA, ROCm, and Metal.

The breadth is build- and license-dependent. Use the installed catalogs and vendor documentation for the licensed artifact as the authority, not function counts or performance claims in the repository overview.

### Prerequisites and Operations

The control file requires superuser installation, marks the extension untrusted, and says to put `neurondb` in `shared_preload_libraries` and restart when using its background workers. The README declares PostgreSQL 16 through 18 support and a proprietary license that restricts commercial use, modification, and redistribution; obtain terms applicable to the deployed binary.

Vector indexes are approximate and add build memory, disk, WAL, and write-maintenance costs. Model and GPU features add native libraries, model files, backend memory, and potentially external-provider credentials or traffic. Benchmark recall and failure behavior, cap resource use, restrict execution privileges, and prove upgrade plus dump/restore on a copy of production-shaped data.
