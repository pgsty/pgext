---
title: Usage
description: Install, Load, create, update PostgreSQL extensions
icon: Package
weight: 800
---


There are unparalleled [**430+**](/list/) extensions available in Pigsty for 14 mainstream [Linux distros](/docs/prepare/linux).

--------

## Overview

It takes 4 steps to deliver an extension: [**downloads**](/usage/download), [**installs**](/usage/install), [**config**](/usage/config), and [**create**](/usage/create):

<Steps>
    <Step>

        [**Download**](#download-extension) : Which extension packages to download

        ```yaml tab="config" title="define which extensions to be downloaded"
        repo_extra_packages: [ postgis, timescaledb, vector ]
        ```
        ```bash tab="apply" title="download package"
        make repo
        ```

    </Step>
    <Step>
        [Install](#install-extension) : Which extensions to be installed

        ```yaml tab="config"
        pg_extensions: [ postgis, pgvector, timescaledb ]
        ```
        ```bash tab="apply"
        ./pgsql.yml -t pg_ext     # install extensions
        ```

    </Step>

    <Step>
        [**Load**](#load-extension) : Which extensions to be pre-loaded

        ```yaml tab="config"
        pg_libs: 'timescaledb, pg_stat_statements, auto_explain'  # add extension to preload libraries (not all extensions need this)
        ```
        ```bash tab="apply" title="edit existing cluster config and reload"
        pg edit-config --force -p shared_preload_libraries='timescaledb, pg_stat_statements, auto_explain'
        ```


    </Step>
    <Step>

        [**Create**](#create-extension) : Create extension in the [database](/docs/pgsql/db)

        ```yaml tab="config"
        pg_databases:
        - { name: meta ,extensions: [ postgis, timescaledb, vector ] }
        ```
        ```sql tab="apply" title="create extension in existing database"
        CREATE EXTENSION postgis CASCADE;
        ```
    </Step>


</Steps>

--------

## Quick Start

You can describe extensions in the [**config inventory**](/docs/config/inventory), and pigsty will
download, install, configure, and enable extensions for you.
This example makes `postgis`, `pgvector`, `timescaledb` out-of-the-box:

```yaml
all:
  children:
    pg-meta:
      hosts: {10.10.10.10: { pg_seq: 1, pg_role: primary }}
      vars:
        pg_cluster: pg-meta
        pg_databases: {name: meta, extensions: [ postgis, vector ]} # create (in database)
        pg_extensions: [ postgis, pgvector ]                        # install (in cluster)
  vars:
    repo_extra_packages: [ postgis, timescaledb, vector ]           # download  (globally)
```

When you init this PG cluster, these extensions will be made available for you in the `pg-meta` cluster.


Here's a more complicated example: launch Postgres with required extensions for [self-hosting supabase](/docs/app/supabase):

```yaml
all:
  children:
    pg-meta:
      hosts: { 10.10.10.10: { pg_seq: 1, pg_role: primary } }
      vars:
        pg_cluster: pg-meta
        pg_databases:
          - name: postgres
            baseline: supabase.sql
            schemas: [ extensions ,auth ,realtime ,storage ,graphql_public ,supabase_functions ,_analytics ,_realtime ]
            extensions:                                 # Extensions to enable in the postgres database
              - { name: pgcrypto  ,schema: extensions } # Encryption functions
              - { name: pg_net    ,schema: extensions } # Asynchronous HTTP
              - { name: pgjwt     ,schema: extensions } # JSON Web Token API for PostgreSQL
              - { name: uuid-ossp ,schema: extensions } # Generate universally unique identifiers (UUIDs)
              - { name: pgsodium        }               # Modern cryptography for PostgreSQL
              - { name: supabase_vault  }               # Supabase Vault extension
              - { name: pg_graphql      }               # GraphQL support
              - { name: pg_jsonschema   }               # JSON schema validation
              - { name: wrappers        }               # Collection of foreign data wrappers
              - { name: http            }               # Web page retrieval within the database
              - { name: pg_cron         }               # Job scheduler for PostgreSQL
              - { name: timescaledb     }               # Time-series data support
              - { name: pg_tle          }               # Trusted Language Extensions for PostgreSQL
              - { name: vector          }               # Vector similarity search
              - { name: pgmq            }               # Lightweight message queue
        # supabase required extensions for loading
        pg_libs: 'timescaledb, plpgsql, plpgsql_check, pg_cron, pg_net, pg_stat_statements, auto_explain, pg_tle, plan_filter'
        pg_parameters:
          cron.database_name: postgres
          pgsodium.enable_event_trigger: off
  vars:
    pg_version: 17
    repo_extra_packages: [pg17-core ,pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-olap ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ]
    pg_extensions:                  [pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ] #,pg17-olap]
```

All available extensions for PG 17 are downloaded and installed, and required ones are loaded & enabled.
