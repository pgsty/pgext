## 用法

来源：

- [已复核提交的 README](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/README.md)
- [扩展控制文件](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/ruid.control)
- [类型、操作符、索引与辅助函数](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/sql/ruid.sql)
- [仓库 PGXN 元数据](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/META.json)
- [PGXN 发行页面](https://pgxn.org/dist/ruid/)

`ruid` 定义了“可读且实用的标识符”类型，其内部值与 PostgreSQL `uuid` 相同，都是 16 字节，但文本表示使用紧凑的改造版 Base64。典型值为 22 个字符，而标准 UUID 形式为 36 个字符。二进制赋值类型转换无需转换函数即可连接两种类型；扩展还提供默认 B-tree 和哈希运算符类。

### 生成与转换标识符

该扩展依赖 `uuid-ossp` 提供 UUID 生成与命名空间辅助函数。

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION ruid;

SELECT ruid_v4();
SELECT '5b0ba39c-6597-11e2-9d36-001b211595d1'::uuid::ruid;
```

辅助函数包括 `ruid_nil`、`ruid_v1`、`ruid_v1mc`、`ruid_v4`、`ruid_v5`，以及命名空间辅助函数 `ruid_dns`、`ruid_oid`、`ruid_url` 和 `ruid_x500`。

### 注意事项

- 紧凑表示只改变文本交换形式；数据库值仍与 `uuid` 一样占用 16 字节。外部系统必须明确理解这种改造版 Base64 字母表。
- 控制文件报告版本 `0.0.4`，而仓库中的 PGXN `META.json` 仍报告 `0.0.3`。对于当前目录项，应以安装 SQL 和控制文件作为相关 API/版本证据。
- README 列出了 `ruid_v3`，但已复核的安装 SQL 并未创建该函数。未检查已安装扩展前，不要依赖它。
- 安装脚本显式把 `search_path` 设置为 `public`，尽管控制文件称扩展可重定位。不要假定扩展模式隔离，应实际测试非 public 模式安装与迁移。
- 源码和文档停留在 2013 年，也没有现代 PostgreSQL 兼容矩阵。请在确切目标版本上测试转储/恢复、类型转换、索引及客户端解码。
