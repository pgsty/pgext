## Usage

Sources:

- [Official version 0.1.7 README](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/README.md)
- [Official control file](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/pgx_uuidv7.control)
- [Official version 0.1.7 release](https://github.com/kaznak/pgx_uuidv7/releases/tag/v0.1.7)

`pgx_uuidv7` generates time-ordered UUID version 7 values and extracts their embedded timestamp or version. It is useful for UUID primary keys with better insertion locality than random UUIDv4, while retaining a standard `uuid` storage type. Version `0.1.7` targets PostgreSQL 17 and 18.

### Core Workflow

Create the extension and use the current-time generator as a column default:

```sql
CREATE EXTENSION pgx_uuidv7;

CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v7_now(),
  payload jsonb NOT NULL
);

INSERT INTO events (payload) VALUES ('{"kind":"signup"}');

SELECT id, uuid_get_version(id), uuid_to_timestamptz(id)
FROM events;
```

UUID values can also be cast directly to `timestamptz` for timestamp filtering.

### API

- `uuid_generate_v7_now()`: generate a UUIDv7 for the current time.
- `uuid_generate_v7(timestamptz)`: generate for an explicit timestamp.
- `uuid_generate_v7_at_interval(interval)`: generate at an offset from the current time.
- `uuid_get_version(uuid)`: extract the UUID version.
- `uuid_to_timestamptz(uuid)`: extract the UUIDv7 timestamp.

On PostgreSQL versions earlier than 18, the build also exposes compatibility aliases `uuidv7()`, `uuidv7(interval)`, `uuid_extract_version(uuid)`, and `uuid_extract_timestamp(uuid)`. They are deliberately omitted from the PostgreSQL 18 build to avoid conflicts with native functions.

### Caveats

The timestamp embedded in UUIDv7 is observable; do not treat the identifier as hiding creation time. Time ordering improves locality but does not replace a business timestamp or guarantee global generation order across clocks. Validate casts and aliases when moving between PostgreSQL 17 and 18, because the native PostgreSQL 18 surface replaces the compatibility names.
