## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/README.md)
- [Rust 实现](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/src/lib.rs)
- [Cargo 特性矩阵](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/Cargo.toml)
- [扩展 control 文件](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/pg_ethopic.control)

`pg_ethopic` 可将 PostgreSQL 日期格式化为埃塞俄比亚历日期，并使用埃塞俄比亚数字格式化整数。

```sql
CREATE EXTENSION pg_ethopic;
SELECT ethopic_date(CURRENT_DATE);
SELECT ethopic_date(DATE '2026-01-01', '{year}/{month}/{day}');
SELECT ethopic_number(420);
```

自定义日期模板支持上游 Rust 库实现的字段；应使用领域样例验证模板与预期日历边界。结果是格式化文本，而不是具备日历语义的独立日期类型，因此排序与运算仍应保留原始 PostgreSQL `date`。

0.0.1 是早期 pgrx 扩展。README 提供 PostgreSQL 13–17 容器镜像，Cargo 还定义 PostgreSQL 12 特性。将格式化结果用于导出记录或法律文档前，应测试转换结果、负数与大数、与区域设置无关的输出，以及升级行为。
