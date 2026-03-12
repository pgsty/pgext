


## Usage

> [pg_checksums: Activate/deactivate/verify checksums in offline Postgres clusters](https://github.com/credativ/pg_checksums)

`pg_checksums_ext` is a command-line tool (based on PostgreSQL's built-in `pg_checksums`) that can verify, activate, or deactivate page-level checksums for a PostgreSQL cluster. It extends the built-in tool with online verification, `SIGUSR1` progress toggling, fine-grained progress reporting, and I/O rate limiting.

### Verify Checksums (can be done online)

```bash
pg_checksums_ext -D /path/to/data --check
```

### Enable Checksums (requires clean shutdown)

```bash
pg_checksums_ext -D /path/to/data --enable
```

### Disable Checksums (requires clean shutdown)

```bash
pg_checksums_ext -D /path/to/data --disable
```

### Additional Options

- `-D, --pgdata` -- data directory path
- `--check` / `--enable` / `--disable` -- operation mode
- `--progress` -- show progress reporting
- `--filenode` -- check only a specific filenode
- `--no-sync` -- skip fsync
- `--verbose` -- verbose output
- `--debug` -- debug output
- Send `SIGUSR1` to toggle progress reporting during operation
