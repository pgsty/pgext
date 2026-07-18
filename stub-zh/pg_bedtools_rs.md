## 用法

来源：

- [已复核 commit 的 pg_bedtools_rs README](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/README.md)
- [已复核 commit 的 pg_bedtools_rs.control](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/pg_bedtools_rs.control)
- [Cargo 软件包元数据](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/Cargo.toml)
- [bed_merge 实现](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/src/bed/merge.rs)
- [生成的扩展 SQL](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/sql/pg_bedtools_rs--0.1.0.sql)

`pg_bedtools_rs` 提供 `bed_merge(table_name, padding, conditions)`，使用 Rust bedrs 库合并重叠的基因组区间。输入关系必须有 `chromosome`、`pos_start` 和 `pos_end` 列；可选 padding 会先向两端扩展区间再合并。

### 合并区间

```sql
CREATE EXTENSION pg_bedtools_rs;

CREATE TABLE genomic_interval (
  chromosome varchar NOT NULL,
  pos_start integer NOT NULL,
  pos_end integer NOT NULL
);
INSERT INTO genomic_interval VALUES
  ('chr1', 5, 10), ('chr1', 7, 15), ('chr1', 22, 30);

SELECT * FROM bed_merge('genomic_interval', 0, '1 = 1');
```

### 注意事项

- `table_name` 和 `conditions` 都通过 Rust `format!` 直接插入 SQL，既没有引用为标识符，也没有参数化。只能接受可信代码提供的值，不要向不可信调用者开放该函数。
- 实现会先把全部选中区间加载到后端内存，再返回合并结果。应验证区间边界，并谨慎控制大型任务规模。
- 版本 `0.1.0` 只支持 PostgreSQL 13。上游一份次要文档把函数称为 `bed_overlap`，但已复核 Rust 导出和 README 使用 `bed_merge`。
- 控制文件要求由 `superuser` 安装，并把扩展标为不可信。安装权限并不能让动态提供的过滤文本变安全。
