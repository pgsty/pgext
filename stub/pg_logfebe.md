## Usage

Sources:

- [Official protocol and configuration documentation](https://github.com/logplex/pg_logfebe/blob/bac15ef2386fc70197ba889e23027bffe00d0f5b/README.rst)
- [Extension control file](https://github.com/logplex/pg_logfebe/blob/bac15ef2386fc70197ba889e23027bffe00d0f5b/pg_logfebe.control)

`pg_logfebe` is a preload-only logging module that sends structured PostgreSQL log records to another process over a Unix-domain socket using PostgreSQL front-end/back-end framing. It has no extension SQL objects, so do not run `CREATE EXTENSION` for it.

### Core Workflow

Run a compatible receiver on a protected Unix socket, configure PostgreSQL, and restart the server to load the module.

```conf
shared_preload_libraries = 'pg_logfebe'
logfebe.unix_socket = '/run/pg_logfebe/log.sock'
logfebe.identity = 'primary-db'
```

The receiver first gets version and identity messages, followed by a stream of structured log records. The identity is an operator-chosen label that lets one receiver distinguish PostgreSQL instances.

### Protocol and Operational Boundaries

- `logfebe.unix_socket` selects the receiving Unix-domain socket.
- `logfebe.identity` supplies the instance label sent during startup.
- `shared_preload_libraries` is required because the module installs logging hooks during server startup.

The wire framing resembles PostgreSQL's FE/BE protocol, but the documented log stream is a separate, versioned protocol and needs a purpose-built receiver. Protect the socket directory and downstream logs because structured records may contain query text, error details, and other sensitive database information. Configuration changes to the preload list require a restart; there is no per-database DDL or SQL API.
