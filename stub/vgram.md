## Usage

Sources:

- [Official extension control file](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram.control)
- [Official upstream documentation](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/README.md)

`vgram` — Variable-length grams

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "vgram";
SELECT extversion
FROM pg_extension
WHERE extname = 'vgram';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
