## Usage

Sources:

- [Official documentation](https://github.com/codybrom/pg_fedi/blob/2616f1f1e9566934bf53261d16fcf4c051d4ce0d/README.md)
- [Extension control file](https://github.com/codybrom/pg_fedi/blob/2616f1f1e9566934bf53261d16fcf4c051d4ce0d/pg_fedi.control)
- [Official repository](https://github.com/codybrom/pg_fedi)

`pg_fedi` Pure-SQL ActivityPub federation primitives for PostgreSQL and pg_tle environments.

### Enablement

Install the files for the intended server, then create `pg_fedi` in the target database:

```sql
CREATE EXTENSION pg_fedi CASCADE;
```

The reviewed control or official workflow requires `pgcrypto`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Create a local actor (app layer generates RSA keypair, passes it in)
SELECT ap_create_local_actor('alice', 'Alice', 'Hello!', public_pem, private_pem);

-- Post a note (queues delivery to followers)
SELECT ap_create_note('alice', '<p>Hello, fediverse!</p>');

-- Process an inbound activity
SELECT ap_process_inbox_activity('{"type":"Follow", ...}'::jsonb);

-- WebFinger lookup
SELECT ap_webfinger('acct:alice@myinstance.social');

-- Actor profile as ActivityStreams JSON-LD
SELECT ap_serialize_actor('alice');

-- Timelines
SELECT * FROM ap_public_timeline;
SELECT * FROM ap_home_timeline('alice');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ap_set_setting` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_process_inbox_activity` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_serialize_actor` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_webfinger` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_create_local_actor` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_create_note` | FUNCTION | Callable function from the reviewed install surface. |
| `ap_deliveries` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `ap_delivery_failure` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
