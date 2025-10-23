---
title: Create
description: Create & Enable PostgreSQL Extension
icon: PackageOpen
---


## Quick Start

You can enable (create) extension using the [**`CREATE EXTENSION`**](https://www.postgresql.org/docs/current/sql-createextension.html) statement:

```sql tab="vector"
CREATE EXTENSION vector; -- no explicit loading required
```
```sql tab="timescaledb"
CREATE EXTENSION timescaledb; -- explicit loading required
```

Extensions need to be [**installed**](/usage/install) first, some extension also require [**preloading**](/usage/config) before using.

Some extensions have dependencies on other extensions.
In such cases, you can either install the dependencies first
or use the `CASCADE` clause to install all dependencies at once.

```sql
CREATE EXTENSION documentdb CASCADE; -- create documentdb extension and all its dependencies
```

You can also provision extension with Pigsty, which will automatically create the extensions for you.


------

## Configure

Extensions (database logical objects) are logically part of [**PostgreSQL databases**](/docs/pgsql/db).
In Pigsty, you can specify which extensions to be created in a database with [**`pg_databases`**](/docs/pgsql/param#pg_databases) parameter.

```yaml
pg_databases:
  - { name: meta ,extensions: [ vector, postgis, timescaledb ] }
```

But you can explicitly specify extension details with the `object` format, like create them in a specific schema.
Or install a specific version. Here's a complete example ([self-hosting supabase](/docs/app/supabase)):

```yaml
pg_databases:
  - name: postgres
    baseline: supabase.sql
    schemas: [ extensions ,auth ,realtime ,storage ,graphql_public ,supabase_functions ,_analytics ,_realtime ]
    extensions:                                 # Extensions to be enabled in the postgres database
      - { name: pgcrypto  ,schema: extensions } # cryptographic functions
      - { name: pg_net    ,schema: extensions } # async HTTP
      - { name: pgjwt     ,schema: extensions } # json web token API for postgres
      - { name: uuid-ossp ,schema: extensions } # generate universally unique identifiers (UUIDs)
      - { name: pgsodium        }               # pgsodium is a modern cryptography library for Postgres.
      - { name: supabase_vault  }               # Supabase Vault Extension
      - { name: pg_graphql      }               # pg_graphql: GraphQL support
      - { name: pg_jsonschema   }               # pg_jsonschema: Validate json schema
      - { name: wrappers        }               # wrappers: FDW collections
      - { name: http            }               # http: allows web page retrieval inside the database.
      - { name: pg_cron         }               # pg_cron: Job scheduler for PostgreSQL
      - { name: timescaledb     }               # timescaledb: Enables scalable inserts and complex queries for time-series data
      - { name: pg_tle          }               # pg_tle: Trusted Language Extensions for PostgreSQL
      - { name: vector          }               # pgvector: the vector similarity search
      - { name: pgmq            }               # pgmq: A lightweight message queue like AWS SQS and RSMQ
```

--------

## Define Extension

The `extensions` field is a list of extension (name or object) to be created in the database.
It will be created under the first schema in dbsu's search_path, (usually the `public` schema).

Here, the `extensions` in the database object is a list where each element can be:

- A simple string representing the extension name, such as `vector`
- Alternatively, A dictionary that may contain the following fields can be used:
  - `name`: Extension name, **REQUIRED**, beware it may differ from the [**extension package name**](/usage/pkg).
  - `schema`: Schema for installing the extension, **OPTIONAL**, defaults to the first schema in the current dbsu search path, usually the default `public`.
  - `version`: Specifies the extension version, **OPTIONAL**, defaults to the latest version, rarely used.

If the database doesn’t exist yet, the extensions defined here will be automatically created when [**creating a cluster**](/docs/pgsql/admin#create-cluster) or [**creating a database**](/docs/pgsql/admin#create-database) through Pigsty.

Re-creating database with non-trivial baseline schema may be dangerous (if you put some `DROP` there)
So for existing clusters / databases, it's advised to use your own schema migration tool to manage extensions. (pgadmin, psql, bytebase, flyway, sqlitch,...)
But it's helpful to enlist them in the [config inventory](/docs/config/inventory) for bookkeeping purposes. (So if you want to fork this cluster, it includes these extensions)



------

## Default Extension

Some built-in extensions and one special [`pg_repack`](/e/pg_repack) are created by default in Pigsty.

These extensions are defined by [**`pg_default_extensions`**](/docs/pgsql/param#pg_default_extensions), created in the `template1` database and the `postgres` database by default.
Newly created databases will inherit these extensions from `template1`, so you don't need to create them again.

```yaml
pg_default_extensions:
  - { name: pg_stat_statements ,schema: monitor }
  - { name: pgstattuple        ,schema: monitor }
  - { name: pg_buffercache     ,schema: monitor }
  - { name: pageinspect        ,schema: monitor }
  - { name: pg_prewarm         ,schema: monitor }
  - { name: pg_visibility      ,schema: monitor }
  - { name: pg_freespacemap    ,schema: monitor }
  - { name: postgres_fdw       ,schema: public  }
  - { name: file_fdw           ,schema: public  }
  - { name: btree_gist         ,schema: public  }
  - { name: btree_gin          ,schema: public  }
  - { name: pg_trgm            ,schema: public  }
  - { name: intagg             ,schema: public  }
  - { name: intarray           ,schema: public  }
  - { name: pg_repack } # <-- The only 3rd-party extension created by default
```

One extra default schema `monitor` is defined by [`pg_default_schemas`](/docs/pgsql/param#pg_default_schemas) is also created by default.
Which is used to contain monitoring related extensions, tables, functions and views.

There are three **3rd-party** extensions that are available by default in Pigsty:

|           Extension           | What                          | Where                                                                                                   |
|:-----------------------------:|-------------------------------|---------------------------------------------------------------------------------------------------------|
| [`pg_repack`](/e/pg_repack) | Online Bloat Control  Tools   | in the [`pg_default_extensions`](/docs/pgsql/param#pg_default_extensions)                               |
|  [`wal2json`](/e/wal2json)  | Changing data capture in JSON | extension [without DDL](/list/attr#without-ddl), [install](/usage/install) means available |
|    [`vector`](/e/vector)    | vector data type & indexes    | in [`pg_databases`](/docs/pgsql/param#pg_databases)  as an example                                      |

The [`pg_repack`](/e/pg_repack) extension is an important utility for maintaining bloat tables online.

[**`vector`**](/e/vector) is a very popular extension for RAG,
It is [**installed**](/usage/install) by default (in the `pgsql-main` [alias](/usage/pkg)) and created in the placeholder `meta` database in most [config template](/docs/config/template).

The [`wal2json`](/e/wal2json) is another important extension for Changing Data Capture (CDC). It is [installed](/usage/install) by default, but it is an [**extension without DDL**](#extension-without-ddl),
So you don't need to `CREATE` it explicitly.


------

## Extension without DDL

[Extension without DDL](/list/attr#without-ddl) does not require the `CREATE EXTENSION` command to work

PostgreSQL extensions typically consist of three parts: a required control file, optional SQL files, and optional libraries.
If an extension does not have SQL file, `CREATE EXTENSION` command is not needed.

| Component    | Description                                           | Required     |
|--------------|-------------------------------------------------------|:------------:|
| Control file | Key metadata, name, dependencies, schema, version,... | **REQUIRED** |
| SQL file     | SQL DDL statements, Types, Functions, etc...          | **OPTIONAL** |
| Library file | binary shared libraries (`.so`, `.dylib`, `.dll`)     | **OPTIONAL** |

Since SQL / LIB files are optional, there are four possible combinations of extension types:

| **[LOAD](/usage/config) / DDL** | Requires `CREATE EXTENSION` | Doesn't require `CREATE EXTENSION` |
|------------------------------------------|-----------------------------|------------------------------------|
| **Requires `LOAD`**                      | Extensions using hooks      | Headless extensions                |
| **Doesn't Require `LOAD`**               | Extensions not using hooks  | Logical decoding output plugins    |

