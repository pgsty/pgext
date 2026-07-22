## Usage

Sources:

- [Upstream README](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/README.md)
- [FDW implementation](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/src/lib.rs)
- [Rust package manifest](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/Cargo.toml)

`orb_fdw` version `0.13.3` is a read-only foreign data wrapper for querying selected Orb billing objects through the Orb API.

### Core Workflow

```sql
CREATE EXTENSION orb_fdw;
CREATE FOREIGN DATA WRAPPER orb_wrapper
  HANDLER orb_fdw_handler VALIDATOR orb_fdw_validator;
CREATE SERVER my_orb_server FOREIGN DATA WRAPPER orb_wrapper
  OPTIONS (api_key '<orb secret key>');

CREATE FOREIGN TABLE orb_customers (
  user_id text,
  organization_id text,
  email text,
  created_at timestamp,
  attrs jsonb
) SERVER my_orb_server OPTIONS (object 'customers');

SELECT user_id, email FROM orb_customers;
```

The server accepts `api_key`; when omitted, the server process reads `ORB_API_KEY`. A foreign table must set `object` to one of `customers`, `subscriptions`, or `invoices`. Define only documented column names and compatible PostgreSQL types; `attrs` retains the full response object.

### Operational Boundaries

The implementation provides scan callbacks only, so treat it as read-only. It paginates the remote API, while PostgreSQL applies filters, sorting, and limits locally; a selective-looking query can therefore make many network requests.

An API key stored in server options is visible to sufficiently privileged catalog readers; prefer the process environment where operationally acceptable and restrict server/table ownership. Plan for HTTPS failures, Orb API schema changes, quotas, rate limits, and latency. The upstream README describes an earlier `0.13.0` prerequisite while the reviewed Rust manifest is `0.13.3`; test the packaged build rather than assuming README compatibility.
