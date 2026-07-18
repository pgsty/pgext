


## 用法

来源：[README](https://github.com/pmpetit/pglinter/blob/2.0.0/README.md)、[how-to](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/how-to/README.md)、[examples](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/examples/README.md)、[rules](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/rules/README.md)、[2.0.0 release](https://github.com/pmpetit/pglinter/releases/tag/2.0.0)

`pglinter` 会分析 PostgreSQL 数据库中的潜在问题、性能问题和最佳实践违规。当前用户文档通过 `pglinter.get_violations()` 暴露检查结果，该函数返回已启用规则的违规行，可过滤，也可连接到 `pg_identify_object()`。

### 运行检查

```sql
SELECT * FROM pglinter.get_violations();
SELECT * FROM pglinter.get_violations() WHERE rule_code = 'B001';

SELECT
  rule_code,
  (pg_identify_object(classid, objid, objsubid)).type AS object_type,
  (pg_identify_object(classid, objid, objsubid)).schema AS object_schema,
  (pg_identify_object(classid, objid, objsubid)).name AS object_name,
  (pg_identify_object(classid, objid, objsubid)).identity AS object_identity
FROM pglinter.get_violations();
```

### 规则管理

```sql
SELECT pglinter.show_rules();                -- Show all rules and their status
SELECT pglinter.explain_rule('B001');        -- Get rule details and suggested fixes
SELECT pglinter.enable_rule('B001');         -- Enable a specific rule
SELECT pglinter.disable_rule('B001');        -- Disable a specific rule
SELECT pglinter.is_rule_enabled('B001');     -- Check if a rule is enabled
SELECT pglinter.enable_all_rules();
SELECT pglinter.disable_all_rules();
SELECT pglinter.show_rule_queries('B001');   -- Inspect the rule query
SELECT pglinter.list_rules();                -- Return a formatted rule list
```

### 规则导入与导出

```sql
SELECT pglinter.export_rules_to_yaml();                -- Export rules to YAML
SELECT pglinter.import_rules_from_yaml('yaml...');     -- Import rules from YAML
SELECT pglinter.export_rules_to_file('/path/to/rules.yaml');
SELECT pglinter.import_rules_from_file('/path/to/rules.yaml');
SELECT pglinter.export_rulemessages_to_yaml();
SELECT pglinter.import_rule_messages_from_yaml('yaml...');
```

### 规则家族

**基础规则（B 系列）：** B001 表缺少主键、B002 冗余索引、B003 外键缺少索引、B004 未使用的索引、B005 大写名称、B006 未使用的表、B007 跨模式外键、B008 外键类型不匹配、B009 共享触发器函数、B010 使用保留关键字、B011 单个模式存在多个所有者、B012 复合主键超过四列、B013 触发器逐行处理且没有 `WHERE` 子句。

**集群规则（C 系列）：** C002 不安全的 pg_hba.conf 条目、C003 使用 MD5 密码加密。

**模式规则（S 系列）：** S001 未设置默认角色授权、S002 名称带环境前缀或后缀、S003 public 模式未加固、S004 对象归系统角色所有、S005 单个模式存在多个所有者。

### 注意事项

版本 `2.0.0` 移除了较旧的 `check()` 和 `check_rule()` 函数；`get_violations()` 现在是唯一的检查 API。当前构建使用 `pgrx` 0.18.1。

上游 `1.1.2` 版本增加了 `B013`。主 README 相比其他文档和导出函数仍有部分滞后，因此本文使用 `get_violations()`，并省略已移除的 `check()` / `check_rule()` 示例。
