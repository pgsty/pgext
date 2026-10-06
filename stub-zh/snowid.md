## 用法

来源：

- [README.md](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/README.md)
- [snowid/snowid--1.1.sql](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/snowid/snowid--1.1.sql)
- [snowid/snowid.control](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/snowid/snowid.control)

`snowid` 保存 64 位标识符，但显示为 UUID 形态的文本。高 12 位编码表标识，低 52 位编码序列值，并不具有完整的 128 位 UUID 空间。

### 核心用法

```sql
CREATE EXTENSION snowid;
CREATE SEQUENCE demo_snowid_seq MAXVALUE 4503599627370495;
SELECT snowid.gen_snowid(1, 'demo_snowid_seq'::regclass);
```

### 运行边界

SQL 安装 snowid 类型、比较算子、B-tree/hash 支持、类型转换，以及 `snowid` 模式下的序列辅助函数。`snowid.pack` 组合各分量，双参数 `snowid.gen_snowid` 消费指定序列。应在预期范围内保持表标识唯一并防止序列溢出；位掩码不会自动避免冲突。

安装还包含用于单参数生成函数的 DDL 事件触发器，启用前须审查它对现有数据库默认值和序列的影响。安装要求超级用户，上游未声明预加载要求、当前 PostgreSQL 主版本矩阵或许可证。
