## 用法

来源：

- [腾讯云 cos_fdw 指南](https://cloud.tencent.com/document/product/409/73657)
- [腾讯云数据库 PostgreSQL 插件矩阵](https://cloud.tencent.com/document/product/409/75121)

`cos_fdw` 是腾讯云数据库 PostgreSQL 的服务商扩展，它把腾讯云对象存储（COS）中不带表头的 CSV 文件映射为只读外部表。它适合查询冷数据以及把数据迁移到本地表；它不是可移植的社区 FDW，也不是通用对象存储写入器。

### 核心流程

先把数据导出为无表头 CSV 并上传到合适地域的 COS 存储桶，再依次创建依赖、包装器、服务器和外部表。`host` 应填写不带 `http://` 或 `https://` 前缀的存储桶访问域名。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION cos_fdw;

CREATE SERVER cos_server
FOREIGN DATA WRAPPER cos_fdw
OPTIONS (
    host 'example.cos.ap-nanjing.myqcloud.com',
    bucket 'example-bucket',
    id 'COS_SECRET_ID',
    key 'COS_SECRET_KEY'
);

CREATE FOREIGN TABLE archive_sensor_log (
    sensor_log_id integer,
    location text OPTIONS (force_not_null 'true'),
    reading bigint,
    reading_date timestamp
)
SERVER cos_server
OPTIONS (
    filepath '/archive/sensor_log.csv',
    format 'csv',
    null 'NULL'
);

SELECT *
FROM archive_sensor_log
WHERE reading_date >= timestamp '2026-01-01';
```

要把多个对象映射到同一张表，可在 `filepath` 中填写逗号分隔且无额外空格的多个路径。使用 `EXPLAIN` 查看 COS URL、各文件大小和总估算大小。当需要本地索引或事务写入时，可通过 `INSERT INTO local_table SELECT ... FROM archive_sensor_log` 把选定行复制到本地表。

### 选项与边界

- 服务器选项包括 `host`、`bucket`、`id` 和 `key`。
- 表选项包括 `filepath`、`format 'csv'`、`delimiter` 与 `null` 等 CSV 解析选项，以及列级 `force_not_null`。
- 扩展用于读取外部数据；不支持把插入路由到 COS 外部分区。
- 腾讯云数据库使用 `pgcrypto` 加密保存 COS 凭据字段，因此它是必需依赖。即使目录中的存储形式已加密，也必须把 `id` 和 `key` 视为秘密。

### 服务商运维

可用性、支持的 PostgreSQL 版本、权限、升级和凭据处理均由腾讯云数据库控制。执行 `CREATE EXTENSION` 前，应针对实例的准确内核版本检查插件矩阵。

读取 COS 会引入网络延迟、请求费用以及外部一致性边界。应限制创建或修改外部服务器的权限，把 COS 凭据收敛到所需存储桶和路径，让数据库与存储桶位于合适地域，并在生产查询中使用外部表前测试坏行和对象缺失场景。
