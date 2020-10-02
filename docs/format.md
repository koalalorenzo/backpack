# Backpack File Format
A `.backapck` file consists in metadata, values files and documentation encoded 
with [msgpack](https://msgpack.org). You can _unpack_ a file by running
`backpack unpack` and this will create a new directory containing:

- `backpack.yaml`: Where the metadata, version and dependencies are specified 
- `values.yaml`: Where the default values for the templates are defined
- any file ending with `.nomad`: representing the various Go Templates of Nomad Jobs
- any file ending with `.md`: used for documentation purposes


