# Nomad's Backpack 🎒

[Backpack](https://backpack.qm64.io) 🎒 is a packaging system for
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
backpack create nginx
```

**Pack** all the files into one single backpack:
```shell
backpack pack ./nginx-0.1.0/
```

**Customize** a pack default values to configure, enable, adjust the jobs:
```shell
backpack unpack values ./nginx-0.1.0.backpack -f ./values.yaml
```

**Run** your Nomad Jobs with my custom values:
```shell
export NOMAD_ADDR="http://127.0.0.1:4646"
export NOMAD_TOKEN=""
backpack run ./nginx-0.1.0.backpack -v ./values.yaml
```

**Get Help** and learn more for each command
```shell
backpack help
```

Happy Backpacking! 🎒😀 

## Read More

* [How to build Backpack binaries from source code](docs/build.md)
* [How to install Backpack from source code](docs/build.md#installing)
* [How to use backpack with Helm](docs/usage.md)

# Copyright and License

Copyright © 2020 Lorenzo Setale https://setale.me
The full license is available in the file [LICENSE](LICENSE)
