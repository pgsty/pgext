## 用法

来源：

- [官方 README](https://github.com/byucesoy/pg_color/blob/master/README.md)
- [类型定义 SQL](https://github.com/byucesoy/pg_color/blob/master/pg_color.sql)
- [1.1 版本更新 SQL](https://github.com/byucesoy/pg_color/blob/master/pg_color--1.0--1.1.sql)

`pg_color` 1.1 是一个教学型扩展，提供紧凑的 RGB 颜色基础类型。文本输入必须恰好是六位大写十六进制字符，且不能带井号前缀；输出采用相同格式。1.1 版只增加了相等和不等比较。

### 核心流程

```sql
CREATE EXTENSION pg_color;

CREATE TABLE swatch (
    name text PRIMARY KEY,
    rgb color NOT NULL
);

INSERT INTO swatch VALUES
    ('salmon', 'FA8072'),
    ('navy',   '000080');

SELECT name, rgb
FROM swatch
WHERE rgb = 'FA8072'::color;
```

### 对象索引

- `color` 使用整数大小的内部表示保存 RGB 值，接受 000000 到 FFFFFF。
- `color = color` 和 `color <> color` 是 1.1 版仅有的两个新增操作符。
- 带井号前缀或包含小写十六进制字符的输入会被解析器拒绝。

### 注意事项

- 扩展没有定义排序操作符、类型转换或 B-tree/hash 操作符类，因此不要假定普通索引能支持等值条件。
- 它只校验表示形式，不处理色彩空间、透明通道、色彩配置或无障碍问题。
- 上游将项目定位为 PostgreSQL 扩展开发示例；将其写入持久 schema 前应在目标服务器大版本上测试兼容性。
