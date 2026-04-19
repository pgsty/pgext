
## 用法

> 来源：[README](https://github.com/pmpetit/pglinter/blob/main/README.md), [1.1.2 release](https://github.com/pmpetit/pglinter/releases/tag/1.1.2)

`pglinter` 会分析数据库中的潜在问题、性能隐患和最佳实践违规项，并以 SARIF 2.1.0 格式输出结果。

### 运行检查

```sql
SELECT pglinter.check();                                -- Run all enabled rules
SELECT pglinter.check_rule('B001');                     -- Run a specific rule
SELECT pglinter.check('/path/to/results.sarif');        -- Save SARIF report to file
SELECT pglinter.check_rule('B001', '/path/to/b001.sarif');
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
```

### 规则配置

```sql
SELECT pglinter.update_rule_levels('B001', 30, 70);   -- Set warning/error thresholds
SELECT pglinter.get_rule_levels('B001');               -- Get current thresholds
SELECT pglinter.export_rules_to_yaml();                -- Export rules to YAML
SELECT pglinter.import_rules_from_yaml('yaml...');     -- Import rules from YAML
```

### 规则家族

**Base (B-series):** B001 tables without PK, B002 redundant indexes, B003 missing FK indexes, B004 unused indexes, B005 uppercase names, B006 unused tables, B007 cross-schema FKs, B008 FK type mismatches, B009 shared trigger functions, B010 reserved keywords, B011 multiple owners per schema.

**Cluster (C-series):** C002 insecure pg_hba.conf entries, C003 MD5 password encryption.

**Schema (S-series):** S001 no default role grants, S002 env prefixes/suffixes, S003 unsecured public schema, S004 system role ownership, S005 multiple owners per schema.

### 版本增量

`1.1.2` release 新增了规则 `B013`，用于检测缺少 `WHERE` 子句的逐行处理触发器。README 还没有更新到这条规则，因此这里应以 release note 作为相对 `1.1.1` 的权威增量说明。
