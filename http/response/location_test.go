package main

import (
	"net/http"
	"net/url"
	"testing"
)

type responseLocationTest struct {
	location string // Response's Location header or ""
	requrl   string // Response.Request.URL or ""
	want     string
	wantErr  error
}

var responseLocationTests = []responseLocationTest{
	{"/foo", "http://bar.com/baz", "http://bar.com/foo", nil},
	{"http://foo.com/", "http://bar.com/baz", "http://foo.com/", nil},
	{"", "http://bar.com/baz", "", http.ErrNoLocation},
	{"/bar", "", "/bar", nil},
}

func TestLocationResponse(t *testing.T) {
	for i, tt := range responseLocationTests {
		res := new(http.Response)
		res.Header = make(http.Header)
		res.Header.Set("Location", tt.location)
		if tt.requrl != "" {
			res.Request = &http.Request{}
			var err error
			res.Request.URL, err = url.Parse(tt.requrl)
			if err != nil {
				t.Fatalf("bad test URL %q: %v", tt.requrl, err)
			}
		}

		got, err := res.Location()
		t.Log(got)
		if tt.wantErr != nil {
			if err == nil {
				t.Errorf("%d. err=nil; want %q", i, tt.wantErr)
				continue
			}
			if g, e := err.Error(), tt.wantErr.Error(); g != e {
				t.Errorf("%d. err=%q; want %q", i, g, e)
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("%d. err=%q", i, err)
			continue
		}
		if g, e := got.String(), tt.want; g != e {
			t.Errorf("%d. Location=%q; want %q", i, g, e)
		}
	}
}
