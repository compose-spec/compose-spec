package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func createComposeSpec(spec string) string {
	return fmt.Sprintf(`
  {
    "services": {
        "a": {
          "image": "alpine"
        }
    },
    %s
  }`, spec)
}

func createComposeSpecService(serviceSpec string) string {
	return fmt.Sprintf(`
  {
    "services": {
      "service-foo": {
          %s
      }
    }
  }`, serviceSpec)
}

func TestComposeSpec(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		invalid bool
	}{
		{
			name: "[#/nertwork] network ipv4 and ipv6 subnet",
			spec: createComposeSpec(`
        "networks": {
          "front-tier": {
            "ipam": {
              "driver": "default",
              "config": [
                {"subnet": "172.16.238.0/24"},
                {"subnet": "0.0.0.0/64"},
                {"subnet": "2001:3984:3989::/64"},
                {"subnet": "172.16.238.0/24"},
                {"subnet": "2001:3984:3989::/64"},
                {"subnet": "2001:3984:3989::/64"},
                {"subnet": "1080:0:0:0:8:800:200C:417A/64"},
                {"subnet": "FF01:0:0:0:0:0:0:101/64"},
                {"subnet": "0:0:0:0:0:0:0:1/64"},
                {"subnet": "0:0:0:0:0:0:0:0/64"},
                {"subnet": "::/64"},
                {"subnet": "1080::8:800:200C:417A/64"},
                {"subnet": "::1/64"}
              ]
            }
          }
        }`),
		},
		{
			name: "[#/x-*] optional x- should be allowed at top level",
			spec: createComposeSpec(`"x-test":"something"`),
		},
		{
			name: "[#/services/service/expose] expose port has right format",
			spec: createComposeSpecService(`
        "expose": [
          "12345-12345/tcp",
          "12345-12345/udp",
          "12345/udp",
          "12345-12345",
          "1234-1234",
          "123-123",
          "12-12",
          "1-1",
          "12345",
          "1234",
          "123",
          "12",
          "1"
        ]`),
		},
		{
			name: "[#/services/service/expose] expose port is too long",
			spec: createComposeSpecService(`
        "expose": [
          "123456"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/expose] expose second port is too long",
			spec: createComposeSpecService(`
        "expose": [
          "1234-123456"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/expose] expose missing first port",
			spec: createComposeSpecService(`
        "expose": [
          "-123"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/expose] expose with wrong protocol",
			spec: createComposeSpecService(`
        "expose": [
          "123-123/random"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] ports must have the [HOST:]CONTAINER[/PROTOCOL] format",
			spec: createComposeSpecService(`
        "ports": [
          "3000",
          "3000-3005",
          "8000:8000",
          "9090-9091:8080-8081",
          "49100:22",
          "8000-9000:80",
          "127.0.0.1:8001:8001",
          "127.0.0.1:5000-5010:5000-5010",
          "6060:6060/udp"
        ]`),
		},
		{
			name: "[#/services/service/ports] ports cannot be empty",
			spec: createComposeSpecService(`
        "ports": [
          ""
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] ports cannot be negative",
			spec: createComposeSpecService(`
        "ports": [
          "-3000"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] ports cannot be imcomplete",
			spec: createComposeSpecService(`
        "ports": [
          ":80"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] ports must have valid protocol",
			spec: createComposeSpecService(`
        "ports": [
          "6060:6060/random"
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] port protocol should be tcp or udp",
			spec: createComposeSpecService(`
        "ports": [
          { "protocol": "tcp" },
          { "protocol": "udp" }
        ]`),
		},
		{
			name: "[#/services/service/ports] port protocol should ONLY be tcp or udp",
			spec: createComposeSpecService(`
        "ports": [
          { "protocol": "random" }
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] port published within range",
			spec: createComposeSpecService(`
        "ports": [
          { "published": "0" },
          { "published": "00000" },
          { "published": "0-0" },
          { "published": "00000-00000" },
          { "published": "65535" },
          { "published": "65535-65535" },
          { "published": 0 },
          { "published": 65535 }
        ]`),
		},
		{
			name: "[#/services/service/ports] port published should not be over 65535 (integer)",
			spec: createComposeSpecService(`
        "ports": [
          { "published": 65536 }
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] published should not be over 65535 (string)",
			spec: createComposeSpecService(`
        "ports": [
          { "published": "65536" }
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/ports] port published should not be over 65535 (string)",
			spec: createComposeSpecService(`
        "ports": [
          { "published": "65535-65536" }
        ]`),
			invalid: true,
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (micro second)",
			spec: createComposeSpecService(`"stop_grace_period": "10us"`),
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (milisecond)",
			spec: createComposeSpecService(`"stop_grace_period": "10ms"`),
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (second)",
			spec: createComposeSpecService(`"stop_grace_period": "10s"`),
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (minute)",
			spec: createComposeSpecService(`"stop_grace_period": "10m"`),
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (hour)",
			spec: createComposeSpecService(`"stop_grace_period": "10h"`),
		},
		{
			name: "[#/services/service/stop_grace_period] must be a valid duration (combined)",
			spec: createComposeSpecService(`"stop_grace_period": "1h2m3s4ms5us"`),
		},
		{
			name:    "[services/service/stop_grace_period] cannot be given a duration with wrong order",
			spec:    createComposeSpecService(`"stop_grace_period": "1s10h"`),
			invalid: true,
		},
		{
			name:    "stop_grace_period cannot be given a duration with wrong unit",
			spec:    createComposeSpecService(`"stop_grace_period": "1kg"`),
			invalid: true,
		},
		{
			name: "[#/services/service/deploy/rollback_config/delay] is a duration",
			spec: createComposeSpecService(`"deploy": {"rollback_config": {"delay": "1h2s3ms4us"}}`),
		},
		{
			name: "[#/services/service/deploy/rollback_config/monitor] is a duration",
			spec: createComposeSpecService(`"deploy": {"rollback_config": {"monitor": "1h2s3ms4us5ns"}}`),
		},
		{
			name:    "[services/service/develop/watch] should require at least one element",
			spec:    createComposeSpecService(`"develop": { "watch": [] }`),
			invalid: true,
		},
		{
			name:    "[services/service/develop/watch] should require at least one element with path and action",
			spec:    createComposeSpecService(`"develop": { "watch": [{"action":"rebuild"}] }`),
			invalid: true,
		},
		{
			name: "[#/services/service/develop/watch] should require action and path",
			spec: createComposeSpecService(`
			"develop": {
				"x-develop": "foo",
				"watch": [
					{ 
						"action": "sync", 
						"path": "foo",
						"x-watch": "foo"
					}
				]
			}`),
		},
		{
			name: "[#/include] should be a string array",
			spec: createComposeSpec(`"include": [ "foo" ]`),
		},
		{
			name: "[#/include] should be a object array",
			spec: createComposeSpec(`"include": [ { "path": "foo"} ]`),
		},
	}
	sch, err := jsonschema.Compile("../schema/compose-spec.json")
	if err != nil {
		t.Fatalf("\033[31mcompose-spec.json is not valid. Skipping tests.\n\033[0m--- FAIL: %s", err)
		return
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var data interface{}
			err = json.Unmarshal([]byte(tt.spec), &data)
			if err != nil {
				t.Logf("an error was thrown on Unmarshal. NOTE: make sure in array test element case does not have a ',' in the last element\n%s\n", tt.spec)
				t.FailNow()
			}

			err = sch.Validate(data)
			if !tt.invalid {
				if err != nil {
					t.Logf("the spec should be valid\n%s\nERROR: %s", tt.spec, err)
					t.FailNow()
				}
			} else {
				if err == nil {
					t.Logf("the spec should NOT be valid\n%s\nERROR: %s", tt.spec, err)
					t.Fail()
				}
			}
		})
	}
}
