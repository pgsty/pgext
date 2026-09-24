## 用法

来源：

- [Cartridge documentation](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Docs/Book/Cartridge.md)
- [Control](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/rdkit.control)
- [SQL template](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/rdkit.sql.in)
- [Upgrade script](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/update_sql/rdkit--4.7.0--4.8.0.sql.in)

`rdkit` 提供分子类型、子结构搜索、指纹和化学描述符。本文依据 RDKit 2026.03.6 随附的 PostgreSQL cartridge。

### 创建扩展

```sql
CREATE EXTENSION rdkit;
```

该 cartridge 会增加 `mol`、`bfp` 和 `sfp` 等化学类型。

### 核心搜索操作符

cartridge 文档覆盖了以下内容：

- `@>` 和 `<@`：用于子结构匹配。
- `@=`：用于分子精确相等判断。
- `%`、`<%>` 和 `<#>` 这一类指纹相似度与 KNN 操作符：用于相似性搜索。

这些操作通常会和建在指纹列上的 GiST 索引一起使用。

### 指纹与相似度

文档中常见的 SQL 指纹函数包括 `morgan_fp`、`morganbv_fp`、`featmorgan_fp`、`rdkit_fp`、`atompair_fp`、`torsion_fp`、`layered_fp` 和 `maccs_fp`。

cartridge docs 中的示例：

```sql
SELECT tanimoto_sml(
  morganbv_fp('c1ccccc1'::mol),
  morganbv_fp('c1ccccc1O'::mol)
);
```

### 描述符与校验

cartridge docs 还公开了校验与描述符辅助函数，例如：

- `is_valid_smiles()`
- `is_valid_ctab()`
- `is_valid_smarts()`
- `mol_amw()`
- `mol_hba()`
- `mol_numrings()`

这些函数构成了 SQL 层面对分子结构做分析时最主要的用户接口。

### 版本与升级边界

RDKit 工具包发行版 2026.03.6 内含的 SQL 扩展版本是 `4.8.0`，两者使用不同的版本号。上游 `4.7.0` 到 `4.8.0` 的 SQL 模板修改函数代价，但其中 `fmcs_smiles` 语句的行注释吞掉了代价子句和终止符。不要假定未经修补的 `ALTER EXTENSION rdkit UPDATE` 路径可用；应检查实际安装的升级脚本，并先在恢复出的副本上测试。打包补丁与上游源码需要区分。
