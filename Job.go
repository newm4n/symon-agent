package symon_agen

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Job interface {
	Execute(ctx MonitorContext)
	Validate(ctx MonitorContext)
}

type HttpCheckJob struct {
}

func (job *HttpCheckJob) Execute(ctx MonitorContext) {

	numCals, err := strconv.Atoi(ctx["http.monitor.count"])
	if err != nil || numCals <= 0 {
		ctx["http.response.haveerror"] = "true"
		ctx["http.response.error"] = "http.monitor.count is not provided or not an integer"
		return
	}

	for callIndex := 0; callIndex < numCals; callIndex++ {
		UrlToInvoke := ctx[fmt.Sprintf("%d.http.request.url", callIndex)]
		Method := ctx[fmt.Sprintf("%d.http.request.method", callIndex)]
		if len(UrlToInvoke) == 0 {
			ctx[fmt.Sprintf("%d.http.response.haveerror", callIndex)] = "true"
			ctx[fmt.Sprintf("%d.http.response.error", callIndex)] = "URL (http.request.url) is not provided in context"
			return
		}
		if len(Method) == 0 {
			ctx[fmt.Sprintf("%d.http.response.haveerror", callIndex)] = "true"
			ctx[fmt.Sprintf("%d.http.response.error", callIndex)] = "HTTP method (http.request.method) is not provided in context"
			return
		}
		request, err := http.NewRequest(Method, UrlToInvoke, nil)
		if err != nil {
			ctx[fmt.Sprintf("%d.http.response.haveerror", callIndex)] = "true"
			ctx[fmt.Sprintf("%d.http.response.error", callIndex)] = err.Error()
			return
		}
		for key, value := range ctx {
			if strings.HasPrefix(key, fmt.Sprintf("%d.http.request.header.", callIndex)) {
				headerKey := key[20:]
				headerValues := strings.Split(value, ",")
				request.Header[headerKey] = headerValues
			}
		}
		if ctx.Contains(fmt.Sprintf("%d.http.request.body", callIndex)) {
			// TODO make this smarter, when request content type is a type of text we do not need to convert into Base64
			base64RequestBody := ctx[fmt.Sprintf("%d.http.request.body", callIndex)]
			byteBody, err := base64.StdEncoding.DecodeString(base64RequestBody)
			if err != nil {
				ctx[fmt.Sprintf("%d.http.response.haveerror", callIndex)] =  "true"
				ctx[fmt.Sprintf("%d.http.response.error", callIndex)] = err.Error()
				return
			}
			if len(byteBody) > 0 {
				request.Body = ioutil.NopCloser(bytes.NewReader(byteBody))
				request.ContentLength = int64(len(byteBody))
			}
		}

		client := &http.Client{}

		if ctx.GetBool(fmt.Sprintf("%d.http.request.insecure-skip-verify", callIndex)) {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client.Transport = tr
		}

		start := time.Now()
		response, err := client.Do(request)
		duration := time.Since(start)
		ctx[fmt.Sprintf("%d.http.response.duration", callIndex)] = fmt.Sprintf("%d", duration / time.Millisecond)
		if err != nil {
			ctx[fmt.Sprintf("%d.http.response.haveerror", callIndex)]= "true"
			ctx[fmt.Sprintf("%d.http.response.error", callIndex)]= err.Error()
			return
		}
		defer response.Body.Close()
		ctx[fmt.Sprintf("%d.http.response.code", callIndex)] = fmt.Sprintf("%d",response.StatusCode)
		for key, value := range response.Header {
			ctx[fmt.Sprintf("%d.http.response.header.%s",callIndex, key)] = strings.Join(value, ",")
		}
		bodyByte, err := ioutil.ReadAll(response.Body)
		// TODO make this smarter, when response content type is a type of text we do not need to convert into Base64
		if err == nil && bodyByte != nil && len(bodyByte) > 0 {
			base64Body := base64.StdEncoding.EncodeToString(bodyByte)
			ctx[fmt.Sprintf("%d.http.response.body",callIndex)] = base64Body
		}
	}
}

func (job *HttpCheckJob) Validate(ctx MonitorContext) {
	numCals, err := strconv.Atoi(ctx["http.monitor.count"])
	if err != nil {
		ctx["http.response.valid"] = "false"
		ctx["http.response.valid.reason"] = "missing response code in context"
		return
	}

	for callIndex := 0; callIndex < numCals; callIndex++ {
		// validate for response code
		if ctx.Contains(fmt.Sprintf("%d.http.response.validation.code", callIndex)) {
			if ctx.Contains(fmt.Sprintf("%d.http.response.code", callIndex)) {
				fmt.Println(ctx[fmt.Sprintf("%d.http.response.code", callIndex)])
				if ctx[fmt.Sprintf("%d.http.response.validation.code", callIndex)] != ctx[fmt.Sprintf("%d.http.response.code", callIndex)] {
					ctx[fmt.Sprintf("%d.http.response.valid", callIndex)] = "false"
					ctx[fmt.Sprintf("%d.http.response.valid.reason", callIndex)] = "invalid response code."
					return
				}
			} else {
				ctx[fmt.Sprintf("%d.http.response.valid", callIndex)]= "false"
				ctx[fmt.Sprintf("%d.http.response.valid.reason", callIndex)]= "missing response code in context"
				return
			}
		}

		if ctx.Contains(fmt.Sprintf("%d.http.response.validation.headerkeys.exist", callIndex)) {
			requiredHeaderkeys := strings.Split(ctx[fmt.Sprintf("%d.http.response.validation.headerkeys.exist", callIndex)] , ",")
			for _,k := range requiredHeaderkeys {
				keyToLook := fmt.Sprintf("%d.http.response.header.%s", callIndex, k)
				if !ctx.Contains(keyToLook) {
					ctx[fmt.Sprintf("%d.http.response.valid", callIndex)]= "false"
					ctx[fmt.Sprintf("%d.http.response.valid.reason", callIndex)]= "missing response code in context"
					return
				}
			}
		}

		// validate for maximum duration
		// validate header existance
		ctx[fmt.Sprintf("%d.http.response.valid", callIndex)] = "true"
	}
}