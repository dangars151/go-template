package postutil

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var client = resty.New().R()
var baseUrl = "https://jsonplaceholder.typicode.com"

// data được lấy từ trang này https://jsonplaceholder.typicode.com/posts
func GetPosts() ([]PostResponse, error) {
	url := baseUrl + "/posts"
	postsResponse := make([]PostResponse, 0)
	res, err := client.SetResult(&postsResponse).Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "utils.postutil.api failed when get posts")
	}
	if res.IsError() {
		return nil, errors.Wrap(fmt.Errorf("%v", res.Error()), "utils.postutil.api failed when get posts")
	}
	return postsResponse, nil
}
