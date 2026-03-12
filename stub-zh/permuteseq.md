

## 用法

> [permuteseq: 可扩展的序列伪随机置换](https://github.com/dverite/permuteseq)

无需存储历史值，即可生成唯一、非连续、看似随机的数字序列。使用 Feistel 密码与循环行走法实现格式保持加密。

```sql
CREATE EXTENSION permuteseq;
```

### 函数

| 函数 | 描述 |
|---|---|
| `permute_nextval(seq_oid, crypt_key bigint)` | 推进序列并返回序列范围内的加密值 |
| `reverse_permute(seq_oid, value bigint, crypt_key bigint)` | 从置换后的元素计算出原始值 |
| `range_encrypt_element(clear_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | 在给定范围内加密一个 bigint |
| `range_decrypt_element(crypt_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | 解密先前加密的值 |

### 示例

```sql
CREATE SEQUENCE s MINVALUE -10000 MAXVALUE 15000;

-- 从序列中生成看似随机的唯一值
SELECT permute_nextval('s'::regclass, 123456789012345)
  FROM generate_series(1, 10);

-- 将置换后的值反向还原为原始值
SELECT reverse_permute('s'::regclass, -545, 123456789012345);
-- -10000

-- 在任意范围内加密/解密
SELECT range_encrypt_element(91919191919, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 83028080992

SELECT range_decrypt_element(83028080992, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 91919191919
```
