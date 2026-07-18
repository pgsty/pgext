## Usage

Sources:

- [Pinned upstream README](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/README.md)
- [Version 1.0 application SQL](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/siguanet_quest--1.0.sql)
- [Pinned extension control file](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/siguanet_quest.control)
- [SIGUANET database schema project](https://github.com/labgeo/siguanet-dbsetup)

siguanet_quest version 1.0 is an application-specific reporting layer for a SIGUANET university built-asset geodatabase. It installs the fixed quest schema with hundreds of Spanish-named PL/pgSQL functions, views, and composite types for spatially aggregating floor area, rooms, staff, buildings, floors, departments, and activity groups.

### Installation and representative query

It requires an already populated SIGUANET database whose tables and views match the companion schema project, plus PostgreSQL 9.2+ and PostGIS 2.0+ according to the historical README:

```sql
CREATE EXTENSION siguanet_quest;

SELECT quest_edificio_superficie('0037');

SELECT *
FROM quest_edificio_obtenerplantas('0037')
ORDER BY indice;
```

These examples are meaningful only when building 0037 and the expected SIGUANET objects exist. Administrators are expected to populate quest_adminroles for roles authorized to view unencrypted protected data such as Spanish NIF identifiers.

### Application and privacy boundary

This is not a general PostGIS analytics library. Its SQL assumes exact public-schema tables, column conventions, activity codes, Spanish institutional meanings, and companion views. It creates a large surface of dynamic SQL, views, types, comments, and privilege-sensitive reporting paths. Installation on a different schema or partially compatible database can fail midway or return misleading results.

The repository has not changed since 2013 and is classified as abandoned. Its stated PostgreSQL and PostGIS minimums are historical, not evidence for current majors. Before using it, restore a representative SIGUANET database in isolation, inventory every created object and dependency, test each exposed function, review dynamic SQL and protected-identifier paths, and replace the old application-level admin table with current role and row-security controls where appropriate.
