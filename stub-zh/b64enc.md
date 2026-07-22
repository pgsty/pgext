## 用法

来源：

- [官方类型定义](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/sql/b64enc.sql)
- [官方 Rust 实现](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/src/lib.rs)
- [官方扩展控制文件](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/b64enc.control)

`b64enc` 是一个实验性的定长基础类型，它把 URL 安全的 Base64 字符串解码为最多八个字节，并以一个按值传递的 64 位 datum 存储。它不是 PostgreSQL `bytea` 加 `encode`、`decode` 函数的通用替代品。

### 核心流程

创建扩展，并把 URL 安全的 Base64 字面量转换成该类型：

```sql
CREATE EXTENSION b64enc;

SELECT 'AQ=='::b64enc;
SELECT 'AQIDBAUGBwg='::b64enc;
```

输出函数总是序列化全部八个已存储字节，因此较短输入会以带前导零字节的规范八字节形式输出。当前审阅的 SQL 只定义了文本输入和输出；二进制发送与接收函数被注释掉了。

### 限制与故障行为

- 解码后超过八字节的输入会被拒绝。没有已记录的比较运算符、类型转换、算术或运算符类，因此除存储和显示外，该类型几乎没有 SQL 行为。
- 如果不同文本编码解码后的字节序列只相差前导零，它们会折叠成相同的 64 位值并产生相同输出。不要把原始 Base64 拼写当作身份属性。
- 输入实现对 Base64 解码结果调用 Rust `unwrap()`。因此非法 ASCII 可能走 Rust panic 路径，而不是受控的 PostgreSQL 输入错误；接受不可信输入前，应在一次性服务器上模糊测试非法和超长值。
- 版本 `0.0.1` 使用旧 Rust crate 与手写 C/Rust FFI 边界。应测试后端进程存活、转储恢复、体系结构与字节序行为、升级以及准确的目标 PostgreSQL ABI。
- 普通二进制数据应优先使用 `bytea` 和 PostgreSQL 维护的 Base64 函数，它们支持任意长度并具有成熟的错误语义。
