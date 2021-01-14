package symon_agen

import (
	"fmt"
	"net/http"
	"strings"
)

func ExtractResponseHeader(response *http.Response, ctx JobContext) {
	for key, value := range response.Header {
		ctx.Set(fmt.Sprintf("http.response.header.%s", strings.ToLower(key)), strings.Join(value, ","))
	}
}
