## 用法

来源：

- [pg_bio.control](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/pg_bio.control)
- [README.md](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/README.md)
- [Cargo.toml](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/Cargo.toml)
- [src/lib.rs](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/src/lib.rs)

`pg_bio` 提供实验性的 Rust SQL 函数，涉及分子坐标、序列比较、指纹和生物数据导入。控制文件使用 Cargo 版本 0.0.0，属于早期源码快照。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_bio;
SELECT distance_angstroms(
  create_residue_coord(0, 0, 0, 'A'),
  create_residue_coord(3, 4, 0, 'B')
);
SELECT sequence_alignment_score('ACGT', 'ACGA');
```

### 运行边界

控制文件要求 `vector` 和超级用户安装，没有预加载要求。坐标函数包括 `create_residue_coord`、`distance_angstroms` 和 `z_order_encode`，后者会将坐标截断到源码规定的范围。另有序列比对、SMILES／指纹及 UniProt 获取接口。网络与文件导入会产生服务器端副作用，需限制权限。带模型名称的嵌入函数不能证明执行了训练模型推理，不能据此推导临床、预测或性能保证。缺乏充分源码证据的 PostgreSQL 主版本兼容性与许可信息保持未指定。
