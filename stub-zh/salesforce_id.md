## 用法

来源：

- [官方 README](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/README.md)
- [官方扩展 SQL](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/salesforce_id--1.0.sql)
- [官方扩展控制文件](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/salesforce_id.control)

`salesforce_id` 提供一种用于 Salesforce 对象标识符的 12 字节紧凑 PostgreSQL 类型。它接受区分大小写的 15 字符形式和不区分大小写的 18 字符形式，统一输出 18 字符，并提供比较和索引能力，适合需要存储大量 Salesforce ID 的数据库。

### 核心流程

创建扩展，在模式边界使用该类型，并确认两种文本形式会比较为同一个标识符：

```sql
CREATE EXTENSION salesforce_id;

SELECT '0012800000CXn0kAAD'::salesforce_id;
SELECT '0012800000CXn0kAAD'::salesforce_id
     = '0012800000CXn0k'::salesforce_id AS same_id;

CREATE TABLE sf_objects (
  id salesforce_id PRIMARY KEY
);
```

输入必须使用 Salesforce 字符集 `0-9A-Za-z`，且长度严格为 15 或 18 个字符。该类型定义了与 `text` 之间的隐式类型转换，因此应仔细检查重载表达式和应用绑定，不能盲目依赖静默转换。

### 对象索引

- `salesforce_id` 是定长标识符类型。
- `=` 和 `<>` 检查相等与不等。
- `#<#`、`#<=#`、`#>#` 和 `#>=#` 提供顺序比较。
- `salesforce_id_ops` 提供默认 B-tree 与 Hash 运算符类。
- `gen_random_salesforce_id()` 和 `check_salesforce_id_internal()` 是上游 SQL 中的诊断辅助函数，不能代替 Salesforce 签发的标识符。

### 注意事项

上游 README 将该版本描述为实验性且未经生产测试，目录也将其视为已废弃项目。其声称的存储节省幅度有限，而且上游没有证明性能优势。应在外部边界验证规范的 18 字符值，在需要对账时保留原始来源标识符，并在采用自定义类型前，针对实际 PostgreSQL 构建测试转储、恢复、驱动、类型转换、排序和索引。
