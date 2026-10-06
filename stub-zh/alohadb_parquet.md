## 用法

来源：

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_parquet/alohadb_parquet--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet--1.0.sql)
- [contrib/alohadb_parquet/alohadb_parquet.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet.c)
- [contrib/alohadb_parquet/parquet_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/parquet_reader.c)
- [contrib/alohadb_parquet/csv_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/csv_reader.c)
- [contrib/alohadb_parquet/json_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/json_reader.c)
- [contrib/alohadb_parquet/alohadb_parquet.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet.control)

`alohadb_parquet` 在 AlohaDB 中把 CSV 与 JSON Lines 读取为 JSONB 行。1.0 版的 Parquet 入口只验证文件并返回元数据，不解码 Parquet 数据行。

### 核心用法

```sql
CREATE EXTENSION alohadb_parquet;
SET alohadb.parquet_allowed_paths = '/srv/import/';
SELECT * FROM read_csv('/srv/import/example.csv', ',', true);
SELECT * FROM read_json('/srv/import/example.jsonl');
```

### 运行边界

`read_csv` 接收服务器本地路径、分隔符和表头标志，`read_json` 读取 JSON Lines。`read_parquet` 检查魔数与文件尾信息并报告元数据；该源码尚未实现完整的 Parquet 解码。

由超级用户安装，不要求预加载。`alohadb.parquet_allowed_paths` 是只能由超级用户设置的目录列表，默认 /tmp。读取使用 PostgreSQL 操作系统账号权限；路径检查基于字符串前缀，不能作为完整隔离边界，应限制函数调用权限及其可访问的服务器文件。处理不可信输入前应核对文件大小、编码及 CSV 解析限制。
