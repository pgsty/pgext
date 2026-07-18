## 用法

来源：

- [上游 README](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/README.md)
- [0.0.2 版安装 SQL](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/fibonacci--0.0.2.sql)
- [C 实现](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/fibonacci.c)

fibonacci 0.0.2 提供 fibonacci(integer)，这是一个返回 32 位整数的迭代式 C 实现。

### 示例

```sql
CREATE EXTENSION fibonacci;
SELECT n, fibonacci(n) FROM generate_series(0, 10) AS n;
```

如需精确的非负 32 位结果，应将输入限制在 0 到 46 的闭区间内。

### 注意事项

- 负数输入会静默返回零。
- 第 47 个斐波那契数已超出有符号 32 位返回类型。更大的输入会触发有符号溢出并可能返回无效结果。
- 运行时间随输入线性增长。非常大的正整数可能长时间占用一个后端 CPU，因此不要让该函数接收无上限输入。
- 上游没有提供许可证、当前兼容矩阵或实质性测试套件。应将其视为教学示例，而非生产数值实现。
