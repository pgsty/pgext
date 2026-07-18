## Usage

Sources:

- [Official extension control file](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/pgsloth.control)
- [Official upstream documentation](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/README.md)
- [Official Rust package manifest](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/Cargo.toml)

`pgsloth` — Intentionally delays every query by a random 100 to 9,999 milliseconds through an executor hook.

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
