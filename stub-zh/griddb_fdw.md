## 用法

来源：

- [目录源码版本的官方 README](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/README.md)
- [目录源码版本的扩展控制文件](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw.control)
- [1.2 版扩展 SQL](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw--1.2.sql)
- [GridDB C 客户端仓库](https://github.com/griddb/c_client)

griddb_fdw 把 GridDB 容器公开为 PostgreSQL 外部表。它支持读写，以及部分表达式、聚合、排序和限制下推，但依赖 GridDB C 客户端，也不提供 PostgreSQL 本地存储语义。

### 核心流程

构建扩展前先安装匹配的 GridDB 客户端库，然后创建服务器、用户映射和外部表。当远端容器使用行键时，必须明确标识对应的列。

```sql
CREATE EXTENSION griddb_fdw;

CREATE SERVER griddb_svr
FOREIGN DATA WRAPPER griddb_fdw
OPTIONS (
    host '239.0.0.1',
    port '31999',
    clustername 'myCluster',
    database 'public'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER griddb_svr
OPTIONS (username 'admin', password 'admin');

CREATE FOREIGN TABLE sensor_readings (
    ts timestamp OPTIONS (rowkey 'true'),
    sensor_id text,
    value double precision
)
SERVER griddb_svr
OPTIONS (table_name 'sensor_readings');

SELECT sensor_id, avg(value)
FROM sensor_readings
GROUP BY sensor_id;
```

固定列表通知模式应提供成员列表，而不是组播主机和端口设置。服务器级控制包括连接保留、可更新性和批量大小；外部表控制包括远端表名、可更新性、成本估算和批量大小。列映射支持远端名称和行键标志。

### 下推与工具函数

目录中的 1.2 版 SQL 安装了连接检查与清理函数、版本函数，以及反解析 GridDB 专有函数所需的本地签名。这些本地桩在 PostgreSQL 中直接执行时会刻意报错；依赖它们之前，查询必须满足远端下推条件。

受支持的读取可启用部分执行，README 也列出了可下推的函数和聚合。应使用实际部署的 GridDB 与扩展版本检查生成的远端查询，不能假定每个 PostgreSQL 表达式都可下推。

### 限制

固定的上游版本支持查询、插入、更新和删除，并记录了外部模式导入功能。它明确没有实现外部数据包装器的截断接口。批量插入取决于本地 PostgreSQL 主版本；由于 GridDB 不提供生成列，这类列由包装器求值。

连接缓存函数只影响拥有相应缓存连接的会话。网络凭据存放在用户映射中，因此应保护系统目录可见性，并使用最小权限的 GridDB 用户。扩展没有声明预加载要求，但服务器必须能装载扩展库和 GridDB 客户端库。
