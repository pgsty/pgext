## 用法

来源：

- [扩展控制文件](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/extension%2Fpgchem_tigress.control/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)
- [版本化安装 SQL](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/extension%2Fpgchem_tigress--3.2.sql/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)
- [官方用户指南源码](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/doc%2Fuserguide%2Fpgchem_usersguide%2Fpgchem_usersguide.tex/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)

`pgchem_tigress` 是 legacy chemoinformatics 扩展，提供 `molecule` 与 fingerprint type、化学性质函数、exact/substructure operator 以及 GiST operator class。

### 启用

安装 3.2 库及其 OpenBabel、Barsoi 与 InChI 运行时依赖后，以超级用户创建可迁移扩展：

```sql
CREATE EXTENSION pgchem_tigress;
SELECT pgchem_version();
```

它没有声明预加载。上游已停止活动，源码时代的安装说明面向 PostgreSQL 8 以上版本并使用硬编码辅助数据路径；现代大版本、库 ABI、dump/restore 与升级行为都需要专门兼容性测试。

### 保存并索引分子

`molecule` 输入函数接受 SMILES、InChI 与 MDL molfile text 等格式。GiST 索引加速 equality 与 substructure containment。

```sql
CREATE TABLE compounds (
  id  bigint PRIMARY KEY,
  mol molecule NOT NULL
);

CREATE INDEX compounds_mol_gist
ON compounds USING gist (mol);

INSERT INTO compounds VALUES
  (1, 'c1ccccc1'::molecule);
```

### 搜索操作符

`=` 检查分子完全相等。`A <= B` 判断 A 是否为 B 的子结构，`A >= B` 判断 A 是否包含 B。`@` 返回 Tanimoto coefficient，且不受 GiST 加速。

```sql
SELECT id
FROM compounds
WHERE 'c1ccccc1'::molecule <= mol;

SELECT id,
       'c1ccccc1'::molecule @ mol AS similarity
FROM compounds
ORDER BY similarity DESC;
```

历史版本之间的 molecule serialization 与 fingerprint 曾发生变化。切换二进制时应按准确 release note 迁移存储值并重建相关索引；绝不能在已有 `molecule` 列下直接替换 shared library 而不做显式迁移测试。
