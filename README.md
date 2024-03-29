# pg-dba
[![CI](https://github.com/jonstacks/pg-dba/actions/workflows/ci.yml/badge.svg)](https://github.com/jonstacks/pg-dba/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/jonstacks/pg-dba/branch/master/graph/badge.svg)](https://codecov.io/gh/jonstacks/pg-dba)

A semi-automated Postgres DBA for helping teams better manage their DB.

## Testing

To run integration tests locally, you must have docker installed. You can run
the integration tests by running:

```
make compose-up
make integration
```

## Docs

You can run the docs server by going to the `docs` folder and running:

```
make doc-server
```