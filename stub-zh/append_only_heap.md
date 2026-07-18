## 用法

来源：

- [已复核 commit 的 append_only_heap README](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/README.md)
- [已复核 commit 的 append_only_heap.control](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap.control)
- [版本 1.0 的安装 SQL](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap--1.0.sql)
- [表访问方法实现](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap.c)

`append_only_heap` 是一个复用 PostgreSQL heap 行为的表访问方法，但会给插入增加 `HEAP_INSERT_SKIP_FSM`。因此新元组会避开空闲空间映射中的孔洞，可减少真正追加写入的变长数据工作负载产生的碎片。

### 创建试验表

```sql
CREATE EXTENSION append_only_heap;
CREATE TABLE event_log (
  event_id bigint GENERATED ALWAYS AS IDENTITY,
  payload jsonb NOT NULL
) USING append_only_heap;

INSERT INTO event_log(payload) VALUES ('{"event":"created"}');
SELECT * FROM event_log;
```

### 注意事项

- 名称并不会强制不可变：`UPDATE`、`DELETE` 和 `TRUNCATE` 仍是继承的 heap 操作。追加写入策略需要另外实施。
- 版本 `1.0` 的实现会调用 `mprotect`，并覆盖 PostgreSQL 原本为常量的进程级 `TableAmRoutine` 函数指针。这种不受支持的内部钩子可能加载失败、与其他模块冲突，或在服务器变更后失效。
- 跳过空闲空间复用会在行被更新或删除时加快关系增长。应把它视为实验性存储代码，在一次性数据上度量膨胀和 vacuum 行为，并测试崩溃恢复与升级。
