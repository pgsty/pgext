## 用法

> 来源: [GitHub README](https://raw.githubusercontent.com/rdkit/rdkit/master/README.md), [Cartridge 文档](https://www.rdkit.org/docs/Cartridge.html), [项目网站](https://www.rdkit.org)
> RDKit 提供 PostgreSQL 的化学信息学 cartridge，用于分子存储、检索、指纹和描述符计算。

```sql
CREATE EXTENSION rdkit;
```

该 cartridge 增加了分子类型和运算符，支持子结构搜索、精确匹配、相似性搜索、指纹生成以及大量描述符函数。

### 核心类型

- `mol` 用于分子和查询分子值。
- `bfp` 用于位向量指纹。
- `sfp` 用于稀疏计数指纹。

### 搜索操作符

- `@>` 判断右侧分子是否为左侧分子的子结构。
- `<@` 执行反向子结构判断。
- `@=` 判断分子是否完全相等。
- `#` 是 Dice 相似度阈值操作符。
- `<%>` 支持 Tanimoto KNN 搜索。
- `<#>` 支持 Dice KNN 搜索。

### 指纹

- `morgan_fp` 和 `morganbv_fp` 用于 Morgan 指纹。
- `featmorgan_fp` 和 `featmorganbv_fp` 用于基于特征的 Morgan 指纹。
- `rdkit_fp` 用于哈希化 RDKit 指纹。
- `atompair_fp`、`atompairbv_fp`、`torsion_fp` 和 `torsionbv_fp`。
- `layered_fp` 和 `maccs_fp`。

### 描述符示例

```sql
SELECT is_valid_smiles('c1ccccc1');
SELECT tanimoto_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));
SELECT mol_amw('c1ccccc1'::mol);
SELECT mol_hba('c1ccccc1O'::mol);
SELECT mol_numrings('c1ccccc1'::mol);
```

### 说明

- 文档还描述了更多分子校验辅助函数，例如 `is_valid_smiles`、`is_valid_ctab` 和 `is_valid_smarts`。
- 该 cartridge 还暴露了大量额外的描述符和指纹大小 GUC，但此 stub 仅保留最核心的用户入口。
