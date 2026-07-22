## 用法

来源：

- [官方 README](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/README.md)
- [官方 FDW 回调](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/src/parquet_fdw.c)
- [官方扩展控制文件](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/parquet_s3_fdw.control)

`parquet_s3_fdw` 把 PostgreSQL 服务器本地、Amazon S3 或 MinIO 兼容端点上的 Parquet 文件映射为外部表。它支持强类型与无模式扫描、行组裁剪、并行或多文件读取器、导入辅助函数，以及通过重写文件实现的 `INSERT`、`UPDATE` 和 `DELETE`。

### 核心流程

创建服务器，在用户映射中保存 S3 凭据，然后让外部表指向一个或多个文件：

```sql
CREATE EXTENSION parquet_s3_fdw;

CREATE SERVER parquet_s3_srv
  FOREIGN DATA WRAPPER parquet_s3_fdw;

CREATE USER MAPPING FOR app_reader
  SERVER parquet_s3_srv
  OPTIONS (user '<access-key>', password '<secret-key>');

CREATE FOREIGN TABLE user_events (
  user_id bigint OPTIONS (key 'true'),
  event_name text,
  observed_at timestamp
)
SERVER parquet_s3_srv
OPTIONS (
  filename 's3://analytics/events.parquet',
  region 'ap-northeast-1'
);

SELECT * FROM user_events WHERE user_id = 42;
```

使用 MinIO 时，设置服务器选项 `use_minio` 并配置文档中的 `endpoint`。本地路径在数据库服务器上解析，访问身份是 PostgreSQL 操作系统账户。

### 重要选项与功能

- `filename` 选择以空格分隔的文件列表，`dirname` 选择目录；单个外部表不能混用本地和 S3 路径。
- `sorted`、`files_in_order`、`max_open_files`、`use_mmap` 和 `use_threads` 用来选择有序、缓存、内存映射或并行解码策略。
- `files_func` 动态返回文件路径，`insert_file_selector` 为新行选择目标文件；二者都应视为有权限的路由逻辑。
- `schemaless = 'true'` 用一个 `jsonb` 列表示每行。强类型模式映射已记录的 Arrow 基础类型、数组与映射，不支持结构体和嵌套列表。
- `IMPORT FOREIGN SCHEMA`、`import_parquet_s3` 和 `import_parquet_s3_explicit` 可以从文件推导外部表。

### 修改与事务边界

- 上游明确说明不支持事务。不能假定 PostgreSQL 回滚会撤销 Parquet/S3 修改；绝不能把这些写入与要求原子性的业务更新组合在一起。
- 更新和删除会重建缓存数据并覆盖旧 Parquet 文件。大文件性能较差，并发修改同一个文件会产生不一致结果。
- 标记为 `key` 的列用于定位更新或删除的行，但 FDW 不强制其唯一性。必须在外部验证唯一性并避免并发写入。
- 不支持 `WITH CHECK OPTION`、`ON CONFLICT` 和 `RETURNING`。压缩仅限 `zstd`、`snappy` 或不压缩，并受文档所列限制。
- 用户映射中的凭据属于系统目录机密。应限制服务器与映射可见性，使用窄权限对象存储策略和 TLS，轮换密钥，并审计端点与地域设置。
- 应固定所需的 Arrow/Parquet、AWS SDK 版本和编译器 ABI。生产使用前，应测试模式演进、时区、空值、数值精度、映射与列表转换、网络故障、部分覆盖、连接缓存和恢复流程。
