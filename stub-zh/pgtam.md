## 用法

来源：

- [上游 README](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/README.md)
- [0.0.1 版安装 SQL](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/pgtam--0.0.1.sql)
- [内存表访问方法](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/pgtam.c)
- [上游演示](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/test.sql)
- [MIT 许可证](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/LICENSE.md)

pgtam 0.0.1 是一个用于讲解 PostgreSQL 表访问方法 API 的教学原型。其 mem 访问方法在普通进程内存中存储少量数据并演示回调；它不是存储引擎。

### 一次性演示

应使用内容可以全部丢失的专用单会话测试服务器：

```sql
CREATE EXTENSION pgtam;
CREATE TABLE mem_demo(a integer) USING mem;
INSERT INTO mem_demo VALUES (23), (101);
SELECT * FROM mem_demo;
```

这些行只存在于当前后端的私有内存中，并会在该进程退出时消失。

### 注意事项

- 没有 WAL、磁盘持久化、共享状态、MVCC、崩溃恢复、复制或事务回滚。其他会话看到不同内存，中止事务中的插入也不会撤销。
- 更新、删除、锁定、推测插入、批量插入、索引查找、清理、分析、抽样以及许多其他回调为空，或返回硬编码成功值。
- 存储上限为 100 个表和每表 100 行，而且边界检查不完整。实现只假定非空 32 位整数数据，扫描也只填充第一列。
- 表仅按未限定关系名匹配，因此不同模式中的同名关系会冲突。删除关系不会释放其私有表数据。
- 加载处理器会向固定的 /tmp/pgtam.log 路径追加调试输出，却不检查文件打开失败，造成可靠性与本地文件安全风险。
- TableAmRoutine 结构与主版本绑定，Makefile 还固定了开发者 pg_config 路径。只能针对精确审查过的 PostgreSQL 源码树构建。
