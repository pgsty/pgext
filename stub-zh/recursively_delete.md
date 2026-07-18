## 用法

来源：

- [锁定提交的上游 README](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/README.md)
- [0.1.5 版实现](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/sql/recursively_delete.sql)
- [锁定提交的 PGXN 元数据](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/META.json)
- [锁定提交的扩展控制文件](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/recursively_delete.control)

recursively_delete 0.1.5 是一个破坏性管理函数，它发现外键依赖对象，并构造一条递归数据修改 CTE。它可以处理某些循环依赖和复合主键，并打印 ASCII 损坏预览。上游作者明确要求备份，并说明它不适用于事务工作负载。

### 预览与回滚试运行

默认为预览。应在显式事务中运行真实试验，并在另一次有意识的运行中选择提交前，检查所有受影响表：

```sql
CREATE EXTENSION recursively_delete;

SELECT recursively_delete(
  'public.users'::regclass,
  4402,
  false
);

BEGIN;

SELECT recursively_delete(
  'public.users'::regclass,
  4402,
  true
);

ROLLBACK;
```

返回值只计算显式删除的根行，不包括每个依赖行。文本标量值需要显式 text 转换。复合键数组必须按主键索引顺序列出值，且所有数组元素必须共享一种 PostgreSQL 类型。

### 预览副作用与图限制

预览模式不是只计算 SELECT 计数。它会在 PL/pgSQL 子事务中执行实际 DELETE CTE，然后有意抛出并捕获异常以回滚。它仍可能获取锁、消耗工作量、推进序列、写日志，并调用其外部副作用不受事务管理的 DELETE 触发器。应先在恢复副本上运行预览，并设置锁与语句超时。

使用 ON DELETE SET NULL 或 SET DEFAULT 的关系不会被递归删除。图发现依赖所涉及子表拥有主键。分区、继承表、行级安全、延迟约束、语句和行触发器、生成列、非常规键类型、并发写入、深度或宽度很大的图和引号标识符，都需要显式测试。最终查询可能变得极大，并持有大量行锁和表锁。

函数以调用者权限运行，但默认向 PUBLIC 授予执行权。应撤销它，再只授权给已拥有所需精确 DELETE 权限的专用维护角色。辅助函数和视图使用未限定名，也没有固定 search_path；应安装在受控模式并使用可信搜索路径，以防名称捕获。

上游只报告在 PostgreSQL 10.10 和 13.1 上使用，其 PGXN 元数据将发行标记为不稳定。0.1.5 是预览级应急工具，不是引用操作引擎。应优先使用显式、已审查删除过程和声明式外键操作；如果无法避免使用该函数，必须创建并验证备份、停止并发写入、保留预览，并在之后验证行数和完整性。
