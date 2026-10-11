## Usage

Sources:

- [Compiler prerequisites](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/docs/DEVELOPER-QUICKSTART.md)

- [README.md](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [crates/plts/Cargo.toml](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/Cargo.toml)
- [crates/plts/src/lib.rs](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/src/lib.rs)
- [crates/plts/plts.control](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/plts.control)

`plts` provides the procedural language and compiler/artifact APIs used by Stopgap. It is an untrusted extension; installation requires a superuser. Cargo features cover PostgreSQL 14–18, with 17 selected by default.

### Core Workflow

```ini
shared_preload_libraries = 'plts'
```

```sql
CREATE EXTENSION plts;
SELECT plts.typecheck_ts('export default () => 42;');
```

### Operational Boundaries

`plts.compile_ts()` compiles TypeScript, `plts.typecheck_ts()` checks it, `plts.compile_and_store()` saves an artifact and `plts.get_artifact()` retrieves one. Standalone modules use a default export; Stopgap application modules use named exports. Database calls execute in the SQL transaction. Runtime availability depends on the source build features and compiler/runtime assets; the README also describes an ongoing compiler migration, so verify the pinned runtime contract before deployment. Compiler service and backend-local runtime settings are distinct; do not assume that creating the extension alone enables every build configuration.

Compilation and typechecking require `shared_preload_libraries = 'plts'` and a server restart to start the shared TSGo compiler worker.
