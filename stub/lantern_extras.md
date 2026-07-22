## Usage

Sources:

- [Official README](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/README.md)
- [Official extension control file](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/lantern_extras.control)
- [Official Rust package manifest](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/Cargo.toml)

`lantern_extras` exposes server-side helpers for loading vector benchmark data, running local ONNX embedding models, calling remote embedding/completion APIs, and scheduling table-wide AI jobs. The helpers are standalone from the main Lantern extension, but they execute file, network, model-inference, and background work inside the PostgreSQL service boundary.

### Core Workflow

Create the extension as a superuser, then discover the runtimes and models built into the installed version:

```sql
CREATE EXTENSION lantern_extras;

SELECT get_available_runtimes();
SELECT get_available_models();
```

Use `ort` for a supported local ONNX model, or select a remote runtime explicitly:

```sql
SELECT llm_embedding(
  model   => 'BAAI/bge-base-en',
  input   => 'PostgreSQL vector search',
  runtime => 'ort'
);

SELECT llm_embedding(
  model     => 'text-embedding-3-small',
  input     => 'PostgreSQL vector search',
  runtime   => 'openai',
  api_token => '<api-token>'
);
```

The extension can also parse server-local `.fvecs` files and stream benchmark datasets:

```sql
SELECT *
FROM parse_fvecs('/srv/vector/siftsmall_base.fvecs', 10000);
```

### Jobs and Daemon

`add_embedding_job` and `add_completion_job` populate destination columns from a source table. Use `get_embedding_job_status`, `get_embedding_jobs`, and `get_completion_jobs` to monitor them; `cancel_embedding_job` and `resume_embedding_job` control embedding work.

The SQL-managed Lantern daemon is explicitly experimental. Its documented setup requires a preload entry, `lantern_extras.enable_daemon`, a server restart for the preload change, and optionally `lantern_extras.daemon_databases`. Verify the module filename against the installed package because the reviewed README spells the preload library `lantern_extra.so` while the extension is named `lantern_extras`.

### Security and Capacity Boundaries

- Calls can download large archives or models and perform CPU-intensive inference in a database backend. Bound concurrency, statement time, disk use, and memory, and isolate experiments from latency-sensitive workloads.
- Local paths are paths on the database server and are read with the PostgreSQL operating-system account's access. Remote URLs and API calls cause server-side network traffic; restrict SQL privileges and outbound connectivity accordingly.
- Remote runtimes send input data to the selected provider. Classify data before use and review provider retention, region, cost, rate-limit, and failure behavior.
- `lantern_extras.llm_token` can supply an API credential. Treat the setting as a secret, understand who can inspect session/config values, and prefer a narrowly scoped credential rather than embedding it in stored SQL or logs.
- Model name, runtime, preprocessing, output dimensions, and model revision form an application contract. Store that metadata and validate vector dimensions before mixing or indexing results.
- The control file marks the extension untrusted and superuser-only. Grant execution on file/network/job functions only to roles that should hold those capabilities.
