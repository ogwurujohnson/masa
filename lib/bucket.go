package bucket

import (
	"net/url"

	keyStorage "github.com/ogwurujohnson/masa/lib/services"
)

func parse(key string) keyStorage.InitType {
	uri, err := url.Parse(key)
	if err != nil {
		panic(err)
	}

	// TODO: figure out the escape and unescape stuff on the ruby repo
	return keyStorage.InitType{
		AdapterType: uri.Scheme,
		Bucket:      uri.Host,
		Key:         uri.Path,
	}
}

//  Given a `key` in the format of `adapter://bucket/key` returns the corresponding
//  adapter that will allow to manipulate (e.g. download, upload or list) such key.
// 
//  Currently supported adapters are `gcs` (Google Cloud Storage), `s3` (AWS S3)
// 
//  @param [String] key The reference key
//  @return [KeyStorage] An interface to the adapter that can handle requests on the given key
//  @example Configure {BucketStore} for Google Cloud Storage
//  BucketStore.for("gcs://the_bucket/a/valid/key")
func For(key string) *keyStorage.ServiceMappers {
	data := parse(key)

	return keyStorage.Initialize(data)
}
