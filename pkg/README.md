# Backpack Format

This Go Package contains all the logic related to Backpack files, format, type
and various encoding used.

TL;DR: In a Backpack "unpacked" you can find these files:

* `backpack.yaml`: defining metadata and default values
* `*.nomad`: [Go Template](https://golang.org/pkg/text/template/) of HCL code for [Nomad Jobs](https://www.nomadproject.io/docs/job-specification)

TL;DR: The Backpack "packed" is formatted in [msgpack](https://msgpack.org)
and contains the metadata, the default values, and all the templates encoded in 
base64.