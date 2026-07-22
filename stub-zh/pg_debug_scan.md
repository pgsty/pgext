## 用法

来源：

- [官方 README](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/README.md)
- [扫描实现](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/src/lib.rs)
- [扩展控制文件](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/pg_debug_scan.control)
- [构建特性矩阵](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/Cargo.toml)

`pg_debug_scan` 使用当前事务快照或调用者提供的快照字段执行原始堆扫描。它是用于检查 MVCC 的教学工具，能够暴露通常不可见或已删除的元组版本；它不是时间旅行、恢复或访问控制机制。

### 核心流程

```sql
CREATE EXTENSION pg_debug_scan;

-- Use the current transaction snapshot.
SELECT * FROM pg_debug_scan('public.temperature', NULL);

-- Supply xmin:xmax:xip1,xip2 explicitly.
SELECT * FROM pg_debug_scan('public.temperature', '775:775:');
```

`pg_debug_scan(table text, snapshot text DEFAULT NULL)` 返回 `xmin bigint`、`xmax bigint` 和 `data text`。自定义快照语法为 `xmin:xmax:xip1,xip2`；每个列出的进行中事务 ID 都必须满足 `xmin <= xip < xmax`。

函数通过每种类型的输出函数转换所有用户列，并将结果序列化为类似 JSON 的文本。所有非空值都是 JSON 字符串，而 SQL 空值表示为字符串 `"NULL"`，不是 JSON null。应据此转换并解释 `data`。

### 安全边界

当前实现通过内部 API 解析并打开关系，没有执行常规 SQL 表权限检查。虽然安装扩展需要超级用户，但函数执行权限在未撤销时默认授予公众。应立即将其限制给专用诊断角色：

```sql
REVOKE EXECUTE ON FUNCTION pg_debug_scan(text, text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_debug_scan(text, text) TO mvcc_diagnostics;
```

否则，获授权者可能检查其通常无权查询的关系，包括仍保存敏感值的废弃元组版本。应把该权限视为低层数据库检查权限，只在受控系统中使用。

### 限制与兼容性

扫描只适用于堆关系，会持有完整的 `AccessShareLock`，并在返回前把所有结果行累积在内存中。大关系可能造成大量 I/O 和后端内存消耗。用户快照值是 32 位事务 ID，而现代 `pg_current_snapshot()` 暴露带 epoch 的 `xid8` 值；该实现没有建模事务回卷和 epoch 信息。

解析器会验证各个 `xip` 值，但不验证整体 `xmin`/`xmax` 关系；格式错误的数值输入可能直接报错。源码使用 PostgreSQL 内部堆与快照 API，因此每个主版本都需要匹配的构建。Cargo 声明 PostgreSQL 11 到 16，但关系名处理代码只显式覆盖 12 到 16；使用前必须核验实际目标构建，尤其是 PostgreSQL 11。
