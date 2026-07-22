## Usage

Sources:

- [Upstream README](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/README.md)
- [Extension control file](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm.control)
- [Version 0.1.0 installation SQL](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm--0.1.0.sql)

`wasm` version `0.1.0` embeds the Wasmer runtime so PostgreSQL can instantiate a server-side WebAssembly file and expose its exported functions as generated SQL functions. Upstream calls this release experimental and documents support for PostgreSQL 10 only.

### Core Workflow

The project installs a PL/pgSQL extension separately from its shared library. Initialize it with the absolute path visible to the PostgreSQL server, then instantiate a module under a SQL-function namespace.

```sql
CREATE EXTENSION wasm;
SELECT wasm_init('/absolute/path/to/libwasm.so');
SELECT wasm_new_instance('/absolute/path/to/simple.wasm', 'ns');
SELECT ns_sum(1, 2);
```

`wasm_new_instance` returns an instance ID and creates one generated function per supported export, prefixed by the namespace. The prototype supports up to ten parameters. Integer inputs and outputs map from WebAssembly `i32` and `i64`; floating-point invocation is incomplete even though introspection maps float signatures to numeric types.

### Introspection and Safety

`wasm_init` creates internal FDWs and the `wasm.instances` and `wasm.exported_functions` foreign tables for inspecting process-local instances and exports. Module paths are read by the backend process, and loaded instances are held in backend memory rather than durable database storage.

Initialization drops and recreates internal wrappers, servers, and the `wasm` schema with `CASCADE`; rerunning it can destroy objects in that schema. It also dynamically creates C-language functions from the supplied library path. Restrict execution to trusted administrators and review module provenance. There is no documented preload requirement, but PostgreSQL 11 and later are explicitly outside the reviewed compatibility claim; test crashes, resource exhaustion, ABI compatibility, and backend restart behavior only in an isolated environment.
