

## 用法

> [babelfishpg_common: SQL Server Transact SQL 数据类型支持](https://babelfishpg.org/)

`babelfishpg_common` 扩展作为 Babelfish 项目的一部分，为 PostgreSQL 提供 SQL Server 兼容的数据类型支持，使 PostgreSQL 能够理解和使用 Microsoft SQL Server 数据类型。

### 启用

```sql
CREATE EXTENSION babelfishpg_common;
```

### SQL Server 数据类型

该扩展添加了以下 SQL Server 兼容数据类型：

- **TINYINT** - 1 字节无符号整数（0 到 255）
- **SMALLMONEY** - 小型货币值
- **MONEY** - 货币值（另见 `babelfishpg_money`）
- **DATETIME** - SQL Server 风格日期时间
- **DATETIME2** - 扩展精度日期时间
- **SMALLDATETIME** - 低精度日期时间
- **DATETIMEOFFSET** - 带时区偏移的日期和时间
- **BIT** - SQL Server 兼容布尔值
- **NCHAR** / **NVARCHAR** - Unicode 字符类型
- **UNIQUEIDENTIFIER** - SQL Server 风格 UUID
- **VARBINARY** - 变长二进制数据
- **IMAGE** - 旧版二进制数据类型
- **SQL_VARIANT** - 通用数据类型容器
- **XML** - SQL Server 兼容 XML 类型
- **SYSNAME** - 系统名称类型（nvarchar(128)）

### 关键特性

- 提供 SQL Server 和 PostgreSQL 类型之间的隐式和显式类型转换
- 支持 SQL Server 风格的排序行为
- 处理 SQL Server 特定的类型强制规则
- 与 `babelfishpg_tsql` 配合实现完整的 T-SQL 兼容性

该扩展通常作为完整 Babelfish for PostgreSQL 安装的一部分部署，是 `babelfishpg_tsql` 的前置依赖。
