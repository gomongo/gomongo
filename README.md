# gomongo

A library and eventual CLI tool for administrating Mongodb, written in Golang.

## Goals

- Enable declarative, idempotent configuration of the following properties of a MongoDB cluster:
  - [] Replication configuration
  - [] Access control, including:
    - [] Users
    - [] Roles
- Accept a configuration file in YAML (and potentially JSON) format, which delcaratively describes the desired state of the cluster. Example:
  ```yaml
  replicationConfig:
    replicationConfig:
      ID: replicaSetName
      members:
        - ID: 0
          host: mongo1
          priority: 1
          votes: 1
        - ID: 1
          host: mongo2
          priority: 1
          votes: 1
        - ID: 3
          host: mongo3
          priority: 1
          votes: 1
  users:
    - name: user1
      passwordSource:
        type: Vault
        key: secrets/mongo/users/user1
      roles:
        - name: databaseRole
          databaseName: database
  roles:
    - name: databaseRole
      databaseName: database
      privileges:
        - resource: resource
          actions:
            - action
            - action2
  ```

**NOTE: It is the intention to re-use the libraries created as part of this tool to create / enhance a Terraform provider for MongoDB. Depending on the level of effort involved, it is possible that this route may be taken _in place of_ the use of declarative configuration through files, since Terraform already handles a great deal of things like state.**