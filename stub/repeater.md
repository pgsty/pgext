## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/README.md)
- [PostgreSQL 12 development patch](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/pg12_devel.patch)
- [Repeater implementation](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/repeater.c)

`repeater` is a demonstration module for sending a serialized raw PostgreSQL query plan to remote nodes and executing the identical plan there. It belongs to the `execplan` research project, requires `postgres_fdw`, and depends on invasive PostgreSQL 12 core and libpq patches.

```conf
shared_preload_libraries = 'repeater'
```

After loading the library, the patched server and client machinery—not standard PostgreSQL SQL APIs—manage portable object identifiers, plan serialization, transfer, remote sessions, and transactions. Do not load it into an unpatched server or mix binaries built from different source trees.

Serialized plans depend on PostgreSQL internal node layouts, catalog objects, types, operators, schemas, statistics assumptions, and binary version. Executing a supplied raw plan bypasses normal planning safety and can crash or corrupt semantics under version or catalog skew. This is PostgreSQL 12-era research code with simplistic demo tests, not a production distributed executor. Any study must use isolated disposable clusters and validate authentication, plan provenance, object mapping, failures, cancellation, remote transactions, and upgrades.
