## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/README.md)
- [0.1 版本 SQL 对象](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/pg_string_extension--0.1.sql)
- [C 实现](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/pg_string_extension.c)

`pg_string_extension` 是一个已放弃的教学型 C 扩展，提供 `echo_text(text)`、`is_palindrome(text)`、`sort_chars(text)` 与 `is_anagram(text, text)`。

```sql
CREATE EXTENSION pg_string_extension;
SELECT is_palindrome('level'),
       sort_chars('postgres'),
       is_anagram('silent', 'listen');
```

已复核文档没有定义 Unicode 规范化、区域设置、大小写折叠、空白、标点或排序规则语义。因此结果只能视为字节或实现级演示，不能当作语言学正确的文本处理。复用前应测试多字节 UTF-8、空串、组合字符、非 ASCII 大小写和长输入。

0.1 版本没有当前兼容矩阵、维护声明或安全审计。应优先使用 PostgreSQL 核心字符串函数或持续维护的语言感知库。如果示例函数用于持久生成值、索引或约束，应固定二进制并在升级间验证精确语义；行为变化后必须重建依赖对象。
