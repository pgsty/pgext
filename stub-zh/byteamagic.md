## 用法

> 来源: [GitHub 仓库](https://github.com/nmandery/pg_byteamagic), [byteamagic 文档](https://raw.githubusercontent.com/nmandery/pg_byteamagic/master/doc/byteamagic.md)
> 扩展名: `byteamagic`
> CSV 中的包条目是 `pg_byteamagic`；上游扩展名是 `byteamagic`。

`byteamagic` 使用 `libmagic` 分析 PostgreSQL 中的 `bytea` 值，用于识别 MIME 类型和文件类型文本。

```sql
CREATE EXTENSION byteamagic;
SELECT byteamagic_mime(data);
SELECT byteamagic_text(data);
```

### 函数

- `byteamagic_mime(bytea)` 返回 MIME 类型文本，行为与 `file --mime-type` 一致。
- `byteamagic_text(bytea)` 返回人类可读的文件类型描述，行为与 `file` 一致。

### 说明

- 该扩展需要 PostgreSQL 开发头文件和 `libmagic` 开发包。
- 适用于将文件或 blob 以 `bytea` 形式存储，并需要在数据库内进行类型识别的场景。
