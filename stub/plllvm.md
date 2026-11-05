## Usage

Sources:

- [Upstream source 1](https://github.com/matejformanek/plllvm/blob/9e439273ae695e74a13aba391636cf4df1ac79d3/README.md)
- [Upstream source 2](https://github.com/matejformanek/plllvm/blob/9e439273ae695e74a13aba391636cf4df1ac79d3/plllvm.control)

`plllvm` 0.3 runs untrusted LLVM IR as PostgreSQL functions and includes an experimental PL/pgSQL JIT language.

### Core Workflow

Functions compile separately in each backend. The entry point returns a PostgreSQL Datum, and non-strict functions must handle nulls explicitly. `plpgsql_jit` accepts supported PL/pgSQL bodies; `plllvm.opt_level` and `plllvm.auto_checks` govern compilation and inserted safety checks. The LLVM IR language has no DO handler.

```sql
CREATE EXTENSION plllvm;
CREATE FUNCTION llvm_add1(x int4) RETURNS int4
LANGUAGE plllvm STRICT AS $$
declare i32 @plllvm_arg_i32(ptr, i32)
define i64 @f(ptr %fcinfo) {
  %x = call i32 @plllvm_arg_i32(ptr %fcinfo, i32 0)
  %r = add i32 %x, 1
  %d = sext i32 %r to i64
  ret i64 %d
}
$$;
SELECT llvm_add1(41);
```

### Operational Boundaries

Only the documented PostgreSQL 20devel/LLVM 22.1.8/macOS arm64 baseline is tested. Match the server's LLVM shared library. Installation and language functions require superuser authority; native IR has no memory sandbox and can crash or corrupt the backend. Optional expression hooks need explicit loading. No in-place extension upgrade is supplied; preserve dependent functions before following any replacement procedure.
