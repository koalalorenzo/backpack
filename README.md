# Nomad's Backpack 🎒

[Backpack](https://backpack.qm64.tech) 🎒 is a packaging system for
[Hashicorp Nomad](https://www.nomadproject.io) that allows to:

* Helps you define and install complex jobs configuration
* Helps building reproducible jobs across multiple Nomad clusters
* Simplifies updates to new version of jobs
* Allows you to publish and share packages of applications

Please, keep in mind that this is designed for Nomad, and it might result as
very different than Helm, as Kubernetes is way more than a scheduler.
Read more [here](https://www.nomadproject.io/intro/vs/kubernetes.html) about
the differences between k8s and nomad

To learn more about the motivation behind the development of this project
check [the blog post on Qm64 website](https://qm64.tech/posts/202011-hashicorp-nomad-backpack/).

If you need some help or you want to stay updated with the latest news,
[join Qm64's chatroom on Matrix](https://matrix.to/#/#qm64:matrix.org?via=matrix.org)

Backpack is currently tested against Nomad version 0.12.8

## TL;DR: Install
You can manually download the latest release from 
[the release page here](https://gitlab.com/Qm64/backpack/-/releases).

Or compile the binaries:
```shell
go get -v gitlab.com/Qm64/backpack
cd $GOPATH/src/gitlab.com/Qm64/backpack/
make install
```

## TL;DR How to Use
**Create** your first pack, by using the boilerplate directory structure:

```shell
backpack create nginx
```

**Pack** all the files into one single pack:
```shell
backpack pack ./nginx-0.1.0/
```

**Customize** the values for the template to configure, enable, adjust the jobs:
```shell
backpack unpack values ./nginx-0.1.0.backpack -f ./values.yaml
```

**Plan** and validate (dry-run) the jobs of a package before running:
```shell
backpack plan ./nginx-0.1.0.backpack -v ./values.yaml
```

**Run** your Nomad Jobs with my custom values:
```shell
backpack run ./nginx-0.1.0.backpack -v ./values.yaml
```

**Check** the status of the job allocations:
```shell
backpack status ./nginx-0.1.0.backpack --all
```

Unpack, customize or Run a backpack **from an URL**:
```shell
backpack unpack values https://backpack.qm64.tech/examples/redis-6.0.0.backpack -f ./values.yaml
backpack run https://backpack.qm64.tech/examples/redis-6.0.0.backpack -v values.yaml
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
