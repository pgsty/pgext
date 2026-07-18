## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/README.md)
- [Rust implementation](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/src/lib.rs)
- [Install-time SQL script](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/sql/symbiotic_python.sql)
- [Extension control file](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/symbiotic_python.control)

`symbiotic_python` is an abandoned one-day pgrx prototype, not an extension that is safe to install. Its install script immediately invokes `create_venv`, creates files under the PostgreSQL service account's home directory, runs `pip` to download `asyncpg` and `asyncpg-listen`, writes launcher scripts, and starts a detached Python listener connected to the current database.

### Critical Security Warning

Public SQL functions interpolate the caller-supplied environment name into `sh -c` commands used by `create_venv` and `drop_venv` without shell quoting. A caller able to invoke them can inject shell syntax and execute operating-system commands as the PostgreSQL service account; `drop_venv` also constructs an unquoted `rm -rf` command.

Do not run `CREATE EXTENSION symbiotic_python`, do not expose its shared library, and do not test it on a host containing data or credentials. Source inspection in an isolated build environment is the only appropriate use. The repository contains no operational documentation, maintenance history, or mitigation for these install-time side effects.
