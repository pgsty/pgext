## Usage

Sources:

- [Official README](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/README.md)
- [Extension control file](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.control)
- [FDW implementation](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.c)

couchdb_fdw is an old, experimental, read-only foreign data wrapper for exposing CouchDB documents as PostgreSQL rows. Upstream calls it beta, reports known bugs, and says it was tested only with PostgreSQL 9.1-era software and macOS 10.6. Treat it as historical source, not a production integration.

### Core Workflow

After installing the required libcurl and YAJL libraries, define the server, credentials, and one foreign table. `_doc` is a reserved column name that returns the entire document JSON; other columns map top-level JSON attributes.

```sql
CREATE EXTENSION couchdb_fdw;

CREATE SERVER couch
  FOREIGN DATA WRAPPER couchdb_fdw
  OPTIONS (address '127.0.0.1', port '5984');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER couch
  OPTIONS (username 'reader', password 'secret');

CREATE FOREIGN TABLE couch_docs (
  _id text,
  _rev text,
  title text,
  _doc text
)
SERVER couch
OPTIONS (database 'example');

SELECT _id, title FROM couch_docs WHERE _id = 'document-id';
```

### Important Objects

- `couchdb_fdw` is the foreign-data wrapper.
- `address` and `port` are server options; their defaults are 127.0.0.1 and 5984.
- `username` and `password` are user-mapping options.
- `database` is the foreign-table option and defaults to test.
- `_id`, `_rev`, and `_doc` are special document columns.

### Security and Query Limits

The implementation constructs a plain HTTP URL containing the username and password and has no TLS or modern authentication support. Equality on `_id` is the only meaningful pushdown, with optional `_rev`; an unqualified scan first obtains all document IDs and then fetches documents individually. It maps only top-level attributes and implements no writes. Credentials remain in PostgreSQL catalogs, remote I/O is synchronous, and failures can interrupt the backend query. Do not expose sensitive credentials or untrusted CouchDB endpoints through this obsolete code.
