


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

**Base (B-series)：** B001 tables without PK、B002 redundant indexes、B003 missing FK indexes、B004 unused indexes、B005 uppercase names、B006 unused tables、B007 cross-schema FKs、B008 FK type mismatches、B009 shared trigger functions、B010 reserved keywords、B011 multiple owners per schema、B012 composite primary keys with more than four columns、B013 row-by-row trigger processing without a `WHERE` clause。

**Cluster (C-series)：** C002 insecure pg_hba.conf entries、C003 MD5 password encryption。

**Schema (S-series)：** S001 no default role grants、S002 env prefixes/suffixes、S003 unsecured public schema、S004 system role ownership、S005 multiple owners per schema。

### 注意事项

版本 `2.0.0` 移除了较旧的 `check()` 和 `check_rule()` 函数；`get_violations()` 现在是唯一 check API。当前构建使用 `pgrx` 0.18.1。

上游 `1.1.2` release 增加了 `B013`。主 README 相比 docs 和导出函数仍有部分滞后，因此此 stub 使用 `get_violations()`，并省略已移除的 `check()` / `check_rule()` 示例。
