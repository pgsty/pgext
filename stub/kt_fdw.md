## Usage

Sources:

- [Archived upstream README at the reviewed commit](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/README.md)
- [FDW SQL definition](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/sql/kt_fdw.sql)
- [Extension control file](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/kt_fdw.control)

`kt_fdw` exposes a Kyoto Tycoon key-value server as a two-column PostgreSQL foreign table. It supports `SELECT`, `INSERT`, `UPDATE`, and `DELETE` against text keys and values.

```sql
CREATE EXTENSION kt_fdw;

CREATE SERVER kt_server
  FOREIGN DATA WRAPPER kt_fdw
  OPTIONS (host '127.0.0.1', port '1978', timeout '-1');

CREATE USER MAPPING FOR PUBLIC SERVER kt_server;
CREATE FOREIGN TABLE cache_entry (key text, value text)
  SERVER kt_server;

SELECT * FROM cache_entry WHERE key = 'example';
```

Transactional mode requires Kyoto Tycoon built with Lua support and the bundled `transactions.lua` loaded by `ktserver`. Builds can remove `USE_TRANSACTIONS` to disable that integration.

### Caveats

The repository now lives in Cloudflare's archive organization, its last source change dates to 2013, and upstream does not document compatibility with current PostgreSQL or Kyoto Tycoon releases. The remote service is outside PostgreSQL durability and access-control boundaries; test failure, timeout, and transaction behavior before using the wrapper with important data.
