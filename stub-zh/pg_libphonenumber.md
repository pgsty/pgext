## 用法

来源：

- [上游 README](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/README.md)
- [0.1.0 版安装 SQL](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/sql/pg_libphonenumber--0.1.0.sql)
- [可由 SQL 调用的 C++ 实现](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/src/pg_libphonenumber.cpp)
- [紧凑表示与比较](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/src/packed_phone_number.h)
- [PGXN 发布元数据](https://pgxn.org/dist/pg_libphonenumber/)

pg_libphonenumber 0.1.0 提供由 Google C++ libphonenumber 支持的八字节 packed_phone_number 类型、显式地区解析器、国际格式文本输出、国家代码访问函数与 B-tree 操作符。

### 使用显式地区解析

```sql
CREATE EXTENSION pg_libphonenumber;
SELECT parse_packed_phone_number('03 7010 1234', 'AU')::text AS normalized;
SELECT phone_number_country_code(
  parse_packed_phone_number('+1 281 901 0011', 'US')
);
```

国家格式输入应使用 parse_packed_phone_number。直接输入 packed_phone_number 文本时固定采用美国地区，对其他国家存在歧义。

### 注意事项

- 上游将扩展标记为 alpha 且仅部分完成。PGXN 只有 2017 年发布的 0.1.0，仓库最后修改于 2020 年。
- 解析过程不会调用 libphonenumber 的有效性检查。语法可解析并不能证明号码已分配、可接通或适合绑定账户。
- 解析与格式化函数被声明为不可变，但结果取决于链接的 libphonenumber 元数据版本。库升级可能改变规范化输出，并破坏存储表达式中的假设。
- 大于与大于等于 SQL 操作符错误地将自身声明为自己的交换操作符。因此规划器重写可能改变谓词含义；修复并重新测试 SQL 前，不要依赖这些操作符或其 B-tree 操作符类。
- 二进制发送与接收直接复制本机八字节对象，没有可移植字节序或语义验证。应避免在不同架构或独立构建版本之间进行二进制传输。
- 排序采用内部紧凑值启发式，源码明确说明结果可能不直观。它不是语言、拨号、地理或数字字符串顺序。
