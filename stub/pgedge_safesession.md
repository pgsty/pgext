## Usage

Sources:

- [Official extension control file](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/pgedge_safesession.control)
- [Official upstream documentation](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/docs/index.md)
- [Official upstream README](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/README.md)

`pgedge_safesession` — Enforce read-only sessions for specified PostgreSQL roles.

The reviewed catalog snapshot records version `1.0-alpha1`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.

The upstream project is associated with `pgEdge`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
