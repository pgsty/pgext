## Usage

Sources:

- [Upstream source 1](https://github.com/emberian/dregg/blob/4caff0718e0a34b9a76281008ebf908e965f856f/pg-dregg/README.md)
- [Upstream source 2](https://github.com/emberian/dregg/blob/4caff0718e0a34b9a76281008ebf908e965f856f/pg-dregg/pg_dregg.control)

`pg_dregg` 0.1.0 is an experimental PostgreSQL adapter for object-capability authorization and row-level security.

### Core Workflow

A DBA provisions the issuer trust key; applications present an issued capability token through `dregg.token`. RLS predicates call `dregg_admits` for an operation and row identifier. `dregg_cap_explain` exposes denial reasoning and `dregg_revoke` supports credential revocation. Enable and, when needed, force RLS on the protected tables.

```sql
CREATE EXTENSION pg_dregg;
SELECT dregg_issuer_status();
```

### Operational Boundaries

The documented live adapter targets PostgreSQL 18 and requires superuser installation. Missing or malformed trust configuration denies authorization. Token issuance and attenuation happen outside this minimal SQL example; do not substitute illustrative keys or tokens in production. Optional proof attestation fails closed when its verifier is absent. Upstream formal-verification claims are not an independent security audit of an installation.
