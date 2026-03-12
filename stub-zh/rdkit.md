


## 用法

> [rdkit: 化学信息学与分子工具 PostgreSQL 扩展](https://github.com/rdkit/rdkit)

RDKit PostgreSQL 扩展提供了用于分子的 `mol` 数据类型、用于指纹的 `fp` 数据类型、子结构和相似性搜索操作，以及 GiST 索引支持。

```sql
CREATE EXTENSION rdkit;
```

### 数据类型

| 类型 | 描述 |
|------|------|
| `mol` | 分子结构（来自 SMILES、SMARTS 等） |
| `bfp` | 位向量指纹 |
| `sfp` | 稀疏（计数）指纹 |

### 分子输入/输出

```sql
-- 从 SMILES 创建分子
SELECT 'c1ccccc1'::mol;

-- 检查 SMILES 是否有效
SELECT is_valid_smiles('c1ccccc1');

-- 将分子转换为 SMILES
SELECT mol_to_smiles('c1ccccc1'::mol);
```

### 子结构搜索

```sql
-- 子结构匹配运算符
SELECT 'c1ccccc1O'::mol @> 'c1ccccc1'::mol;   -- true（苯酚包含苯环）
SELECT 'c1ccccc1'::mol <@ 'c1ccccc1O'::mol;    -- true

-- 使用 SMARTS 模式
SELECT 'c1ccccc1O'::mol @> 'c1ccc(O)cc1'::mol;
```

### 相似性搜索

```sql
-- Tanimoto 相似度（返回 0 到 1 之间的值）
SELECT tanimoto_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));

-- Dice 相似度
SELECT dice_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));
```

### 指纹函数

```sql
-- Morgan 指纹（位向量）
SELECT morganbv_fp('c1ccccc1'::mol);
SELECT morganbv_fp('c1ccccc1'::mol, 2);  -- radius=2

-- RDKit 指纹
SELECT rdkit_fp('c1ccccc1'::mol);

-- 拓扑扭转指纹
SELECT torsionbv_fp('c1ccccc1'::mol);

-- 原子对指纹
SELECT atompairbv_fp('c1ccccc1'::mol);
```

### 描述符计算

```sql
SELECT mol_amw('c1ccccc1'::mol);          -- 平均分子量
SELECT mol_logp('c1ccccc1'::mol);         -- LogP
SELECT mol_hba('c1ccccc1O'::mol);         -- 氢键受体
SELECT mol_hbd('c1ccccc1O'::mol);         -- 氢键供体
SELECT mol_numrotatablebonds('c1ccccc1'::mol); -- 可旋转键
SELECT mol_numatoms('c1ccccc1'::mol);     -- 原子数
SELECT mol_numheavyatoms('c1ccccc1'::mol);    -- 重原子数
SELECT mol_numrings('c1ccccc1'::mol);     -- 环数
```

### GiST 索引支持

创建索引以加速子结构和相似性搜索：

```sql
-- 子结构搜索索引
CREATE INDEX idx_mol ON molecules USING gist(molecule);

-- 指纹相似性索引
CREATE INDEX idx_fp ON molecules USING gist(morganbv_fp(molecule));
```

使用索引查询：

```sql
-- 子结构搜索
SELECT * FROM molecules WHERE molecule @> 'c1ccccc1'::mol;

-- 相似性搜索（带阈值）
SET rdkit.dice_threshold = 0.5;
SELECT * FROM molecules WHERE morganbv_fp(molecule) % morganbv_fp('c1ccccc1O'::mol);
```

### GUC 参数

| 参数 | 默认值 | 描述 |
|------|--------|------|
| `rdkit.tanimoto_threshold` | 0.5 | Tanimoto 相似度运算符 `<%>` 的阈值 |
| `rdkit.dice_threshold` | 0.5 | Dice 相似度运算符 `%` 的阈值 |
