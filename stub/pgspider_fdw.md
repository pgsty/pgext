## Usage

Sources:

- [Official README](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/README.md)
- [Extension control file](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/pgspider_fdw.control)
- [Installation SQL](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/pgspider_fdw--1.4.sql)
- [Build configuration](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/Makefile)

`pgspider_fdw` connects PGSpider nodes and supports the fork's optional compressed transfers through cloud functions. Use the matching PGSpider distribution; it is a separate extension from `pgspider_core_fdw` and `pgspider_ext`.

### Remote-node Workflow

With compatible PGSpider nodes and an existing remote table, an administrator can create a server and mapping. Substitute the actual host, database and credentials; PGSpider uses port 4813 by default.

```sql
CREATE EXTENSION pgspider_fdw;
CREATE SERVER remote_spider FOREIGN DATA WRAPPER pgspider_fdw
  OPTIONS (host '127.0.0.1', port '4813', dbname 'pgspider');
CREATE USER MAPPING FOR CURRENT_USER SERVER remote_spider
  OPTIONS (user 'app_user', password 'replace-me');
CREATE FOREIGN TABLE remote_t1(i integer, t text, __spd_url text)
  SERVER remote_spider OPTIONS (table_name 't1');
SELECT * FROM remote_t1;
```

### Connections and Transfer

`pgspider_fdw_get_connections()` lists cached remote connections; `pgspider_fdw_disconnect(text)` and `pgspider_fdw_disconnect_all()` release them. The optional transfer workflow uses `endpoint`, `proxy`, `batch_size` and fork-specific migration commands, and requires separately deployed compatible cloud functions. Follow the official deployment recipe before using that path; the SQL example above does not configure it.

### Dependencies and Limits

The module links libpq, libcurl and LZ4. SQL/control version `1.4` is separate from the surrounding PGSpider kernel version. Grant only needed server/table privileges and protect user-mapping credentials. Operations depend on remote availability and FDW transaction behavior; no distributed atomic-commit guarantee is implied. The source provides no current stock PostgreSQL compatibility matrix or required preload setting.

### Installation Objects and Name Conflicts

The 1.4 install script also creates procedure `pgspider_create_or_replace_stub(text, text, regtype)`, types `mysql_string_type`, `time_unit`, `path_value`, and several hundred pushdown stub functions and aggregates. These routines require `plpgsql` and raise errors when executed locally; they do not implement the corresponding remote functions locally.

The script uses `CREATE OR REPLACE FUNCTION` and `CREATE OR REPLACE AGGREGATE`, and some type-existence checks use only the type name without a schema qualifier. Review same-name functions in the installation schema and same-name types across the database before installation to avoid collisions with application objects.
