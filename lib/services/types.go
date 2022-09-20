package services

type Response struct {
	Bucket  string
	Key     string
	Keys    []string
	Content interface{}
}
