---
title: 创建
description: 创建与启用 PostgreSQL 扩展
icon: PackageOpen
---


## 快速开始

你可以使用 [**`CREATE EXTENSION`**](https://www.postgresql.org/docs/current/sql-createextension.html) 语句来启用（创建）扩展：

```sql tab="vector"
CREATE EXTENSION vector; -- 无需显式加载
```
```sql tab="timescaledb"
CREATE EXTENSION timescaledb; -- 需显式加载
```

扩展需先[**安装**](/zh/usage/install)，部分扩展在使用前还需[**预加载**](/zh/usage/config)。

部分扩展依赖其他扩展。
此时你可以先安装依赖，
或使用 `CASCADE` 子句一次性安装所有依赖。

```sql
CREATE EXTENSION documentdb CASCADE; -- 创建 documentdb 及其所有依赖扩展
```

你也可以通过 Pigsty 自动化创建扩展。

------

## 配置

扩展（数据库逻辑对象）属于[**PostgreSQL 数据库**](/zh/docs/pgsql/db)的一部分。
在 Pigsty 中，你可以通过 [**`pg_databases`**](/zh/docs/pgsql/param#pg_databases) 参数指定数据库中需创建的扩展。

```yaml
pg_databases:
  - { name: meta ,extensions: [ vector, postgis, timescaledb ] }
```

你也可以用 `object` 格式显式指定扩展细节，如指定安装到特定 schema，或指定版本。以下为完整示例（[自托管 supabase](/zh/docs/app/supabase)）：

```yaml
pg_databases:
  - name: postgres
    baseline: supabase.sql
    schemas: [ extensions ,auth ,realtime ,storage ,graphql_public ,supabase_functions ,_analytics ,_realtime ]
    extensions:                                 # 需在 postgres 数据库启用的扩展
      - { name: pgcrypto  ,schema: extensions } # 加密函数
      - { name: pg_net    ,schema: extensions } # 异步 HTTP
      - { name: pgjwt     ,schema: extensions } # PostgreSQL 的 JSON Web Token API
      - { name: uuid-ossp ,schema: extensions } # 生成 UUID
      - { name: pgsodium        }               # 现代加密库
      - { name: supabase_vault  }               # Supabase Vault 扩展
      - { name: pg_graphql      }               # GraphQL 支持
      - { name: pg_jsonschema   }               # JSON Schema 校验
      - { name: wrappers        }               # FDW 扩展集合
      - { name: http            }               # 数据库内网页抓取
      - { name: pg_cron         }               # PostgreSQL 任务调度器
      - { name: timescaledb     }               # 时序数据扩展
      - { name: pg_tle          }               # 可信语言扩展
      - { name: vector          }               # 向量检索扩展
      - { name: pgmq            }               # 轻量级消息队列
```

--------

## 扩展定义

`extensions` 字段为需在数据库中创建的扩展（名称或对象）列表。
扩展将被创建在 dbsu 搜索路径中的第一个 schema（通常为 `public`）。

此处，数据库对象中的 `extensions` 字段为列表，每个元素可以是：

- 扩展名字符串，如 `vector`
- 或包含如下字段的字典对象：
  - `name`：扩展名称，**必填**，注意可能与[**扩展包名**](/zh/usage/pkg)不同。
  - `schema`：安装 schema，**可选**，默认为当前 dbsu 搜索路径第一个 schema，通常为 `public`。
  - `version`：指定扩展版本，**可选**，默认为最新版本，极少用到。

若数据库尚未创建，则在通过 Pigsty [**创建集群**](/zh/docs/pgsql/admin#create-cluster)或[**创建数据库**](/zh/docs/pgsql/admin#create-database)时会自动创建这些扩展。

对于已有集群/数据库，若基线 schema 非空（如含有 DROP），重新创建数据库可能有风险。
因此建议对现有数据库使用自有 schema 迁移工具管理扩展（如 pgadmin、psql、bytebase、flyway、sqlitch 等）。
但在[配置清单](/zh/docs/config/inventory)中列出扩展有助于资产管理（如集群分叉时自动包含这些扩展）。

------

## 默认扩展

Pigsty 默认会创建部分内置扩展及一个特殊的 [`pg_repack`](/zh/e/pg_repack)。

这些扩展由 [**`pg_default_extensions`**](/zh/docs/pgsql/param#pg_default_extensions) 定义，默认创建在 `template1` 和 `postgres` 数据库中。
新建数据库会从 `template1` 继承这些扩展，无需重复创建。

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
  - { name: pg_repack } # <-- 唯一默认创建的第三方扩展
```

Pigsty 还默认创建了一个额外的 `monitor` schema，由 [`pg_default_schemas`](/zh/docs/pgsql/param#pg_default_schemas) 定义。
用于存放监控相关扩展、表、函数与视图。

Pigsty 默认可用的三款**第三方**扩展：

|           扩展           | 作用                        | 位置                                                                                                   |
|:-----------------------:|-----------------------------|--------------------------------------------------------------------------------------------------------|
| [`pg_repack`](/zh/e/pg_repack) | 在线膨胀治理工具           | [`pg_default_extensions`](/zh/docs/pgsql/param#pg_default_extensions)                                   |
|  [`wal2json`](/zh/e/wal2json)  | JSON 格式变更数据捕获       | [无 DDL 扩展](/zh/list/attr#without-ddl)，[安装](/zh/usage/install) 即可用                              |
|    [`vector`](/zh/e/vector)    | 向量数据类型与索引          | [`pg_databases`](/zh/docs/pgsql/param#pg_databases) 示例中                                              |

[`pg_repack`](/zh/e/pg_repack) 是重要的在线膨胀治理工具。

[**`vector`**](/zh/e/vector) 是 RAG 场景下非常流行的向量扩展，
在 Pigsty 默认安装（`pgsql-main` [别名](/zh/usage/pkg)），并在大多数[配置模板](/zh/docs/config/template)的 `meta` 数据库中创建。

[`wal2json`](/zh/e/wal2json) 是另一款重要的变更数据捕获（CDC）扩展，默认安装，但属于[**无 DDL 扩展**](#extension-without-ddl)，无需显式 `CREATE`。

------

## 无 DDL 扩展

[无 DDL 扩展](/zh/list/attr#without-ddl) 无需执行 `CREATE EXTENSION` 命令即可使用。

PostgreSQL 扩展通常包含三部分：必需的控制文件、可选的 SQL 文件、可选的库文件。
若扩展无 SQL 文件，则无需 `CREATE EXTENSION`。

| 组件         | 说明                                             | 是否必需     |
|--------------|--------------------------------------------------|:------------:|
| 控制文件     | 关键元数据，名称、依赖、schema、版本等           | **必需**     |
| SQL 文件     | SQL DDL 语句、类型、函数等                       | **可选**     |
| 库文件       | 二进制共享库（.so、.dylib、.dll）                | **可选**     |

由于 SQL / LIB 文件均为可选，扩展类型有四种组合：

| **[LOAD](/zh/usage/config) / DDL** | 需 `CREATE EXTENSION` | 无需 `CREATE EXTENSION` |
|------------------------------------|----------------------|------------------------|
| **需 `LOAD`**                      | 使用钩子的扩展        | 无头扩展                |
| **无需 `LOAD`**                    | 非钩子扩展            | 逻辑解码输出插件        |

