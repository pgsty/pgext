## 用法

来源：

- [当前上游固定版本 README](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/README.md)
- [固定版本配置实现](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/src/guc.rs)
- [固定版本执行器钩子实现](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/src/hooks.rs)
- [固定版本 Cargo 元数据](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/Cargo.toml)
- [固定版本扩展控制文件](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/pg_no_seqscan.control)

pg_no_seqscan 0.1.1 是开发与 CI 的防护栏。执行器钩子遍历查询计划，在指定数据库、模式或表中查找顺序扫描节点，并抛出错误或发送提示。上游明确说明它不适合生产环境。

### 预加载并试用防护栏

```conf
shared_preload_libraries = 'pg_no_seqscan'
enable_seqscan = off
jit_above_cost = 800000000000
pg_no_seqscan.check_schemas = 'public'
pg_no_seqscan.level = 'Error'
```

```sql
CREATE EXTENSION pg_no_seqscan;

CREATE TABLE no_seqscan_demo AS
SELECT generate_series(1, 10000) AS id;

SELECT * FROM no_seqscan_demo WHERE id = 123;
```

添加合适索引后，该例可采用索引计划并通过检查。0.1.1 修复了并行 EXPLAIN 行为，并声明通过 pgrx 构建 PostgreSQL 14 至 18。

### 防护栏而不是强制策略

小表或大范围查询使用顺序扫描可能完全正确，而索引扫描也仍可能很慢。扩展还会把 enable_seqscan 设为关闭，并建议极高的 JIT 阈值，因此测试计划会有意偏离正常生产规划。错误可能阻断迁移、维护、分析、目录工具和合理的全表读取。

pg_no_seqscan.level 是 USERSET，任意用户都能把它设为 Off；包含文档所述跳过注释的查询也会绕过检查。因此它无法对不受信任用户强制执行安全或性能策略。范围配置仅限超级用户，但表匹配基于名称列表，必须覆盖模式、分区、CTE、子查询、视图和工具命令测试。

错误与提示消息包含查询文本和渲染后的计划，可能把字面量写入客户端输出或服务器日志。应使用合成 CI 数据、限制日志访问，只在专用开发实例预加载，并测量复杂计划上的钩子开销；普通 EXPLAIN 审查和生产可观测性仍是权威流程。
