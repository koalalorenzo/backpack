# Nomad's Backpack

[Backpack](https://backpack.qm64.io) is a packaging system for
[Hashicorp Nomad](https://www.nomadproject.io) that allows to:

* Helps you define and install complex jobs configuration
* Helps building reproducible jobs across multiple Nomad clusters
* Simplifies updates to new version of jobs
* Allows you to publish and share packages of applications

Please, keep in mind that this is designed for Nomad, and it might result as
very different than Helm, as Kubernetes is way more than a scheduler.
Read more [here](https://www.nomadproject.io/intro/vs/kubernetes.html) about
the differences between k8s and nomad

## TL;DR: Install

```shell
make install
```

## TL;DR How to Use

**Create** your first backpack, by creating the directory structure:

```shell
backpack create hello-world
```

**Pack** all the files into one single backpack:
```shell
backpack pack ./hello-world
```

**Customize** a pack default values to configure, enable, adjust the jobs:
```shell
backpack unpack values -f ./values.yaml
```

**Run** your Nomad Jobs with my custom values:
```shell
export NOMAD_ADDR="http://127.0.0.1:4646"
export NOMAD_TOKEN=""
backpack run ./hello-world-0.1.0.backpack -v ./values.yaml
```

## Read More

* [How to build Backpack binaries from source code](docs/build.md)
* [How to install Backpack from source code](docs/build.md#installing)
* [Commands available](docs/cli.md)
  * [create](docs/cli/create.md)
  * [pack](docs/cli/pack.md)
  * [unpack](docs/cli/unpack.md)
  * [run](docs/cli/run.md)

# Copyright and License

Copyright © 2020 Lorenzo Setale
The full license is available in the file [LICENSE](LICENSE)
