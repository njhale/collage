# collage

A naive OCI composition tool.

Looking to smash container images together in an arbitrary fashion with reckless abandon? Look no further, `collage` lets users select and combine layers from multiple distinct images into a single output image.

## Usage

Combining and squashing layers from `etcd`, `redis`, and `parse-server` images into an image named `quay.io/njhale/trash`:

```bash
collage -s glue quay.io/coreos/etcd:v3.3.12 docker.io/redis:latest registry.redhat.io/bitnami/parse-server -o quay.io/njhale/trash
```

Combining selected layer ranges from `helm-operator` and `ansible-operator` images into an image named `quay.io/njhale/garbage`:

```bash
collage glue quay.io/operator-framework/helm-operator:v0.5.0[2:] quay.io/operator-framework/ansible-operator:v0.6.0[21:27]
```

## Installation

\<install-docs>

## Development

\<dev-docs>

## License

collage is made available under the Apache 2.0 license.
See the [LICENSE](LICENSE) file for details.