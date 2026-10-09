## Usage

Sources:

- [README.md](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/README.md)
- [tensor_dancer/sql/tensor_dancer--0.1.0.sql](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/sql/tensor_dancer--0.1.0.sql)
- [tensor_dancer/copilota.cpp](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/copilota.cpp)
- [tensor_dancer/meson.build](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/meson.build)
- [tensor_dancer/tensor_dancer.control](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/tensor_dancer.control)

`tensor_dancer` is an early Tensor Dancer prototype. The pinned SQL exposes `copilota(content text)` returning text; its implementation sends the prompt to a local Ollama endpoint.

### Core Workflow

```sql
CREATE EXTENSION tensor_dancer;
SELECT oid::regprocedure FROM pg_proc WHERE proname = 'copilota';
```

### Operational Boundaries

The endpoint is hard-coded to `http://localhost:11434`, as seen from the database server. The companion `vector_extra` extension has a separate control identity and is not automatically installed. The build needs BLAS, LAPACK, GGML and HTTP/JSON dependencies. No PostgreSQL-major support range is verified. The SQL declares the function `IMMUTABLE` and `PARALLEL SAFE` despite its external request; those declarations do not make network results immutable or provide transactional rollback. Treat this as experimental, protect prompt contents, and avoid using it in index expressions or assuming repeated calls are side-effect free.
