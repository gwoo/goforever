# greq
Very simple go http request library. Every method returns the string body, `http.Response`, and err.

## Supported Methods
There is a more generic `Do` function. But the following wrappers will be easier.

- Get
- Post
- Put
- Delete

## Usage

	import ("github.com/gwoo/greq")
	req := greq.New("http://google.com", false) //Use form-encoded data. If true, use json body.
	body, response, err := req.Get("/")

	var data map[string]interface{}
	body, response, err := req.Post("/search", data{"q": "sailboats"})
