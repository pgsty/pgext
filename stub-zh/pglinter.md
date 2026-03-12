

## 用法

> [pglinter: PostgreSQL 检查与分析扩展](https://github.com/pmpetit/pglinter)

pglinter 分析你的数据库以发现潜在问题、性能隐患和最佳实践违规。结果以 SARIF 2.1.0 格式输出。

### 运行检查

```sql
SELECT pglinter.check();                                -- 运行所有已启用的规则
SELECT pglinter.check_rule('B001');                     -- 运行特定规则
SELECT pglinter.check('/path/to/results.sarif');        -- 将 SARIF 报告保存到文件
SELECT pglinter.check_rule('B001', '/path/to/b001.sarif');
```

### 规则管理

```sql
SELECT pglinter.show_rules();                -- 显示所有规则及其状态
SELECT pglinter.explain_rule('B001');        -- 获取规则详情和修复建议
SELECT pglinter.enable_rule('B001');         -- 启用特定规则
SELECT pglinter.disable_rule('B001');        -- 禁用特定规则
SELECT pglinter.is_rule_enabled('B001');     -- 检查规则是否已启用
SELECT pglinter.enable_all_rules();
SELECT pglinter.disable_all_rules();
```

### 规则配置

```sql
SELECT pglinter.update_rule_levels('B001', 30, 70);   -- 设置警告/错误阈值
SELECT pglinter.get_rule_levels('B001');               -- 获取当前阈值
SELECT pglinter.export_rules_to_yaml();                -- 将规则导出为 YAML
SELECT pglinter.import_rules_from_yaml('yaml...');     -- 从 YAML 导入规则
```

### 可用规则

**基础（B 系列）：** B001 无主键的表、B002 冗余索引、B003 缺失的外键索引、B004 未使用的索引、B005 大写命名、B006 未使用的表、B007 跨模式外键、B008 外键类型不匹配、B009 共享触发器函数、B010 保留关键字、B011 每个模式多个所有者。

**集群（C 系列）：** C002 不安全的 pg_hba.conf 条目、C003 MD5 密码加密。

**模式（S 系列）：** S001 无默认角色授权、S002 环境前缀/后缀、S003 不安全的 public 模式、S004 系统角色所有权、S005 每个模式多个所有者。
