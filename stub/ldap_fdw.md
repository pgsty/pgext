## Usage

Sources:

- [Official repository](https://github.com/guedes/ldap_fdw/tree/b277b085981af0dd014850d5e72d5f43cc2f00cf)
- [Extension control file](https://github.com/guedes/ldap_fdw/blob/b277b085981af0dd014850d5e72d5f43cc2f00cf/ldap_fdw.control)
- [PGXN usage documentation](https://pgxn.org/dist/ldap_fdw/0.1.1/)

ldap_fdw is an experimental, read-only foreign data wrapper that exposes LDAP entries as two-column PostgreSQL rows. The 0.1.1 documentation targets PostgreSQL 9.2 and explicitly says not to use the software in production. It lacks TLS, paging, write support, and a modern security model.

### Core Workflow

Define an LDAP server and binding credentials, then create a foreign table with exactly `dn` and `object_body` text columns:

```sql
CREATE EXTENSION ldap_fdw;

CREATE SERVER directory
  FOREIGN DATA WRAPPER ldap_fdw
  OPTIONS (address 'ldap.internal', port '389');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER directory
  OPTIONS (user_dn 'cn=reader,dc=example,dc=org', password 'secret');

CREATE FOREIGN TABLE directory_entries (
  dn text,
  object_body text
)
SERVER directory
OPTIONS (
  base_dn 'ou=people,dc=example,dc=org',
  attributes 'uid,cn,mail'
);

SELECT dn, object_body
FROM directory_entries
WHERE dn = 'uid=alice,ou=people,dc=example,dc=org';
```

### Important Objects

- `ldap_fdw` is the wrapper and uses a synchronous LDAP client.
- `address` and `port` are server options; the default port is 389.
- `user_dn` and `password` are user-mapping options.
- `base_dn` is required for a foreign table, while `attributes` limits returned attributes.
- `dn` is the entry distinguished name and `object_body` is Hstore-compatible text.

### Security and Query Limits

Only equality on `dn` is pushed down. Other scans perform a one-level search below the base DN, without paging or SQL LIMIT pushdown, so an LDAP server result cap can silently truncate results. The implementation uses a simple bind over an unencrypted connection and stores credentials in PostgreSQL catalogs. Keep it away from real credentials and sensitive directory attributes; use a maintained integration with TLS and paging instead.
