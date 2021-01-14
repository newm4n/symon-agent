package symon_agen

import (
	"testing"
)

func TestHttpCheckJob_Execute(t *testing.T) {
	ctx := NewDefaultContext()
	ctx.Set("http.monitor.count", "2")

	ctx.Set("0.http.request.url", "https://google.com")
	ctx.Set("0.http.request.method", "GET")
	ctx.Set( "0.http.request.header.User-Agent", "symon-agent.v0.0.1")
	ctx.Set( "0.http.response.validation.code", "200")

	ctx.Set("1.http.request.url", "https://detik.com")
	ctx.Set("1.http.request.method", "GET")
	ctx.Set( "1.http.request.header.User-Agent", "symon-agent.v0.0.1")
	ctx.Set( "1.http.request.header.Authentication", "Bearer ${0.http.response.body.$JSON.access}")
	ctx.Set( "1.http.response.validation.code", "200")
	ctx.Set( "1.http.response.validation.headerkeys.exist", "Content-Type,Server")

	job := &HttpCheckJob{}
	job.Execute(ctx)

	t.Log(ctx.String())

	job.Validate(ctx)

	if ctx.Get("1.http.response.valid") != "true" {
		t.Log("It should not be valid")
		t.Fail()
	}
}
