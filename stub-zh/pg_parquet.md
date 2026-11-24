## 用法

来源：

- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.1/README.md)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.1/pg_parquet.control)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.1/Cargo.toml)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.1/sql/pg_parquet--0.6.0--0.6.1.sql)

`pg_parquet` 0.6.1 通过 PostgreSQL COPY 读写 Parquet，并检查文件元数据。源码目标为 PostgreSQL 14–19，可访问服务器本地文件以及 S3、Azure、GCS 和 HTTP(S) 端点；这些位置及凭据不包含在普通 PostgreSQL 备份中。

### 启用与复制

保留其他预加载项，加入该库后重启。由超级用户创建不可重定位的扩展。服务器进程必须有权访问所选本地路径或远程对象。以下示例由具备相应 COPY 权限的管理员执行：

```ini
shared_preload_libraries = 'pg_parquet'
```

```sql
CREATE EXTENSION pg_parquet;
CREATE TABLE parquet_demo (id bigint, label text);
INSERT INTO parquet_demo VALUES (1, 'first'), (2, 'second');
COPY parquet_demo TO '/tmp/parquet_demo.parquet' WITH (format parquet);
CREATE TABLE parquet_copy (LIKE parquet_demo);
COPY parquet_copy FROM '/tmp/parquet_demo.parquet' WITH (format parquet, match_by name);
SELECT * FROM parquet.schema('/tmp/parquet_demo.parquet');
SELECT * FROM parquet.file_metadata('/tmp/parquet_demo.parquet');
```

可导出查询结果以筛选数据。STDIN、STDOUT 和 PROGRAM 流须显式指定 Parquet 格式。外部输出不随 SQL 事务回滚，失败或回滚的导出可能留下文件或未完成的上传。

### 模式与类型处理

```sql
CREATE TABLE inferred_demo () WITH (load_from = '/tmp/parquet_demo.parquet');
SET pg_parquet.max_parallel_workers = 2;
SELECT * FROM pg_stat_progress_copy;
```

`load_from` 表选项推断列并加载数据。复合类型建立在 `parquet_structs` 中，创建者需要该模式的 CREATE 权限，安装时会向 PUBLIC 授予此权限。删除表不会删除生成类型；以 CASCADE 删除扩展可能删除依赖列。映射类型需要 `crunchy_map`，无法推断的类型应先显式定义目标表。

`match_by name` 同样按名称匹配嵌套复合字段，默认按位置匹配。0.6.0 改进字典、decimal、INT96 及无限日期/时间戳读取，保留 JSON 文本并拒绝含 NUL 的字符串。COPY TO 拒绝多维数组。PostGIS geometry/geography 使用 Parquet 地理空间类型，曲线等不支持的几何类型须先转换或显式导出为二进制。

### 文件、性能与凭据

`parquet.metadata`、`parquet.kv_metadata`、`parquet.column_stats` 和 `parquet.schema` 用于检查文件。导出选项 `bloom_filter` 与 `dictionary` 可选择列；COPY FROM 读取全部行，不使用布隆过滤器筛选行。

`pg_parquet.max_parallel_workers` 默认为零。符合条件的复制可使用并行工作进程；查询导出、临时表、触发器、默认值等场景可能保持串行。导入工作进程负责解码，主进程负责插入。通过 `pg_stat_progress_copy` 观察进度，并按实际文件及表配置资源。

`file_size_bytes` 生成文件目录。目录非空时默认拒绝，须显式指定 `overwrite true`；但无法列举目录时不能确定其内容，仍沿用旧写入行为。覆盖可能删除上次复制留下的文件。同一 PostgreSQL 实例会排斥并发写入者，不同实例仍需外部协调。单文件目标保留原有替换行为，应规划失败导出的清理。

只有角色被 `parquet.user_profile_mapping` 允许时，`pg_parquet.profile_name` 才能选择相应对象存储身份。映射仅由超级用户或 `parquet_profile_manager` 管理，每次复制都会重新检查权限。凭据应置于受保护的服务器文件或环境中。`pg_parquet.allow_credentials_in_parameters` 默认关闭，因为 SQL 中的凭据可能进入日志及活动视图。HTTP 导出在内存中缓冲单文件，上限为 10 MiB；对象存储的并行分段上传要求各异，要求等长分段的存储存在限制。

### 升级

部署匹配服务器主版本的文件，按预加载维护流程操作，并采用发行版提供的 SQL 更新路径。迁移到 0.6.0 前应检查目录导出任务、推断复合类型的依赖、凭据映射及类型转换行为。

0.6.1 明确了空间转换边界：只有 PostGIS 3.7 或更高版本的 `ST_Multi` 才会将 `TIN` 或 `POLYHEDRALSURFACE` 转成 `MULTIPOLYGON`；旧版本原样返回这些输入。没有合适转换时，可用显式二进制列保留不受支持的原始几何对象。
