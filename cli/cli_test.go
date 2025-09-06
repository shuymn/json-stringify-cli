package cli

import (
	"bytes"
	"io"
	"testing"
)

func TestRun(t *testing.T) {
	testcases :=
		[]struct {
			path, want string
		}{
			{
				path: "testdata/sample-01.json",
				want: `"[100,500,300,200,400]"`,
			},
			{
				path: "testdata/sample-02.json",
				want: `"[{\"color\":\"red\",\"value\":\"#f00\"},{\"color\":\"green\",\"value\":\"#0f0\"},{\"color\":\"blue\",\"value\":\"#00f\"},{\"color\":\"cyan\",\"value\":\"#0ff\"},{\"color\":\"magenta\",\"value\":\"#f0f\"},{\"color\":\"yellow\",\"value\":\"#ff0\"},{\"color\":\"black\",\"value\":\"#000\"}]"`,
			},
			{
				path: "testdata/sample-03.json",
				want: `"{\"color\":\"red\",\"value\":\"#f00\"}"`,
			},
			{
				path: "testdata/sample-04.json",
				want: `"{\"id\":\"0001\",\"type\":\"donut\",\"name\":\"Cake\",\"ppu\":0.55,\"batters\":{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"},{\"id\":\"1002\",\"type\":\"Chocolate\"},{\"id\":\"1003\",\"type\":\"Blueberry\"},{\"id\":\"1004\",\"type\":\"Devil's Food\"}]},\"topping\":[{\"id\":\"5001\",\"type\":\"None\"},{\"id\":\"5002\",\"type\":\"Glazed\"},{\"id\":\"5005\",\"type\":\"Sugar\"},{\"id\":\"5007\",\"type\":\"Powdered Sugar\"},{\"id\":\"5006\",\"type\":\"Chocolate with Sprinkles\"},{\"id\":\"5003\",\"type\":\"Chocolate\"},{\"id\":\"5004\",\"type\":\"Maple\"}]}"`,
			},
			{
				path: "testdata/sample-05.json",
				want: `"[{\"id\":\"0001\",\"type\":\"donut\",\"name\":\"Cake\",\"ppu\":0.55,\"batters\":{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"},{\"id\":\"1002\",\"type\":\"Chocolate\"},{\"id\":\"1003\",\"type\":\"Blueberry\"},{\"id\":\"1004\",\"type\":\"Devil's Food\"}]},\"topping\":[{\"id\":\"5001\",\"type\":\"None\"},{\"id\":\"5002\",\"type\":\"Glazed\"},{\"id\":\"5005\",\"type\":\"Sugar\"},{\"id\":\"5007\",\"type\":\"Powdered Sugar\"},{\"id\":\"5006\",\"type\":\"Chocolate with Sprinkles\"},{\"id\":\"5003\",\"type\":\"Chocolate\"},{\"id\":\"5004\",\"type\":\"Maple\"}]},{\"id\":\"0002\",\"type\":\"donut\",\"name\":\"Raised\",\"ppu\":0.55,\"batters\":{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"}]},\"topping\":[{\"id\":\"5001\",\"type\":\"None\"},{\"id\":\"5002\",\"type\":\"Glazed\"},{\"id\":\"5005\",\"type\":\"Sugar\"},{\"id\":\"5003\",\"type\":\"Chocolate\"},{\"id\":\"5004\",\"type\":\"Maple\"}]},{\"id\":\"0003\",\"type\":\"donut\",\"name\":\"Old Fashioned\",\"ppu\":0.55,\"batters\":{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"},{\"id\":\"1002\",\"type\":\"Chocolate\"}]},\"topping\":[{\"id\":\"5001\",\"type\":\"None\"},{\"id\":\"5002\",\"type\":\"Glazed\"},{\"id\":\"5003\",\"type\":\"Chocolate\"},{\"id\":\"5004\",\"type\":\"Maple\"}]}]"`,
			},
			{
				path: "testdata/sample-06.json",
				want: `"{\"id\":\"0001\",\"type\":\"donut\",\"name\":\"Cake\",\"image\":{\"url\":\"images/0001.jpg\",\"width\":200,\"height\":200},\"thumbnail\":{\"url\":\"images/thumbnails/0001.jpg\",\"width\":32,\"height\":32}}"`,
			},
			{
				path: "testdata/sample-07.json",
				want: `"{\"items\":{\"item\":[{\"id\":\"0001\",\"type\":\"donut\",\"name\":\"Cake\",\"ppu\":0.55,\"batters\":{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"},{\"id\":\"1002\",\"type\":\"Chocolate\"},{\"id\":\"1003\",\"type\":\"Blueberry\"},{\"id\":\"1004\",\"type\":\"Devil's Food\"}]},\"topping\":[{\"id\":\"5001\",\"type\":\"None\"},{\"id\":\"5002\",\"type\":\"Glazed\"},{\"id\":\"5005\",\"type\":\"Sugar\"},{\"id\":\"5007\",\"type\":\"Powdered Sugar\"},{\"id\":\"5006\",\"type\":\"Chocolate with Sprinkles\"},{\"id\":\"5003\",\"type\":\"Chocolate\"},{\"id\":\"5004\",\"type\":\"Maple\"}]}]}}"`,
			},
		}

	for _, tc := range testcases {
		buf := new(bytes.Buffer)
		c := New(tc.path, buf)
		err := c.Run()
		if err != nil {
			t.Fatalf("want no error. got: %s", err)
		}

		if tc.want != buf.String() {
			t.Errorf("want: %q. got: %q", tc.want, buf.String())
		}
	}
}

func TestRun_error(t *testing.T) {
	testcases := []struct {
		subtitle, path, want string
	}{
		{
			subtitle: "invalid json",
			path:     "testdata/sample-08.json",
			want:     "failed to compact JSON: invalid character '\\n' in string literal",
		},
		{
			subtitle: "empty file",
			path:     "testdata/sample-09.json",
			want:     "failed to compact JSON: unexpected end of JSON input",
		},
		{
			subtitle: "file not exist",
			path:     "testdata/sample-99.json",
			want:     "failed to read file: open testdata/sample-99.json: no such file or directory",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.subtitle, func(t *testing.T) {
			c := New(tc.path, io.Discard)
			err := c.Run()
			if err == nil {
				t.Fatal("want error. got nil")
			}
			if err.Error() != tc.want {
				t.Errorf("want: %q. got: %q", tc.want, err.Error())
			}
		})
	}
}
