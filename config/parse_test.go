package config_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/telenornms/skogul"
	"github.com/telenornms/skogul/config"
	"github.com/telenornms/skogul/receiver"
	"github.com/telenornms/skogul/sender"
)

func TestFile(t *testing.T) {
	defer func() {
		if skogul.AssertErrors > 0 {
			t.Errorf("File() paniced")
		}
	}()
	c, err := config.File("testdata/test.json")
	if err != nil {
		t.Errorf("File() failed: %v", err)
	}
	if c == nil {
		t.Errorf("File() returned nil config")
	}
}

func TestByte_ok(t *testing.T) {
	defer func() {
		if skogul.AssertErrors > 0 {
			t.Errorf("Byte() paniced")
		}
	}()
	var okData []byte
	okData = []byte(`
{
  "senders": {
    "tnet_alarms": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INERT INTO tnet_Alarms values(${foo})"
    },
    "mysql_log": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INSERT INTO liksomlog VALUES(${timestmap.timestamp},${metadata.name},${key})"
    },
    "forward": {
	    "type": "http",
	    "url": "http://localhost:8888"
    },
    "duplicate": {
	    "type": "dupe",
	    "next": ["forward","mysql_log","tnet_alarms"]
    },
    "batch": {
      "type": "batch",
      "next": "duplicate",
      "interval": "5s"
    },
    "det": {
	    "type": "detacher",
	    "next": "batch"
    }
  },
  "handlers": {
    "plain": {
      "parser": "json",
      "sender": "batch",
      "transformers": ["foo"]
    }
  },
  "transformers": {
	  "foo": {
		  "type": "templater"
	  }
  },
  "receivers": {
    "http": {
      "type": "http",
      "handlers": {
	      "/": "plain"
      }
    }
  }
}
`)
	c, err := config.Bytes(okData)
	if err != nil {
		t.Errorf("Bytes() failed: %v", err)
	}
	if c == nil {
		t.Errorf("Bytes() returned nil config")
	}
	badData := []byte(`{ "senders": { "x": { "type": "sql", "ConnStr": 5 } } }`)
	c, err = config.Bytes(badData)
	if err == nil {
		t.Errorf("Bytes() test 2 failed, sent bad data, didn't get error.")
	}
	if c != nil {
		t.Errorf("Bytes() with bad data returned valid config.")
	}
	noURL := []byte(`{ "senders": { "x" : { "type": "http" }}}`)
	c, err = config.Bytes(noURL)
	if err == nil {
		t.Errorf("Bytes() test 3 failed, http sender with no URL didn't get error.")
	}
	if c != nil {
		t.Errorf("Bytes() with bad data returned valid config.")
	}

}

func TestHelpModule(t *testing.T) {
	_, err := config.HelpModule(sender.Auto, "sql")
	if err != nil {
		t.Errorf("HelpModule(sender.Auto,\"sql\") didn't work: %v", err)
	}
}

func testBadConf(t *testing.T, badData string) {
	t.Helper()
	_, err := config.Bytes([]byte(badData))
	if err == nil {
		t.Errorf("Bytes() was ok, despite bad data")
	}
}

