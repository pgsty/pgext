## Usage

Sources:

- [pgauditlogtofile v1.8.5 README](https://github.com/fmbiete/pgauditlogtofile/blob/v1.8.5/README.md)
- [Changes from v1.8.4 to v1.8.5](https://github.com/fmbiete/pgauditlogtofile/compare/v1.8.4...v1.8.5)

pgauditlogtofile is a pgAudit add-on that routes pgAudit records to a dedicated CSV or JSON file. Use it to separate audit retention and access controls from the ordinary PostgreSQL server log while keeping pgAudit's event selection and semantics.

### Preload and Create the Extension

Load pgAudit first, then pgauditlogtofile:

    shared_preload_libraries = 'pgaudit,pgauditlogtofile'

Restart PostgreSQL, then install both extensions in the postgres database:

    CREATE EXTENSION pgaudit;
    CREATE EXTENSION pgauditlogtofile;

Upstream recommends creating pgauditlogtofile only in the postgres database, not independently in every application database.

### Configure Audit Files

    pgaudit.log_directory = 'log'
    pgaudit.log_filename = 'audit-%Y%m%d_%H%M.log'
    pgaudit.log_format = 'csv'
    pgaudit.log_rotation_age = 1440
    pgaudit.log_file_mode = 0600

An empty pgaudit.log_directory or pgaudit.log_filename disables the separate target and lets records fall back to the normal server logger. Relative directories are resolved under the PostgreSQL data directory.

### Compression

Version 1.8 supports compressed audit files:

    pgaudit.log_compression = 'zstd'
    pgaudit.log_compression_level = 6

pgaudit.log_compression accepts off, gzip, lz4, or zstd when the corresponding support is available. The level range is 0 through 22, but valid and useful levels depend on the selected algorithm. Compression consumes backend CPU, so test both log throughput and rotation latency.

### Parameter Index

- pgaudit.log_format: csv or json output.
- pgaudit.log_directory and pgaudit.log_filename: destination and strftime-style filename.
- pgaudit.log_file_mode: permissions for newly created files.
- pgaudit.log_rotation_age: time-based rotation interval in minutes.
- pgaudit.log_compression and pgaudit.log_compression_level: compression method and effort.
- pgaudit.log_connections and pgaudit.log_disconnections: include connection lifecycle events when PostgreSQL's matching log settings are enabled.
- pgaudit.log_execution_time and pgaudit.log_execution_memory: add execution measurements; these require a restart.
- pgaudit.log_autoclose_minutes: experimental inactivity-based file-handler close.

### Rotation and Operations

A PostgreSQL configuration reload rotates the audit file. The extension's background worker can signal backends to close audit file handles; pg_rotate_logfile() does not rotate the independent audit file.

Version 1.8.5 improves background-worker signaling, hook restoration, and PostgreSQL 19 build compatibility. It does not introduce a required configuration migration from 1.8.4.

### Caveats

- File separation is not retention management. Ship, rotate, protect, and expire audit files explicitly.
- Ensure the PostgreSQL operating-system account can create the destination and that file permissions meet the audit policy.
- Abrupt backend or host failure can leave the last compressed file incomplete; validate ingestion behavior.
- Enabling timing, memory, connection, or verbose pgAudit classes can materially increase overhead and log volume.
