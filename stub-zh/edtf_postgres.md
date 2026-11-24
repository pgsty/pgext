## 用法

来源：

- [2.0.0 extension README](https://github.com/monumental-archive/edtf/blob/v2.0.0/crates/edtf-postgres/README.md)
- [2.0.0 control](https://github.com/monumental-archive/edtf/blob/v2.0.0/crates/edtf-postgres/edtf_postgres.control)
- [2.0.0 release](https://github.com/monumental-archive/edtf/releases/tag/v2.0.0)

`edtf_postgres` 2.0.0 按 ISO 8601-2:2019 Annex A 验证和规范化扩展日期时间表达式，计算日期边界并比较时间关系。

### 启用与查询

2.0.0 支持 PostgreSQL 15–19。扩展可信且可重定位，拥有数据库 `CREATE` 权限的角色即可安装，无需预加载。

```sql
CREATE EXTENSION edtf_postgres;
SELECT edtf_valid('1985-04-12');
SELECT edtf_level('1985-04-12/..');
SELECT edtf_canonical('1985-04-12');
SELECT edtf_min('1985-04'), edtf_max('1985-04');
SELECT edtf_relation('1985', '1986');
```

SQL 接口处理文本及派生日期，不引入存储型 EDTF 基础类型。限定符或不确定性重要时应保留原始输入，仅将派生边界用于相应查询语义。

### 2.0 版本边界

1.3.5 是最后支持 PostgreSQL 14 的版本。2.0.0 中，`edtf_min` 与 `edtf_max` 对超出 PostgreSQL 有限日期范围的边界返回负或正无穷，替代原先的 NULL；同时修正零年及更早年份的边界。采用新版本时应复核应用谓词与已经存储的派生值。

带 Y 前缀且达到 19 位或更多位数的年份会被拒绝。这个大版本也改变部分歧义范围的规范化与解释；替换已有规范化结果前，应使用代表性的历史输入重新验证。
