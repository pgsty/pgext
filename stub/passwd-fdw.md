## Usage

Sources:

- [Official PGXN README](https://pgxn.org/dist/passwd-fdw/0.4.0/README.html)
- [Official PGXN control file](https://api.pgxn.org/src/passwd-fdw/passwd-fdw-0.4.0/passwd-fdw.control)
- [Official PGXN distribution](https://pgxn.org/dist/passwd-fdw/)

`passwd-fdw` exposes the host's Unix user and group databases through foreign tables. The data normally comes from `/etc/passwd` and `/etc/group`, but the system's Name Service Switch may resolve it from LDAP or another provider. This can disclose host accounts and services, so access is an administrative privilege boundary.

### Core Workflow

```sql
CREATE EXTENSION "passwd-fdw";

CREATE SERVER passwd_svr
    FOREIGN DATA WRAPPER passwd_fdw;

CREATE FOREIGN TABLE passwd_users (
    name text,
    passwd text,
    uid integer,
    gid integer,
    gecos text,
    dir text,
    shell text
) SERVER passwd_svr;

CREATE FOREIGN TABLE passwd_groups (
    name text,
    gid integer,
    members text[]
) SERVER passwd_svr;

SELECT name, uid, gid, dir, shell
FROM passwd_users
ORDER BY uid;
```

The extension name contains a hyphen and therefore needs quoting; the installed wrapper name is `passwd_fdw`.

### Data and Query Boundary

The user table mirrors password-database fields, while the group table exposes group names, IDs, and member arrays. Results represent the database server host's current name-service view, not the client machine and not a transactionally stable PostgreSQL catalog.

The upstream documentation does not describe write callbacks, pushdown, server options, or user mappings. Treat the wrapper as a read-only local-system inventory and validate behavior on the packaged build instead of adding undocumented options.

### Security and Compatibility

Revoke public access to the foreign server and tables; grant only the columns and rows required by a controlled administrative workflow. Even without password hashes, account names, UIDs, home directories, shells, group membership, and NSS latency/failures are sensitive operational information.

Version `0.4.0` is an old, minimally documented implementation. Test target PostgreSQL compatibility, concurrent NSS lookups, LDAP outages, unusual field contents, large group membership, and dump/restore. Foreign-table definitions restore, but the underlying host identities are environment-specific and must be revalidated on every server.
