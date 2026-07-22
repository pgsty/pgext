## 用法

来源：

- [官方 README](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/README.md)
- [扩展 SQL](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji--1.0.sql)
- [扩展控制文件](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji.control)

`emoji` 是一个纯 SQL 编解码器，可将二进制数据或 UTF-8 文本表示成 emoji 序列。它适合可逆的可视化编码，不应作为加密、认证或可靠纠错方案：其头部只有 9 位校验和。

### 核心流程

安装扩展，然后用相应函数编码和解码：

```sql
CREATE EXTENSION emoji;

SELECT emoji.encode('\x0123456789abcdef'::bytea);
SELECT emoji.decode(emoji.encode('\x0123456789abcdef'::bytea));

SELECT emoji.from_text('Hello, PostgreSQL!');
SELECT emoji.to_text(emoji.from_text('Hello, PostgreSQL!'));
```

`emoji.from_text(text)` 会先把文本转换为 UTF-8 字节，`emoji.to_text(text)` 则把解码后的字节解释为 UTF-8。无效 emoji 输入、映射缺失或校验和不匹配都会使 `emoji.decode(text)` 返回 `NULL`；文本封装函数也会相应返回 `NULL`。

### API

- `emoji.encode(bytea) → text` 使用固定 emoji 字母表编码字节。
- `emoji.decode(text) → bytea` 验证头部并还原字节。
- `emoji.from_text(text) → text` 是 UTF-8 文本编码器。
- `emoji.to_text(text) → text` 是 UTF-8 文本解码器。

第一个输出 emoji 保存一个填充标志和 9 位校验和。上游给出的偶然错误检出概率是 511/512，约 99.8%；这并不是密码学完整性保证。

### 运维说明

版本 `1.0` 会安装固定的 `emoji` 模式，以及基于 Emoji 13.1 前 1,024 个单码点条目的 1,024 行映射。编码数据因此与该表绑定：不要重排或修改 `emoji.chars`，解码旧值时也要保留相同的扩展数据。

SQL 实现调用 `sha512(bytea)`，但没有声明扩展依赖。安装前应确认目标 PostgreSQL 构建提供该函数。核验的上游源码只有 2021 年发布的 `1.0`，因此应在实际服务器和客户端版本上测试兼容性以及 Unicode/终端渲染行为。
