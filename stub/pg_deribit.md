## Usage

Sources:

- [README.md](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/README.md)
- [pg_deribit.control](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/pg_deribit.control)
- [Makefile](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/Makefile)
- [Dockerfile](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/Dockerfile)
- [sql/functions/environment.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/environment.sql)
- [sql/functions/auth.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/auth.sql)
- [sql/functions/internal_url_endpoint.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/internal_url_endpoint.sql)
- [sql/endpoints/public_get_currencies.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/endpoints/public_get_currencies.sql)

`pg_deribit` 1.0 exposes Deribit JSON-RPC endpoints through SQL. It requires the Omnigres HTTP extensions; the upstream container uses PostgreSQL 17. Public calls can read market metadata, while authenticated endpoints can perform account operations.

### Public Testnet Workflow

```sql
CREATE EXTENSION pg_deribit CASCADE;
SELECT deribit.enable_test_net();
SELECT currency FROM deribit.public_get_currencies() ORDER BY currency;
```

### Dependencies and Configuration

The non-relocatable extension uses the `deribit` schema and requires `pgcrypto`, `omni_http` and `omni_httpc`. Installation uses the default superuser restriction. It contains SQL functions and types, with no extension-specific shared library or preload setting. The dependencies must be installed and usable.

Production is the default endpoint. `deribit.enable_test_net()` and `deribit.disable_test_net()` change `deribit.set_test_net` for the current session; enable testnet again in each fresh connection. Calls need outbound network access and depend on the remote API.

### Authenticated Calls

`deribit.set_client_auth(client_id, client_secret)` and `deribit.set_access_token_auth(client_id, client_secret, access_token, refresh_token)` set session credentials; `deribit.get_auth()` reads them. Grant SQL access deliberately and keep credentials out of query logs and shared examples. Private wrappers include account-changing operations: a database transaction rollback does not undo an already completed remote action. Review individual endpoint signatures before granting access.

Credentials are stored in session settings. The upstream setters interpolate values with `%s` instead of SQL-literal quoting, so never pass untrusted credential strings. Functions retain default `PUBLIC` execution permissions; restrict them for your deployment. Private request/response data is persisted in `deribit.internal_archive` and needs an explicit retention policy.
