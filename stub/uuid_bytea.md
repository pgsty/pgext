## Usage

Sources:

- [postgresql/uuid_bytea/README](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/README)
- [postgresql/uuid_bytea/uuid_bytea--1.0.sql](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea--1.0.sql)
- [postgresql/uuid_bytea/uuid_bytea.c](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea.c)
- [postgresql/uuid_bytea/uuid_bytea.control](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea.control)

`uuid_bytea` is a small experimental C extension for explicit conversion between uuid and bytea. The upstream author describes it as toy code tested on PostgreSQL 9.1.

### Core Workflow

```sql
CREATE EXTENSION uuid_bytea;
SELECT uuid_to_bytea('0fcc6350-118d-11e4-a559-7de5338eb025'::uuid);
SELECT bytea_to_uuid(decode('0fcc6350118d11e4a5597de5338eb025','hex'));
```

### Operational Boundaries

`uuid_to_bytea` returns the UUID’s 16 bytes; `bytea_to_uuid` converts a 16-byte value back. Explicit casts are also installed. Do not supply arbitrary-length binary values or assume that UUID textual byte order describes an external protocol.

Installation requires a superuser. No preload is specified. The historical test environment does not establish current compatibility, and the reviewed directory provides no license declaration.
