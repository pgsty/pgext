

## 用法

> [babelfishpg_money: SQL Server Money 数据类型](https://babelfishpg.org/)

`babelfishpg_money` 扩展作为 Babelfish 项目的一部分，为 PostgreSQL 提供 SQL Server 兼容的 `MONEY` 和 `SMALLMONEY` 数据类型实现。

### 启用

```sql
CREATE EXTENSION babelfishpg_money;
```

### 数据类型

- **MONEY** - 8 字节货币值，范围从 -922,337,203,685,477.5808 到 922,337,203,685,477.5807，固定 4 位小数
- **SMALLMONEY** - 4 字节货币值，范围从 -214,748.3648 到 214,748.3647，固定 4 位小数

### 行为

该扩展实现了 SQL Server 的货币运算规则：

- 精确 4 位小数的定点表示
- SQL Server 兼容的货币计算舍入行为
- MONEY 与其他数值类型之间的正确转换
- 算术运算遵循 SQL Server 语义（例如 money / money = money，而非 float）

### 注意事项

- Babelfish for PostgreSQL 项目的一部分
- 与提供基础类型基础设施的 `babelfishpg_common` 配合使用
- PostgreSQL 内置的 `money` 类型精度和行为不同；该扩展提供 SQL Server 兼容的变体
