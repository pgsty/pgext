## 用法

来源：

- [阿里云扩展支持矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_address_standardizer_data_us` 版本 `7.0` 是阿里云 ApsaraDB RDS for PostgreSQL 提供的、基于 PAGC 标准的美国地址标准化数据插件。它是服务商提供的数据组件，并不是一组有独立文档的 SQL 函数。

### 启用方式

首先确认精确的 ApsaraDB RDS PostgreSQL 引擎版本与版本系列支持该扩展及版本。在实例允许用 SQL 安装时，服务商文档给出了标准扩展机制：

```sql
CREATE EXTENSION ganos_address_standardizer_data_us;
```

部分托管扩展必须改用服务控制台启用。应遵循当前实例的服务商流程，并在 `pg_extension` 中确认安装结果。

### 运维边界

服务商矩阵分别列出了 `ganos_address_standardizer` 与 `ganos_address_standardizer_data_us`。地址处理操作应通过另行受支持的标准化组件完成；服务商页面没有为该数据插件本身记录独立的函数接口。

可用性会随引擎版本与版本系列变化。如果扩展不存在，阿里云建议升级引擎小版本或核对适用的支持列。不要把此流程用于自建 PostgreSQL：这里引用的制品与启用契约仅适用于阿里云。
