## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/sha/sha-1.1.0/README)
- [Official extension SQL (sha.sql)](https://api.pgxn.org/src/sha/sha-1.1.0/sql/sha.sql)

`sha` — The module implements sha1, sha224, sha256, sha384 sha512 and md5hash datatypes You can apply basic comparison operators on these types and use them with indexes. Both btree and hash indexes are supported. See sql/sha.sql for usage examples. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Requirements and Caveats

- The catalog records version `1.1.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
