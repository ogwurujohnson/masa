package bucket

import (
	"net/url"

	"github.com/ogwurujohnson/masa/lib/services"
)


func parse(key string) services.InitType {
	uri, err := url.Parse(key)
	if err != nil {
		panic(err)
	}

	// TODO: figure out the escape and unescape stuff on the ruby repo
	return services.InitType{
		AdapterType: uri.Scheme,
		Bucket: uri.Host,
		Key: uri.Path,
	}
}

func For(key string) *services.ServiceMappers {
	ctx := parse(key)
	return services.Initialize(ctx)
}