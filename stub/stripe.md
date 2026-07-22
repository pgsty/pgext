## Usage

Sources:

- [Official README for v0.2.1](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/README.md)
- [Official control file for v0.2.1](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/stripe.control)
- [Official changelog for v0.2.1](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/CHANGELOG.md)
- [Official coupon SQL for v0.2.1](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/sql/stripe_coupons.sql)

`stripe` is an experimental, generated PL/Python extension that exposes Stripe REST API resources as PostgreSQL functions and composite types. Every API call leaves the database process and is subject to Stripe credentials, network behavior, rate limits, and external side effects.

The reviewed upstream distribution is `v0.2.1`, but its control file and generated extension SQL still declare extension version `0.0.1`. Pin the distribution commit as well as the SQL extension version; do not infer the generated API surface from `extversion` alone.

### Prerequisites and Read-Only Query

Upstream requires PostgreSQL 14+, Python 3.9+, `plpython3u`, and the matching `stainless_commons_stripe` Python package installed in the interpreter used by PostgreSQL.

```sql
CREATE EXTENSION IF NOT EXISTS plpython3u;
CREATE EXTENSION stripe;

SET stripe.secret_key = '<session-injected-secret>';

SELECT *
FROM stripe_coupons.list()
LIMIT 20;
```

The API is grouped into schemas such as `stripe_accounts`, `stripe_coupons`, `stripe_customers`, `stripe_invoices`, and `stripe_payment_intents`. Resource functions perform retrieval, listing, creation, or other endpoint operations; `make_*` helpers construct nested composite parameters.

### Pagination and Caching

List functions fetch additional pages automatically. Place `LIMIT` directly above the set-returning function as shown so PostgreSQL can stop evaluation early; a limit hidden behind views or other query layers may not be pushed down, causing all pages to be requested and buffered.

For repeated analytics, materialize a read-only pull and refresh it deliberately. `pg_cron` can schedule refreshes but is optional and is not installed by `stripe`.

### Credentials and Network Boundary

The client reads `stripe.secret_key` and `stripe.base_url`. Persisting a secret with database/role settings can expose it to privileged catalog readers and backups; prefer controlled per-session injection, restrict connection logging, and rotate credentials. Changing `stripe.base_url` also creates an outbound destination/SSRF boundary, so restrict who can set it and enforce egress policy outside PostgreSQL.

`plpython3u` is untrusted and its installation requires a superuser boundary. Revoke broad `EXECUTE` privileges on network and mutation functions, then grant only the exact read/write endpoints each application role needs.

### Transaction and Version Caveats

Stripe mutations are external side effects: rolling back a PostgreSQL transaction does not undo a successful remote create, update, refund, or delete. Design idempotency keys, retry rules, reconciliation, and webhook processing around that split transaction boundary; never assume ordinary SQL atomicity covers the API.

The project is explicitly experimental, and `v0.2.0` contained breaking generated-name changes while the SQL extension version remained `0.0.1`. Validate exact function signatures from the pinned SQL, test against a Stripe test account, and monitor request counts, latency, partial failures, pagination, and rate limiting before production use.
