## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/kibae/pg_onnx/blob/ce634755381b3fb27b5724a3c58f39228d94eef8/README.md)
- [Function reference](https://github.com/kibae/pg_onnx/blob/ce634755381b3fb27b5724a3c58f39228d94eef8/docs/functions.md)
- [Extension source](https://github.com/kibae/pg_onnx/tree/ce634755381b3fb27b5724a3c58f39228d94eef8/pg_onnx)

`pg_onnx` runs ONNX inference from PostgreSQL. A dynamic background worker hosts `onnxruntime-server` and reuses one runtime session per imported model, avoiding a separate large runtime allocation in every client backend.

```sql
CREATE EXTENSION pg_onnx;
SELECT pg_onnx_import_model(
  'score_model', 'v1', $1::bytea,
  '{"cuda": false}'::jsonb, 'reviewed model'
);
SELECT pg_onnx_execute_session(
  'score_model', 'v1', '{"input": [[1.0, 2.0]]}'::jsonb
);
```

Pass model bytes as a parameter; do not use a caller-controlled server file path. Restrict import, drop, session, and execution functions. An ONNX model and optional custom-operation library are executable native inputs: accept only reviewed, hashed artifacts and constrain extension-library paths. Treat outputs as model results, not deterministic database facts.

Plan CPU, GPU, RAM, worker slots, timeouts, cancellation, model size, session recycling, and concurrency. Test worker crash/restart, transaction rollback around model metadata, failover, replicas, backup/restore, CUDA/ONNX Runtime ABI, and upgrades. Trigger-based inference adds remote computation to writes and can block or fail transactions; prefer asynchronous inference unless synchronous semantics are essential.
