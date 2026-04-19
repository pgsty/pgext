## 用法

- 来源：[project README](https://github.com/rdkit/rdkit/blob/master/README.md)，[cartridge docs](https://www.rdkit.org/docs/Cartridge.html)，[2025.03.6 release](https://github.com/rdkit/rdkit/releases/tag/Release_2025.03.6)

RDKit 自带 PostgreSQL cartridge，用于化学信息学场景下的分子存储、检索、指纹和描述符计算。cartridge docs 仍然是主要的上游用法参考；2025.03.6 release note 没有提到 cartridge 相关的用户侧变化。

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
