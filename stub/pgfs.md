## Usage

Sources:

- [Version 0.0.1 Rust implementation](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/src/lib.rs)
- [Extension control file](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/pgfs.control)
- [Rust package manifest](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/Cargo.toml)

`pgfs` exposes one server-side filesystem operation: `pgfs_copy_dir(source_dir, dest_dir)`. It recursively copies a directory as the PostgreSQL operating-system account. This is a privileged administrative primitive, not a client file-transfer facility.

### Copy a Controlled Directory

Version `0.0.1` declares superuser-only extension installation. After installation, revoke the function's default execution privilege and grant it only to a tightly controlled maintenance role if required:

```sql
CREATE EXTENSION pgfs;
REVOKE EXECUTE ON FUNCTION pgfs_copy_dir(text, text) FROM PUBLIC;

SELECT pgfs_copy_dir(
  '/srv/postgresql/export/source',
  '/srv/postgresql/export/copy'
);
```

The function creates missing destination directories, traverses subdirectories, copies regular files, and returns `true` on success. It returns `false` and emits an informational message for many missing-path, directory-create, read, or copy errors.

### Security and Filesystem Semantics

Both paths are interpreted on the database server and with the database service account's permissions. Never grant this function to untrusted SQL roles or accept paths directly from users. Constrain access with operating-system permissions and use only reviewed absolute paths outside PostgreSQL's data directory and configuration directories.

Existing destination files with the same name can be overwritten. Despite a source comment, the implementation does not explicitly preserve directory or file permissions. Its directory test can follow symbolic links, recursive links can cause unsafe traversal, non-UTF-8 paths can panic, and a partial copy is not rolled back when a later item fails. The function provides no fsync, atomic publish, ownership preservation, checksum, cancellation cleanup, or audit trail. Use operating-system backup/copy tooling for production workflows; if this abandoned extension must be used, rehearse symlinks, failures, disk exhaustion, overwrite behavior, and recovery on disposable data first.
