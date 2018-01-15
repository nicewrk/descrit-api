package dotenv

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	testVars := map[string]string{
		"DB_HOST": "guru-staging.abc123.us-west-2.rds.amazonaws.com",
		"DB_NAME": "guru-staging",
		"DB_PASS": "notArealPa55word",
		"DB_PORT": "5432",
		"DB_USER": "guru-staging",
	}

	tmpfilepath, err := setUp(testVars)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Remove(tmpfilepath)
		if err != nil {
			t.Fatalf("error: removing file %s: %s", tmpfilepath, err)
		}
	}()

	// Pre-Run() check empty env var values,
	// otherwise post-Run() tests will be unreliable.
	for k, v := range testVars {
		t.Run(fmt.Sprintf("pre-Run(): k=%s, v=%s pair", k, v), func(t *testing.T) {
			if os.Getenv(k) == v {
				t.Errorf(`got "%s"; want ""`, v)
			}
		})
	}

	err = Run(tmpfilepath)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range testVars {
		t.Run(fmt.Sprintf("post-Run(): k=%s, v=%s pair", k, v), func(t *testing.T) {
			if os.Getenv(k) != v {
				t.Errorf(`got "%s"; want "%v"`, os.Getenv(k), v)
			}
		})
	}
}

func setUp(envVars map[string]string) (string, error) {
	filename := ".env-test"

	var buf bytes.Buffer
	for k, v := range envVars {
		buf.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}

	tmpfile, err := ioutil.TempFile("", filename)
	if err != nil {
		return "", err
	}

	_, err = tmpfile.Write(buf.Bytes())
	if err != nil {
		return "", err
	}

	err = tmpfile.Close()
	if err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}
