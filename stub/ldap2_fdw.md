## Usage

Sources:

- [README.md](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/README.md)
- [ldap2_fdw--1.0.sql](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw--1.0.sql)
- [ldap2_fdw.c](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw.c)
- [ldap2_fdw.control](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw.control)

`ldap2_fdw` 1.0 maps LDAP entries to foreign tables and supports writes. The author explicitly warns against production use because testing, memory management and error handling remain incomplete.

### Core Workflow

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION ldap2_fdw;
CREATE SERVER ldap_directory FOREIGN DATA WRAPPER ldap2_fdw
    OPTIONS (uri 'ldap://localhost');
```

### Operational Boundaries

The official setup creates `uuid-ossp` first; generated import definitions can use `uuid_generate_v4()`. Set the server `uri`, user-mapping `username` and `password`, and table `basedn`, `filter` and `scope` for the actual directory. Keep credentials in the user mapping and grant access deliberately. LDAP entries do not always map cleanly to relations, server size limits can truncate results, and remote filtering is limited to a single DN predicate without other conditions. Installation defaults to superuser-only. No supported PostgreSQL-major range or standalone license is declared.
