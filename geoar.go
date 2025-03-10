package geoar

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os"
	"strings"

	"github.com/ctomatis/georefar-go/credentials"
	"github.com/ctomatis/georefar-go/internal"
	qs "github.com/google/go-querystring/query"
)

const (
	BASE_URL = "https://apis.datos.gob.ar/georef/api"

	JSON    Output = "json"
	CSV     Output = "csv"
	XML     Output = "xml"
	GEOJSON Output = "geojson"
	NDJSON  Output = "ndjson"
)

type (
	Output  string
	Payload map[string][]any

	Config struct {
		BaseUrl, Key, Secret string
		Debug                bool
	}

	httpClient struct {
		config   *Config
		client   *http.Client
		response *http.Response
		request  *http.Request
		ctx      *context.Context
	}

	requestError struct {
		ErrorInfo    string
		ErrorMessage string
		StatusCode   int
		Errores      []struct {
			Mensaje string `json:"mensaje"`
		} `json:"errores"`
	}
)

func (e requestError) Error() string {
	return fmt.Sprintf("%s %d. %s\n%s", http.StatusText(e.StatusCode), e.StatusCode, e.ErrorMessage, e.ErrorInfo)
}

func New(cfg *Config) *httpClient {
	if cfg.BaseUrl == "" {
		cfg.BaseUrl = BASE_URL
	}
	return &httpClient{
		client: &http.Client{},
		config: cfg,
	}
}

func (c *httpClient) WithContext(ctx context.Context) *httpClient {
	c.ctx = &ctx
	return c
}

func (c *httpClient) Send(resource any, filters ...*filter) *httpClient {
	url := buildUrl(c.config.BaseUrl, resource, filters...)
	req, err := createRequest(url, nil)
	if err != nil {
		return c
	}
	c.request = req
	return c
}

func (c *httpClient) Json() (data *Json, err error) {
	err = c.doRequest()
	if err != nil {
		return nil, err
	}
	defer c.reset()
	err = json.NewDecoder(c.response.Body).Decode(&data)
	return
}

func (c *httpClient) Csv() (data []byte, err error) {
	c.formato(CSV)
	err = c.doRequest()
	if err != nil {
		return nil, err
	}
	defer c.reset()
	data, err = io.ReadAll(c.response.Body)
	return
}

func (c *httpClient) Download(resource any, output Output) *httpClient {
	url := buildUrl(c.config.BaseUrl, resource)
	url = fmt.Sprintf("%s.%s", strings.ReplaceAll(url, "-", "_"), output)

	req, err := createRequest(url, nil)
	if err != nil {
		return c
	}
	c.request = req
	return c
}

func (c *httpClient) Bulk(resources []any, filters ...*filter) (data *BulkJson, err error) {
	path := internal.Resource(resources[0])
	url := fmt.Sprintf("%s/%s", c.config.BaseUrl, path)

	payload := buildPayload(path, resources, filters...)

	req, err := createRequest(url, payload)
	if err != nil {
		return
	}

	c.request = req
	err = c.doRequest()
	if err != nil {
		return
	}
	defer c.reset()

	err = json.NewDecoder(c.response.Body).Decode(&data)
	return
}

func (c *httpClient) Save(fname ...string) (b int, err error) {
	err = c.doRequest()
	if err != nil {
		return
	}
	defer c.reset()

	body, err := io.ReadAll(c.response.Body)
	if err != nil {
		return
	}

	var filename string
	if len(fname) > 0 {
		filename = fname[0]
	} else {
		i := strings.LastIndex(c.request.URL.Path, "/") + 1
		filename = c.request.URL.Path[i:]
	}

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	return f.Write(body)
}

func (c *httpClient) doRequest() error {
	if c.request == nil {
		return errors.New("request must be non-nil")
	}

	if c.request.Method == http.MethodPost {
		c.request.Header.Set("Content-Type", "application/json")
	}

	if c.config.Secret != "" && c.config.Key != "" {
		if token, err := credentials.New(c.config.Key, c.config.Secret).GetAuthentication(); err == nil {
			c.request.Header.Set("Authorization", token)
		}
	}

	if c.ctx != nil {
		c.request = c.request.WithContext(*c.ctx)
	}

	res, err := c.client.Do(c.request)
	if err != nil {
		return err
	}

	if err = checkResponseError(res); err != nil {
		return err
	}

	c.response = res
	return err
}

func (c *httpClient) formato(f Output) {
	qs := c.request.URL.Query()
	qs.Add("formato", string(f))
	c.request.URL.RawQuery = qs.Encode()
}

func (c *httpClient) reset() {
	if c.response != nil {
		c.response.Body.Close()
	}
}

func createRequest(url string, body any) (*http.Request, error) {
	method := http.MethodGet
	var payload []byte
	if body != nil {
		method = http.MethodPost
		payload, _ = json.Marshal(body)
	}
	return http.NewRequest(method, url, bytes.NewBuffer(payload))
}

func buildPayload(key string, resources []any, filters ...*filter) Payload {
	nf := len(filters)
	payload := make(Payload)
	for i := range resources {
		m1 := internal.ToMap(resources[i])
		if i < nf {
			maps.Copy(m1, internal.ToMap(filters[i]))
		}
		payload[key] = append(payload[key], m1)
	}
	return payload
}

func buildUrl(base string, resource any, filters ...*filter) string {
	buf := new(strings.Builder)
	buf.WriteString(base + "/")
	buf.WriteString(internal.Resource(resource))

	params := []string{}
	if v, _ := qs.Values(resource); len(v) > 0 {
		params = append(params, v.Encode())
	}

	if len(filters) > 0 {
		if v, _ := qs.Values(filters[0]); len(v) > 0 {
			params = append(params, v.Encode())
		}
	}

	if len(params) > 0 {
		buf.WriteString("?")
		buf.WriteString(strings.Join(params, "&"))
	}
	return buf.String()
}

func checkResponseError(res *http.Response) error {
	if res == nil {
		return fmt.Errorf("response must be non-nil")
	}

	if res.StatusCode != 200 {
		Err := requestError{
			StatusCode: res.StatusCode,
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			Err.ErrorMessage = "unable to read response body"
			Err.ErrorInfo = err.Error()
			return Err
		}

		if err = json.Unmarshal(b, &Err); err != nil {
			Err.ErrorMessage = "unexpected server response: %s" + string(b)
			Err.ErrorInfo = "json unmarshal error: " + err.Error()
			return Err
		}

		if len(Err.Errores) > 0 {
			Err.ErrorMessage = Err.Errores[0].Mensaje
			Err.ErrorInfo = string(b)
		}
		return Err
	}
	return nil
}
