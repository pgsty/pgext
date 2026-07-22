## 用法

来源：

- [目录版本对应的官方 README](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/README.md)
- [目录版本对应的 Rust 实现](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/src/lib.rs)
- [目录版本中提交的扩展 SQL](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/sql/pg_bedtools_rs--0.1.0.sql)

`pg_bedtools_rs` 0.1.0 是使用 Rust bedrs 库合并基因组区间的 pgrx 实验项目。它要求输入关系包含染色体、起点和终点列。上游只面向 PostgreSQL 13，而且固定版本仓库中没有可直接使用的生成后扩展 SQL 产物。

### 核心流程

先构建完整的 pgrx 软件包，并确认安装脚本确实声明了函数；然后只对可信且固定的关系和谓词调用它。

```sql
CREATE TABLE intervals (
    chromosome text NOT NULL,
    pos_start integer NOT NULL,
    pos_end integer NOT NULL
);

INSERT INTO intervals VALUES
    ('chr1', 10, 20),
    ('chr1', 18, 30),
    ('chr2', 5, 8);

SELECT *
FROM bed_merge('intervals', 0, '1 = 1');
```

### 函数索引

- `bed_merge(text, integer, text)` 排序并合并区间后，返回染色体、起点和终点列。
- 第二个参数为 padding，默认为零；它会在合并前扩展区间起止位置。
- 第三个参数为 SQL 条件字符串，默认是恒真谓词。

### 安装与安全边界

- 源码要求列名为 chromosome、pos_start 和 pos_end，上游也只报告支持 PostgreSQL 13。
- 固定版本提交的 SQL 文件会递归执行 CREATE EXTENSION，而不是声明 Rust 函数。因此源码快照本身并非可安装产物；应使用匹配的 pgrx 工具链生成并打包，再检查产出的 SQL。
- 表名和条件字符串直接插入 SQL，没有标识符引用或参数绑定。绝不能传入用户可控值；只应通过审核过的包装函数访问固定关系和谓词。
- control 文件将扩展标记为仅超级用户可安装且不可迁移 schema。用于生产前还应验证区间范围和 NULL 处理。
