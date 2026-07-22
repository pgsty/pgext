## 用法

来源：

- [上游文档](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/doc/utinyint.md)
- [SQL 类型与类型转换](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.sql.in.c)
- [C 类型实现](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.c)

`utinyint` 版本 `0.1.1` 定义了一个用于表示无符号值的单字节整数类型。它的主要收益是紧凑的标量存储；根据相邻列布局，PostgreSQL 行对齐可能抵消这部分节省。

### 核心流程

```sql
CREATE EXTENSION utinyint;

CREATE TABLE sensor_value (
  sensor_id bigint,
  channel utinyint
);

INSERT INTO sensor_value VALUES (1, 0), (1, 128), (1, 255);
SELECT channel, channel::integer FROM sensor_value;
```

本次核对的安装 SQL 定义了二进制/文本输入输出，以及 `utinyint` 与 `smallint`、`integer`、`boolean` 之间的类型转换。转为更宽整数是隐式转换，转回时只能用于赋值。布尔 false 映射为零，true 映射为一，任何非零值映射为 true。

### 兼容性注意事项

C 头文件使用无符号字节，转换函数接受到 `UCHAR_MAX` 的值，但上游文字称范围为 `-127` 到 `127`。同一文档还声称支持算术、聚合、B-tree 与 GIN，而这些对象并未出现在本次核对的安装 SQL 中。应把源码/安装脚本作为 API 边界，并在实际打包构建上验证 `0`、`128` 和 `255`。

这个自定义基础类型的磁盘格式与二进制发送格式绑定于其 C 实现。应测试范围错误、类型转换、转储恢复、复制、升级和行大小效果。如果可预期的内置操作符、索引、可移植性与维护性比每个标量节省一个字节更重要，应优先使用 `smallint`。
