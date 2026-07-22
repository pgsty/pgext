## 用法

来源：

- [官方用法文档](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/doc/osm_fdw.md)
- [扩展 SQL](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/sql/osm_fdw.sql)
- [建表辅助 SQL](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/sql/utils.sql)

`osm_fdw` 是读取 OpenStreetMap PBF 文件的只读外部数据包装器。它把 OSM 节点、路径与关系暴露为行，便于先用 SQL 转换，再复制到普通 PostgreSQL 表或物化视图中。

### 核心流程

```sql
CREATE EXTENSION osm_fdw;
CREATE SERVER osm_files FOREIGN DATA WRAPPER osm_fdw;

SELECT create_osm_table(
    'malta_osm',
    'osm_files',
    '/srv/osm/malta-latest.osm.pbf'
);

SELECT id, type, lat, lon, tags
FROM malta_osm
WHERE type = 'NODE'
  AND tags ? 'amenity';

CREATE MATERIALIZED VIEW malta_osm_data AS
SELECT * FROM malta_osm WITH DATA;
```

`create_osm_table(table_name, server_name, filename)` 会创建所需的外部表结构。手工定义时必须保持文档规定的列顺序与类型：

```sql
CREATE FOREIGN TABLE osm_data (
    id bigint,
    type text,
    lat double precision,
    lon double precision,
    tags jsonb,
    refs bigint[],
    members jsonb,
    version integer,
    modified timestamp,
    changeset bigint,
    user_id integer,
    username text,
    visible boolean
)
SERVER osm_files
OPTIONS (filename '/srv/osm/region.osm.pbf');
```

### 行模型

`type` 为 `NODE`、`WAY` 或 `RELATION`。节点填充坐标，`refs` 保存路径的节点标识符，`members` 保存关系成员。`tags` 以 `jsonb` 存储 OSM 标签映射；其余列暴露版本与贡献者元数据。

### 注意事项

PostgreSQL 服务器进程必须能读取指定文件路径。FDW 每次查询都会重新读取 PBF 文件，且无法在源文件上建立索引，因此上游建议物化或复制数据后再建立普通 PostgreSQL 索引。应把文件视为外部快照：文件变化时显式刷新派生表。手工创建外部表时既要保持类型，也要保持列位置；大型数据集应专项测试，因为扫描会重新解码源文件。不需要预加载或重启。
