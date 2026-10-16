## 用法

来源：

- [README.md](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/README.md)
- [arrow--0.1.sql](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrow--0.1.sql)
- [sql/basic.sql](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/sql/basic.sql)
- [arrowam_handler.c](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrowam_handler.c)
- [LICENSE](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/LICENSE)
- [arrow.control](https://github.com/mkindahl/pg_arrow/blob/5e2ab4e15ea0a1a549b27900c4b95efbcf1a44e6/arrow.control)

`arrow` 0.1 注册基于内存 Arrow 列缓冲区的表访问方法。上游回归示例验证整数和浮点数列的插入与扫描；这是实验性的存储实现。

### 核心用法

```sql
CREATE EXTENSION arrow;
CREATE TABLE arrow_example (a integer, b double precision) USING arrow;
INSERT INTO arrow_example VALUES (1, 2.5);
SELECT * FROM arrow_example;
```

### 运行边界

SQL 对象包括 `arrowam_handler(internal)` 和表访问方法 `arrow`。control 默认要求超级用户安装。实现尚未完成 MVCC 与快照处理，若干索引读取和元组操作直接返回空结果。只应用于可丢弃的实验数据，不能作为持久化、具备事务语义的堆表替代品。上游未声明支持的 PostgreSQL 大版本范围。
