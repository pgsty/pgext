## 用法

来源：

- [官方扩展 README](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/README.md)
- [扩展控制文件](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/edtf_postgres.control)
- [pgrx 清单](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/Cargo.toml)

`edtf_postgres` 为 ISO 8601-2:2019 Annex A 字符串提供 Extended Date/Time Format 校验、规范化、边界计算与时间关系判断。

### 启用

1.2.3 版本为 amd64 与 arm64 上的 PostgreSQL 14–18 发布构件。安装匹配构件后执行：

```sql
CREATE EXTENSION edtf_postgres;
```

扩展可迁移且为 trusted，因此拥有数据库 `CREATE` 权限的角色可以安装。它不要求预加载。

### 校验与规范形式

接收外部文本前使用 `edtf_valid`，用 `edtf_level` 查看支持的 EDTF 级别，并用 `edtf_canonical` 统一等价表示。

```sql
SELECT edtf_valid('1985-04-12');
SELECT edtf_level('1985-04-12/..');
SELECT edtf_canonical('1985-04-12');
```

`edtf_min` 与 `edtf_max` 返回便于索引的日期边界。`edtf_relation` 比较两个表达式并返回三值时间关系。

```sql
SELECT edtf_min('1985-04'), edtf_max('1985-04');
SELECT edtf_relation('1985', '1986');
```

### 数据建模边界

扩展处理文本与派生日期；它不会代替应用校验，也不会引入持久化 EDTF 基础类型。若不确定性与限定符很重要，应保留原始表达式；只有在查询语义明确时才为派生边界建立索引。

