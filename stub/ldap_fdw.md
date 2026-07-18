## Usage

Sources:

- [Official extension control file](https://github.com/guedes/ldap_fdw/blob/b277b085981af0dd014850d5e72d5f43cc2f00cf/ldap_fdw.control)
- [Official upstream documentation](https://pgxn.org/dist/ldap_fdw/0.1.1/doc/ldap_fdw.html)
- [Official PGXN distribution page](https://pgxn.org/dist/ldap_fdw/)

`ldap_fdw` — Foreign data wrapper for querying LDAP servers

The reviewed catalog snapshot records version `0.1.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "ldap_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'ldap_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
