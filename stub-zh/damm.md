## 用法

来源：

- [官方 README](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/README.md)
- [官方 1.0 版 SQL](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm--1.0.sql)
- [官方 control 文件](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm.control)

`damm` 用纯 SQL 实现 Damm 十进制校验位算法。它可生成能检测所有单数字错误与相邻数字换位的人类可见标识符、验证已有编码，或从 PostgreSQL 序列派生带校验位的递增 ID。

### 核心流程

创建序列，并把 `nextdamm` 用作 `damm_code` 域的默认值：

```sql
CREATE EXTENSION damm;
CREATE SEQUENCE thing_id_seq;

CREATE TABLE things (
  id damm_code PRIMARY KEY DEFAULT nextdamm('thing_id_seq'),
  label text NOT NULL
);

INSERT INTO things (label) VALUES ('sample') RETURNING id;

SELECT generate_damm(20140426);
SELECT valid_damm_code(201404265);
```

`generate_damm` 追加一位十进制校验数字；`valid_damm_code` 重新计算该数字并返回布尔值。

### API

- `damm_check_digit(bigint)`：计算校验位但不追加。
- `generate_damm(bigint)`：返回输入乘十后加上校验位的值。
- `valid_damm_code(bigint)`：验证最后一位。
- `damm_code`：由 `valid_damm_code` 约束的 `bigint` 域。
- `nextdamm(varchar)`：对指定序列调用 `nextval`，再返回带校验位的值。

### 注意事项

Damm 校验位用于检测常见抄录错误；它不是密码学校验和、唯一性机制或授权控制。唯一性来自序列，域约束只负责验证十进制表示。

扩展拥有一个查找表，安装后不能动态重定位，不过安装时可选择模式。如果函数和域不在应用搜索路径中，应使用模式限定名。追加一位数字还可能让足够大的输入溢出 `bigint`，因此需要测试取值上界。
