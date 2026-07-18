## Usage

Sources:

- [Upstream README](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/README.md)
- [Extension control file](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm.control)
- [Version 0.1.0 installation SQL](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm--0.1.0.sql)

`wasm` embeds the Wasmer runtime so PostgreSQL can instantiate a WebAssembly module and expose its exported functions as generated SQL functions. Upstream version `0.1.0` is experimental and supports PostgreSQL 10 only; the README explicitly says PostgreSQL 11 is not supported.

After building both the shared library and PL/pgSQL extension, initialize it with the absolute server-side library path, then load a WebAssembly module:

```sql
CREATE EXTENSION wasm;
SELECT wasm_init('/absolute/path/to/libwasm.so');
SELECT wasm_new_instance('/absolute/path/to/simple.wasm', 'ns');
SELECT ns_sum(1, 2);
```

The namespace argument prefixes generated functions; the example creates `ns_sum`. The prototype maps WebAssembly `i32`, `i64`, and `v128` values to PostgreSQL types, while floating-point support is incomplete.

`wasm_init()` dynamically creates C-language bindings and internal FDWs. Its installation SQL drops and recreates the internal FDWs, servers, and the `wasm` schema with `CASCADE`, so do not rerun it in a database that contains objects under that schema without reviewing the consequences. Module and library paths must be absolute paths visible to the PostgreSQL server process. No preload setting is documented.
