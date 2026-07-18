## Usage

Sources:

- [Official extension control file](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.control)
- [Official upstream documentation](https://pgxn.org/dist/kmeans/1.1.0/doc/kmeans.html)
- [Official PGXN distribution page](https://pgxn.org/dist/kmeans/)

`kmeans` — K-means clustering window function for PostgreSQL

The reviewed catalog snapshot records version `1.1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "kmeans";
SELECT extversion
FROM pg_extension
WHERE extname = 'kmeans';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
