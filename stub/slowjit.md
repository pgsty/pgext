## Usage

Sources:

- [Official upstream README](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/README.md)
- [Official extension control file (slowjit.control)](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/slowjit.control)
- [Official implementation source](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/slowjit.c)

`slowjit` is an educational PostgreSQL JIT-provider implementation. It emits C code at runtime, invokes a C compiler, loads the resulting shared library, and currently handles only a small set of expression operators.

### Core Workflow

Build and install the provider library, select it through PostgreSQL's JIT-provider configuration, and force a trivial expression through JIT in an isolated server:

```ini
jit_provider = 'slowjit'
jit_above_cost = 0
```

```sql
EXPLAIN (SETTINGS ON)
SELECT 1;
```

The README uses this query only to demonstrate that one function was JIT-compiled. The module does not install a user-facing SQL API or establish a standalone `CREATE EXTENSION` workflow.

### Important Setting

- `slowjit.cc_path` selects the C compiler executable; the implementation default is `cc`.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0.0`.
- The control file marks the extension as relocatable.
- Upstream explicitly calls the provider very inefficient and documents support for only a few operators.
- Runtime compilation executes a server-side compiler and creates loadable code. Restrict filesystem permissions and never treat this prototype as a hardened provider.
- JIT provider interfaces are version-sensitive; build and test it against the exact PostgreSQL server source.
