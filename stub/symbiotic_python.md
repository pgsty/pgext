## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/README.md)
- [Install-time SQL script](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/sql/symbiotic_python.sql)
- [Rust implementation](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/src/lib.rs)
- [Python notification listener](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/scripts/symbiotic.py)
- [Build metadata](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/Cargo.toml)

`symbiotic_python` is an abandoned pgrx prototype that tries to create a per-database Python virtual environment and launch an asynchronous notification listener. It is unsafe to install: extension installation itself performs operating-system, filesystem, package-download, and process-launch side effects as the PostgreSQL service account.

### Intended Workflow

The SQL API includes `create_venv(name)`, `drop_venv(name)`, `venv_path(name)`, and `run_symbiotic(venv, channel)`, plus zero-argument wrappers. The install script automatically invokes environment creation for the current database. It creates files beneath the service account's `$HOME/.symbiotic`, runs `pip` for `asyncpg` and `asyncpg-listen`, writes shell launchers, and starts a detached listener.

The listener connects back to the local database, subscribes to a notification channel, appends payloads to a log, and exits on a shutdown payload. There is no reliable SQL status or stop interface, package lockfile, process supervision, or cleanup transaction for these external side effects.

### Critical Security Boundary

`create_venv(name)` and `drop_venv(name)` interpolate the caller-controlled name into shell commands without quoting. This permits shell-command injection as the PostgreSQL operating-system account; the deletion path also constructs an unquoted `rm -rf`. These functions are not made safe merely by placing them in the `symbiotic` schema, and default function privileges can expose them more broadly than intended.

Do not install or load `symbiotic_python` on any host containing data, credentials, or network access. Do not run its SQL even for a trial. The only appropriate evaluation is offline source inspection or a throwaway sandbox with no secrets and no reachable systems. Its old pgrx dependencies cover only historical PostgreSQL majors and do not mitigate the design-level command-execution flaw.
