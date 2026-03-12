

## Usage

> [permuteseq: scalable pseudo-random permutations of sequences](https://github.com/dverite/permuteseq)

Generate unique, non-sequential, random-looking series of numbers without storing previous values. Uses a Feistel cipher with cycle-walking for format-preserving encryption.

```sql
CREATE EXTENSION permuteseq;
```

### Functions

| Function | Description |
|---|---|
| `permute_nextval(seq_oid, crypt_key bigint)` | Advance sequence and return encrypted value within sequence bounds |
| `reverse_permute(seq_oid, value bigint, crypt_key bigint)` | Compute original value from its permuted element |
| `range_encrypt_element(clear_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | Encrypt a bigint in a given range |
| `range_decrypt_element(crypt_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | Decrypt a previously encrypted value |

### Examples

```sql
CREATE SEQUENCE s MINVALUE -10000 MAXVALUE 15000;

-- Generate random-looking unique values from a sequence
SELECT permute_nextval('s'::regclass, 123456789012345)
  FROM generate_series(1, 10);

-- Reverse a permuted value back to the original
SELECT reverse_permute('s'::regclass, -545, 123456789012345);
-- -10000

-- Encrypt/decrypt within an arbitrary range
SELECT range_encrypt_element(91919191919, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 83028080992

SELECT range_decrypt_element(83028080992, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 91919191919
```
