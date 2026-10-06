## 用法

来源：

- [extensions/addressql-postgres/README.md](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/README.md)
- [extensions/addressql-postgres/sql/addressql--0.1.0.sql](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/sql/addressql--0.1.0.sql)
- [LICENSE_POLICY.md](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/LICENSE_POLICY.md)
- [extensions/addressql-postgres/addressql.control](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/addressql.control)

`addressql` 0.1.0 是可执行的 SQL/PLpgSQL 原型，用于国家元数据、邮编格式、地址标准化和匹配。种子数据为合成数据，结果不能证明真实投递覆盖或地址身份。

### 核心用法

```sql
CREATE EXTENSION addressql;
SELECT addressql.address_parse('1 Main Street', 'US');
SELECT addressql.country_address_profile('US');
SELECT addressql.postgis_available();
```

### 运行边界

固定 `addressql` 模式中的函数返回 JSONB，包含来源版本、置信度、警告和适用边界。表保存国家配置、邮政区域、投递区域及合成示例；要得到有意义的查询结果，须先加载预期的数据集。

`addressql.address_parse`、`addressql.address_normalize`、`addressql.country_address_profile` 及邮政辅助函数构成核心工作流。空间函数可选用 `postgis`，没有安装时会返回明确的降级警告。控制文件标记为 trusted，实现没有自己的共享库，也不要求预加载；控制文件中的库名只是未使用的占位值。上游未声明主版本矩阵，不能把原型标准化结果当成可投递性验证或隐私保护。
