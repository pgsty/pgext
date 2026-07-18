## 用法

来源：

- [上游 README](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/README.md)
- [0.0.1 版安装 SQL](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/pg_random_string--0.0.1.sql)
- [C 实现](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/pg_random_string.c)

pg_random_string 0.0.1 版公开 random_string(integer)，尝试返回指定长度的字母数字字符串。该实现只是小型概念验证，按当前发布状态并不安全。

### 仅供演示

只能在审查或修复 C 代码后，于可丢弃、隔离的测试服务器中调用该函数：

```sql
CREATE EXTENSION pg_random_string;
SELECT random_string(16);
```

不得将其输出用于密码、令牌、盐值、需要不可预测性的标识符或任何其他安全用途。

### 注意事项

- 该函数每次调用都用当前时间重新播种进程全局 rand 生成器。同一秒内的调用可能重复，而且该生成器不具备密码学安全性。
- 请求的 integer 长度会被截断为短整数。负数或过大输入都没有验证。
- C 代码只分配与请求字节数相等的空间，却在调用 pstrdup 前没有追加终止零。这可能越界读取，返回损坏数据、泄露相邻内存，或使后端崩溃。
- 分配使用 malloc，却没有对应的 free，且绕过了 PostgreSQL 内存上下文。
- 上游没有发布测试、许可证、发布历史或当前 PostgreSQL 兼容性说明。应用维护良好、经过审查的 SQL 或密码学原语替代它，而不是部署该实现。
