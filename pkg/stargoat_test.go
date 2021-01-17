package stargoat

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	faker "github.com/bxcodec/faker/v3"

	"github.com/loov/hrtime/hrtesting"
)

func TestAuthAndRoundtrip(t *testing.T) {
	keyspace, ok := os.LookupEnv("ASTRA_DB_KEYSPACE")
	if !ok {
		t.Log("Missing ASTRA_DB_KEYSPACE env needed for test run")
		t.FailNow()
	}
	namespace := keyspace

	astraDatabaseBaseURL, ok := os.LookupEnv("ASTRA_DB_BASE_URL")
	if !ok {
		t.Log("Missing ASTRA_DB_BASE_URL env needed for test run")
		t.FailNow()
	}

	stargateURL := astraDatabaseBaseURL + "/api/rest"
	authURL := astraDatabaseBaseURL + "/api/rest/v1/auth"

	getToken := func() StargateToken {
		creds := &EnvVarCredsProvider{
			UsernameEnv: "ASTRA_DB_USERNAME",
			PasswordEnv: "ASTRA_DB_PASSWORD",
		}
		got, err := Authenticate(creds, authURL)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		return *got
	}
	t.Run("test auth and roundtrip", func(t *testing.T) {
		client, err := NewClient(stargateURL, getToken(), true, nil)
		assert.NoError(t, err)

		type MyDoc struct {
			First string
			Last  string
			Age   int
		}

		doc := MyDoc{"Jason", "Dickinson", 6}
		id, err := client.PostDoc(namespace, "abc", doc)
		assert.NoError(t, err)
		assert.Len(t, id, 36)

		otherID, err := client.PutDoc(namespace, "abc", "11112222", doc)
		assert.NoError(t, err)
		assert.NotEqual(t, id, otherID)
		assert.Equal(t, "11112222", otherID)

		doc.Age = 7
		id2, err := client.PatchDoc(namespace, "abc", id, doc)
		assert.NoError(t, err)
		assert.Equal(t, id, id2)
		assert.Len(t, id2, 36)

		newDoc, err := client.GetDoc(namespace, "abc", id)
		assert.NoError(t, err)
		JSONEq(t, doc, newDoc)

		newDoc2, err := client.GetDoc(namespace, "abc", otherID)
		assert.NoError(t, err)
		assert.NotEqual(t, newDoc, newDoc2)

		ok, err := client.DeleteDoc(namespace, "abc", faker.UUIDHyphenated())
		assert.NoError(t, err)
		assert.True(t, ok)

		ok, err = client.DeleteDoc(namespace, "abc", otherID)
		assert.NoError(t, err)
		assert.True(t, ok)

		ok, err = client.DeleteDoc(namespace, faker.Word()+faker.Word(), otherID)
		assert.Error(t, err, "DeleteDoc on a collection that does not exist should error")
		assert.False(t, ok)
	})
}

func JSONEq(t *testing.T, b1, b2 interface{}) bool {
	j1, err := json.Marshal(b1)
	assert.NoError(t, err)
	j2, err := json.Marshal(b1)
	assert.NoError(t, err)
	return assert.JSONEq(t, string(j1), string(j2))
}

func BenchmarkParallelPost(b *testing.B) {
	bench := hrtesting.NewBenchmark(b)
	defer bench.Report()

	keyspace, ok := os.LookupEnv("ASTRA_DB_KEYSPACE")
	if !ok {
		b.Log("Missing ASTRA_DB_KEYSPACE env needed for test run")
		b.FailNow()
	}
	namespace := keyspace

	astraDatabaseBaseURL, ok := os.LookupEnv("ASTRA_DB_BASE_URL")
	if !ok {
		b.Log("Missing ASTRA_DB_BASE_URL env needed for test run")
		b.FailNow()
	}

	stargateURL := astraDatabaseBaseURL + "/api/rest"
	authURL := astraDatabaseBaseURL + "/api/rest/v1/auth"

	getToken := func() StargateToken {
		creds := &EnvVarCredsProvider{
			UsernameEnv: "ASTRA_DB_USERNAME",
			PasswordEnv: "ASTRA_DB_PASSWORD",
		}
		got, _ := Authenticate(creds, authURL)
		return *got
	}
	client, _ := NewClient(stargateURL, getToken(), false, nil)

	type GitHubProject struct {
		Org         string `faker:"username"`
		Name        string `faker:"word"`
		Stars       int    `faker:"boundary_start=1, boundary_end=10000"`
		LastChange  int64  `faker:"unix_time"`
		Description string `faker:"paragraph"`
	}

	projects := []GitHubProject{}
	for i := 0; i < 2*b.N; i++ {
		prj := GitHubProject{}
		err := faker.FakeData(&prj)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, prj)
	}
	numProjects := len(projects)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for bench.Next() {
				doc := projects[rand.Intn(numProjects)]
				client.PostDoc(namespace, "mycollection", doc)
			}
		}
	})
}
