## Usage

Sources:

- [Pinned upstream README](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/README.md)
- [Pinned version 1.0 installation SQL](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/pg_headerkit--1.0.sql)
- [Pinned extension control file](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/pg_headerkit.control)

pg_headerkit version 1.0 is a pure-SQL PostgREST header helper. It creates the fixed hdr schema, allow-list and deny-list tables, and functions for reading the request.headers setting, extracting proxy and client fields, testing IP lists, and recognizing simple mobile user-agent patterns.

### Intended use after correction

The following illustrates the intended API, but the pinned 1.0 installation script must be corrected before this sequence can run successfully:

```sql
CREATE EXTENSION pg_headerkit;

INSERT INTO hdr.allow_list (ip)
VALUES ('203.0.113.10');

SELECT hdr.header('user-agent');
SELECT hdr.ip();
SELECT hdr.in_allow_list('203.0.113.10'::inet);
```

The functions only expose request context supplied by PostgREST or another trusted gateway. They do not themselves enforce rate limits, routing, request filtering, or authorization; applications must call them from policies or functions that implement those decisions.

### Current defects and security

At the reviewed commit, each of the hdr.projectref and hdr.ref function bodies contains two adjacent SELECT statements without a separator. Consequently, the distributed version 1.0 SQL has a syntax error and CREATE EXTENSION pg_headerkit fails atomically. Patch and review the installation SQL or wait for an upstream correction before packaging it.

Several functions that read request settings or allow/deny tables are incorrectly declared immutable. PostgreSQL may pre-evaluate or cache such results, so they must use appropriate volatility before they are trusted for security decisions. The IP helper trusts the first x-forwarded-for value; accept it only when a controlled reverse proxy overwrites that header. Design explicit schema, table, and function privileges, because the upstream SQL supplies no complete role model and its functions run with invoker rights.

The README lists rate limiting and access-control capabilities as uses that can be built with the helpers, not installed enforcement. Audit corrected SQL, revoke unneeded public execution, and test every policy through the actual PostgREST role and connection-pooling mode.
