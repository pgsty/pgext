## 用法

来源：

- [官方 README](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/README.md)
- [1.0 版本扩展 SQL](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen--1.0.sql)
- [C 实现](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen.c)

`pg_testgen` 生成随机整数和字母数字文本，既提供标量函数，也提供集合返回函数。它用于快速填充开发和测试表，不适合生成凭据、安全令牌或统计上严谨的样本。

### 核心流程

```sql
CREATE EXTENSION pg_testgen;

CREATE TABLE test_rows (id integer, txt text);

INSERT INTO test_rows
SELECT rows_int(5000, 1, 100),
       rows_text(5000, 20, 30);
```

该示例插入 5,000 行：整数位于闭区间 `[1, 100]`，字符串由 20 到 30 个 ASCII 字母或数字组成。

### 对象索引

- `rand_int()` 返回非负 32 位范围随机整数。
- `rand_int(a, b)` 返回闭区间 `[a, b]` 内的整数。
- `rand_text()`、`rand_text(length)` 和 `rand_text(min_length, max_length)` 返回随机 ASCII 字母数字文本。
- `rows_int(count [, min, max])` 返回 `count` 个随机整数。
- `rows_text(count [, length])` 和 `rows_text(count, min_length, max_length)` 返回 `count` 个随机字符串。

### 运维说明

版本 `1.0` 未声明扩展依赖或预加载要求。实现使用 C 库 `rand()` 和取模运算；它不具备密码学安全性，不保证不同后端独立播种，并且可能有偏差。绝不要用于密码、密钥、要求唯一的标识符或基准测试级分布。

应由调用方校验参数。C 代码不会安全拒绝反向或溢出的整数范围、负文本长度以及非法行数；这类输入可能导致错误或不安全的内存分配行为。生成大数据集时，应分批插入并监控 WAL、表大小、索引和事务时长。
