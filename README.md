# Masa

Masa is inspired by the [bucket storage](https://github.com/gocardless/bucket-store) gem by Gocardless 

An abstraction layer on the top of file cloud storage systems such as Google Cloud
Storage or S3. This module exposes a generic interface that allows interoperability
between different storage options. Callers don't need to worry about the specifics
of where and how a file is stored and retrieved as long as the given key is valid.

Keys within `Masa` are URI strings that can universally locate an object
in the given provider. A valid key example would be
`gcs://a-gcs-bucket/file/path.json`.

## Usage
This library is distributed as a Go package, and we recommend adding it to your Go mod:

```go
go get "github.com/ogwurujohnson/masa"
```

## Design and Architecture
The main principle behind `Masa` is that each resource or group of resources must
be unequivocally identifiable by a URI. The URI is always composed of three parts:

- the "adapter" used to fetch the resource (see "adapters" below)
- the "bucket" where the resource lives
- the path to the resource(s)

As an example, all the following are valid URIs:

- `gcs://gcs-bucket/path/to/file.xml`
- `s3://bucket/separator/file.xml`

Even though `Masa`'s main goal is to be an abstraction layer on top of systems such
as S3 or Google Cloud Storage where the "path" to a resource is in practice a unique
identifier as a whole (i.e. the `/` is not a directory separator but rather part of the
key's name), we assume that clients will actually want some sort of hierarchical
separation of resources and assume that such separation is achieved by defining each
part of the hierarchy via `/`.

This means that the following are also valid URIs in `Masa` but they refer to
all the resources under that specific hierarchy:

- `gcs://gcs-bucket/path/subpath/`
- `s3://bucket/separator/`


## Adapters

`Masa` comes with 2 built-in adapters:

- `gcs`: the Google Cloud Storage adapter
- `s3`: the S3 adapter

### GCS adapter
This is the adapter for Google Cloud Storage. `Masa` assumes that the  authorization
for accessing the resources has been set up outside of the gem.

### S3 adapter
This is the adapter for S3. `Masa` assumes that the authorization for accessing
the resources has been set up outside of the gem (see also
https://docs.aws.amazon.com/sdk-for-ruby/v3/api/index.html#Configuration).


## Examples

### Uploading a file to a bucket
```go
bucket.For("gcs://bucket/path/file.xml").Upload(context.Background(), "hello world")
=> "gcs://bucket/path/file.xml"
```

### Accessing a file in a bucket
```go
bucket.For("gcs://bucket/path/file.xml").Download(context.Background())
=> {bucket: "bucket", key: "path/file.xml", content: "hello world"}
```

### Listing all keys under a prefix
```go
bucket.For("gcs://bucket/path/").List(context.Background(), 3000)
=> ["gcs://bucket/path/file.xml"]
```

### Delete a file
```go
bucket.For("gcs://bucket/path/file.xml").Delete(context.Background())
=> true
```

## Development

### Running tests
WIP

## License & Contributing

* Masa is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
* Bug reports and pull requests are welcome on GitHub at https://github.com/ogwurujohnson/masa.