// Useful for visual confirmation - e.g. - check that the arrow points at
// the right thing. Important to test errors early in a config, in the
// middle and towards the end.
func Test_syntaxError(t *testing.T) {
	testBadConf(t,
		`{
  "senders": {
    "tnet_alarms":: {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INERT INTO tnet_Alarms values(${foo})"
    }
  }
}`)

	testBadConf(t,
		`x{
  "senders": {
    "tnet_alarms":: {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INERT INTO tnet_Alarms values(${foo})"
    }
  }
}`)
	testBadConf(t,
		`{
  "senders": {
    "tnet_alarms": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INERT INTO tnet_Alarms values(${foo})"
    },
    "mysql_log": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INSERT INTO liksomlog VALUES(${timestmap.timestamp},${metadata.name},${key})"
    },
    "forward": {
	    "type": "http",
	    "url": "http://localhost:8888"
    },
    "duplicate": {
	    "type": "dupe",
	    "next": ["forward","mysql_log","tnet_alarms"]
    },
    "batch": {
      "type" "batch",
      "next": "duplicate",
      "interval": "5s"
    },
    "det": {
	    "type": "detacher",
	    "next": "batch"
    }
  },
  "handlers": {
    "plain": {
      "parser": "json",
      "sender": "batch",
      "transformers": []
    }
  },
  "receivers": {
    "http": {
      "type": "http",
      "handlers": {
	      "/": "plain"
      }
    }
  }
}`)
	testBadConf(t,
		`{
  "senders": {
    "tnet_alarms": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INERT INTO tnet_Alarms values(${foo})"
    },
    "mysql_log": {
      "type": "sql",
      "Driver": "mysql",
      "ConnStr": "root:lol@/mydb",
      "Query": "INSERT INTO liksomlog VALUES(${timestmap.timestamp},${metadata.name},${key})"
    },
    "forward": {
	    "type": "http",
	    "url": "http://localhost:8888"
    },
    "duplicate": {
	    "type": "dupe",
	    "next": ["forward","mysql_log","tnet_alarms"]
    },
    "batch": {
      "type": "batch",
      "next": "duplicate",
      "interval": "5s"
    },
    "det": {
	    "type": "detacher",
	    "next": "batch"
    }
  },
  "handlers": {
    "plain": {
      "parser": "json",
      "sender": "batch",
      "transformers": []
    }
  },
  "receivers": {
    "http": {
      "type": "http",
      "handlers": {
	      "/": "plain"
      }
    }
  }`)
	testBadConf(t,
		`{
	"senders": {
		"tnet_alarms":: {
			"type": "sql",
			"Driver": "mysql",
			"ConnStr":: "root:lol@/mydb",
			"Query": "INERT INTO tnet_Alarms values(${foo})"
		}
	}
}`)

	_, err := config.Bytes([]byte(`{
    "receivers": {
      "udp": {
        "type": "udp",
        "address": "[::1]:5015",
        "handler": "protobuf"
      }
    },
    "handlers": {
      "protobuf": {
        "parser": "protobuf",
        "transformers": ["remove", "remove2", "remove3", "remove4", "remove5"],
        "sender": "print"
      }
    },
    "transformers": {
      "sw": {
        "type": "switch",
        "cases": [
          {
            "when": "sensorName",
            "is": "junos_system_linecard_intf-exp:/junos/system/linecard/intf-exp/:/junos/system/linecard/intf-exp/:PFE",
            "transformers": ["remove"]
          }
        ]
      },
      "remove": {
        "type": "data",
        "remove": ["interfaceExp_stats"]
      },
      "remove2": {
        "type": "data",
        "remove": ["interfaceExp_stats"]
      },
      "remove3": {
        "type": "data",
        "remove": ["interfaceExp_stats"]
      },
      "remove4": {
        "type": "data",
        "remove": ["interfaceExp_stats"]
      },
      "remove5": {
        "type": "data",
        "remove": ["interfaceExp_stats"]
      }
    },
    "senders": {
      "print": {
        "type": "debug",
        "prefix": "DEBUG"
      }
    }
  }`))
	if err != nil {
		t.Errorf("Expected config to pass: %v", err)
	}
}
func TestFindSuperfluousReceiverConfigProperties(t *testing.T) {
	rawConfig := []byte(`{"receivers": {
		"foo": {
		  "type": "udp",
		  "address": "[::1]:5015",
		  "superfluousField": "this is not needed"
		}
	  }
	}`)

	var c map[string]interface{}
	err := json.Unmarshal(rawConfig, &c)

	if err != nil {
		t.Error("Failed to parse config")
	}

	configStruct := reflect.TypeOf(receiver.UDP{})
	superfluousProperties := config.VerifyOnlyRequiredConfigProps(&c, "receivers", "foo", configStruct)

	if len(superfluousProperties) != 1 {
		t.Errorf("Expected 1 superfluous property but got %d", len(superfluousProperties))
	}

	if superfluousProperties[0] != "superfluousField" {
		t.Errorf("Expected to find '%s' in the superfluous fields list", "superfluousField")
	}
}
