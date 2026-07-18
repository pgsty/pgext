## Usage

Sources:

- [Official extension control file](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser.control)
- [Official upstream documentation](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/README)

`twitter_parser` — Minimal C full-text-search parser that tokenizes Twitter-style hashtags and mentions.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "twitter_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'twitter_parser';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
