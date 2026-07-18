## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/README.md)
- [1.0 版本 SQL 对象](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/pgstrtranslate--1.0.sql)
- [C 实现](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/pgstrtranslate.c)

`pgstrtranslate` 把 PostgreSQL 按字符工作的 `translate` 操作扩展为有序的多字符替换。函数 `pgstrtranslate(boolean, text, text[], text[])` 返回文本。

```sql
CREATE EXTENSION pgstrtranslate;
SELECT pgstrtranslate(
  false,
  'abcdefghijkl',
  ARRAY['ab', 'efg', '2cd'],
  ARRAY['012', '3', '78']
);
```

第一个参数为 false 时，所有搜索项都针对原始输入匹配，因此替换结果不会产生新匹配；为 true 时，替换按顺序执行，后续搜索项可以匹配前面生成的文本。所以顺序会改变结果。

应确保搜索与替换数组按预期配对，并测试空元素、重复、重叠、空值、多字节文本和膨胀。函数被标记为 immutable，因此持久索引或生成值会假定逐字节语义稳定。1.0 版本没有当前兼容矩阵；用于规范化、安全过滤或标识符前，应压测最坏重叠模式并验证输出。
