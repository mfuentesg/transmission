package transmission

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithURL(t *testing.T) {
	var client Client
	fakeURL := "http://fake.com/tranmission/rpc"
	WithURL(fakeURL)(&client)

	assert.Equal(t, fakeURL, client.URL)
}

func TestWithBasicAuth(t *testing.T) {
	var client Client
	WithBasicAuth("user", "password")(&client)

	assert.Equal(t, "user", client.Username)
	assert.Equal(t, "password", client.Password)
}

func TestWithHttpClient(t *testing.T) {
	var client Client

	t.Run("should assign set default http client with an invalid parameter", func(st *testing.T) {
		WithHttpClient(nil)(&client)
		assert.NotNil(st, client.HttpClient)
	})

	t.Run("should assign valid http client", func(st *testing.T) {
		httpClient := &http.Client{
			Timeout: 15000,
		}
		WithHttpClient(httpClient)(&client)
		assert.Equal(st, httpClient.Timeout, client.HttpClient.Timeout)
	})
}

func TestNew(t *testing.T) {
	fakeURL := "http://fake.com/tranmission/rpc"
	cl := New(WithURL(fakeURL))
	assert.NotNil(t, cl)
	assert.IsType(t, &Client{}, cl)
	assert.Equal(t, fakeURL, cl.URL)
}

func TestFillStruct(t *testing.T) {
	type user struct {
		Name string `json:"name"`
	}
	tests := []struct {
		name              string
		base              interface{}
		target            *user
		isErrorExpected   bool
		shouldCheckTarget bool
	}{
		{
			name:            "should get an error for marshal",
			isErrorExpected: true,
			base:            make(chan int),
		},
		{
			name:            "should get an error for unmarshal with nil, nil",
			isErrorExpected: true,
		},
		{
			name:            "should get an error for unmarshal with target = nil",
			isErrorExpected: true,
			base:            map[string]interface{}{"name": "tranmission"},
		},
		{
			name:              "should return a valid struct",
			base:              map[string]interface{}{"name": "transmission"},
			shouldCheckTarget: true,
			target:            &user{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			err := fillStruct(test.base, test.target)

			if test.isErrorExpected {
				assert.NotNil(st, err)
				assert.Error(st, err)
			} else {
				assert.Nil(st, err)
				assert.NoError(st, err)
			}

			if test.shouldCheckTarget {
				assert.Equal(st, "transmission", test.target.Name)
			}
		})
	}
}

func BenchmarkFillStruct(b *testing.B) {
	data := map[string]interface{}{
		"arguments": map[string]interface{}{
			"torrents": []interface{}{
				map[string]interface{}{
					"activityDate":      1585100850,
					"addedDate":         1585100763,
					"bandwidthPriority": 0,
					"comment":           "",
					"corruptEver":       0,
					"creator":           "",
					"dateCreated":       0,
					"desiredAvailable":  0,
					"doneDate":          0,
					"downloadDir":       "/media/storage/Downloads",
					"downloadLimit":     100,
					"downloadLimited":   false,
					"downloadedEver":    25924291,
					"error":             0,
					"errorString":       "",
					"eta":               -1,
					"etaIdle":           -1,
					"fileStats": []interface{}{
						map[string]interface{}{
							"bytesCompleted": 2291211,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 5262451,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 4334885,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 2392312,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 239395,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 2161788,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 4922076,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 4264672,
							"priority":       0,
							"wanted":         true,
						},
						map[string]interface{}{
							"bytesCompleted": 11757,
							"priority":       0,
							"wanted":         true,
						},
					},
					"files": []interface{}{
						map[string]interface{}{
							"bytesCompleted": 2291211,
							"length":         226686475,
							"name":           "Greys anatomy/Grey's Anatomy S01E01 A Hard Day's Night.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 5262451,
							"length":         217074803,
							"name":           "Greys anatomy/Grey's Anatomy S01E05 Shake Your Groove Thing.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 4334885,
							"length":         211723557,
							"name":           "Greys anatomy/Grey's Anatomy S01E03 Winning a Battle, Losing the War.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 2392312,
							"length":         208388344,
							"name":           "Greys anatomy/Grey's Anatomy S01E07 The Self-Destruct Button.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 239395,
							"length":         207857443,
							"name":           "Greys anatomy/Grey's Anatomy S01E08 Save Me.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 2161788,
							"length":         205585532,
							"name":           "Greys anatomy/Grey's Anatomy S01E02 The First Cut is the Deepest.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 4922076,
							"length":         202513116,
							"name":           "Greys anatomy/Grey's Anatomy S01E04 No Man's Land.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 4264672,
							"length":         202085088,
							"name":           "Greys anatomy/Grey's Anatomy S01E06 If Tomorrow Never Comes.mkv",
						},
						map[string]interface{}{
							"bytesCompleted": 11757,
							"length":         190852589,
							"name":           "Greys anatomy/Grey's Anatomy S01E09 Who's Zoomin' Who.mkv",
						},
					},
					"hashString":              "f2599a954d5acb8a06371e3b32b4c5f46c55376c",
					"haveUnchecked":           4898816,
					"haveValid":               20981731,
					"honorsSessionLimits":     true,
					"id":                      2,
					"isFinished":              false,
					"isPrivate":               false,
					"isStalled":               false,
					"leftUntilDone":           1846886400,
					"magnetLink":              "magnet:?xt=urn:btih:f2599a954d5acb8a06371e3b32b4c5f46c55376c&dn=Greys%20anatomy&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80&tr=udp%3A%2F%2Fopen.demonii.com%3A1337&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Fexodus.desync.com%3A6969",
					"manualAnnounceTime":      -1,
					"maxConnectedPeers":       50,
					"metadataPercentComplete": 1,
					"name":                    "Greys anatomy",
					"peer-limit":              50,
					"peers":                   []interface{}{},
					"peersConnected":          0,
					"peersFrom": map[string]interface{}{
						"fromCache":    0,
						"fromDht":      0,
						"fromIncoming": 0,
						"fromLpd":      0,
						"fromLtep":     0,
						"fromPex":      0,
						"fromTracker":  0,
					},
					"peersGettingFromUs": 0,
					"peersSendingToUs":   0,
					"percentDone":        0.0138,
					"pieceCount":         894,
					"pieceSize":          2097152,
					"pieces":             "gAAAAAAAAAAAAAAAAAgAAAAAAAgAAAAAAAARAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAAAAABA==",
					"priorities": []interface{}{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					"queuePosition":      0,
					"rateDownload":       0,
					"rateUpload":         0,
					"recheckProgress":    0,
					"secondsDownloading": 56,
					"secondsSeeding":     0,
					"seedIdleLimit":      15,
					"seedIdleMode":       0,
					"seedRatioLimit":     0,
					"seedRatioMode":      0,
					"sizeWhenDone":       1872766947,
					"startDate":          1585100806,
					"status":             0,
					"torrentFile":        "/etc/transmission-daemon/torrents/Grey's Anatomy Season 1 Complete HDTV x264 [i_c].f2599a954d5acb8a.torrent",
					"totalSize":          1872766947,
					"trackerStats": []interface{}{
						map[string]interface{}{
							"announce":              "udp://tracker.leechers-paradise.org:6969",
							"announceState":         0,
							"downloadCount":         997,
							"hasAnnounced":          true,
							"hasScraped":            true,
							"host":                  "udp://tracker.leechers-paradise.org:6969",
							"id":                    0,
							"isBackup":              false,
							"lastAnnouncePeerCount": 0,
							"lastAnnounceResult":    "Success",
							"lastAnnounceStartTime": 0,
							"lastAnnounceSucceeded": true,
							"lastAnnounceTime":      1585100859,
							"lastAnnounceTimedOut":  false,
							"lastScrapeResult":      "Could not connect to tracker",
							"lastScrapeStartTime":   1585415390,
							"lastScrapeSucceeded":   true,
							"lastScrapeTime":        1585415391,
							"lastScrapeTimedOut":    0,
							"leecherCount":          22,
							"nextAnnounceTime":      0,
							"nextScrapeTime":        1585417200,
							"scrape":                "udp://tracker.leechers-paradise.org:6969",
							"scrapeState":           1,
							"seederCount":           84,
							"tier":                  0,
						},
						map[string]interface{}{
							"announce":              "udp://tracker.openbittorrent.com:80",
							"announceState":         0,
							"downloadCount":         -1,
							"hasAnnounced":          true,
							"hasScraped":            true,
							"host":                  "udp://tracker.openbittorrent.com:80",
							"id":                    1,
							"isBackup":              false,
							"lastAnnouncePeerCount": 0,
							"lastAnnounceResult":    "Connection failed",
							"lastAnnounceStartTime": 0,
							"lastAnnounceSucceeded": false,
							"lastAnnounceTime":      1585411710,
							"lastAnnounceTimedOut":  false,
							"lastScrapeResult":      "Connection failed",
							"lastScrapeStartTime":   0,
							"lastScrapeSucceeded":   false,
							"lastScrapeTime":        1585409515,
							"lastScrapeTimedOut":    0,
							"leecherCount":          -1,
							"nextAnnounceTime":      0,
							"nextScrapeTime":        1585416770,
							"scrape":                "udp://tracker.openbittorrent.com:80",
							"scrapeState":           1,
							"seederCount":           -1,
							"tier":                  1,
						},
						map[string]interface{}{
							"announce":              "udp://open.demonii.com:1337",
							"announceState":         0,
							"downloadCount":         -1,
							"hasAnnounced":          true,
							"hasScraped":            true,
							"host":                  "udp://open.demonii.com:1337",
							"id":                    2,
							"isBackup":              false,
							"lastAnnouncePeerCount": 0,
							"lastAnnounceResult":    "Connection failed",
							"lastAnnounceStartTime": 0,
							"lastAnnounceSucceeded": false,
							"lastAnnounceTime":      1585412045,
							"lastAnnounceTimedOut":  false,
							"lastScrapeResult":      "Connection failed",
							"lastScrapeStartTime":   0,
							"lastScrapeSucceeded":   false,
							"lastScrapeTime":        1585409805,
							"lastScrapeTimedOut":    0,
							"leecherCount":          -1,
							"nextAnnounceTime":      0,
							"nextScrapeTime":        1585417010,
							"scrape":                "udp://open.demonii.com:1337",
							"scrapeState":           1,
							"seederCount":           -1,
							"tier":                  2,
						},
						map[string]interface{}{
							"announce":              "udp://tracker.coppersurfer.tk:6969",
							"announceState":         0,
							"downloadCount":         238,
							"hasAnnounced":          true,
							"hasScraped":            true,
							"host":                  "udp://tracker.coppersurfer.tk:6969",
							"id":                    3,
							"isBackup":              false,
							"lastAnnouncePeerCount": 0,
							"lastAnnounceResult":    "Success",
							"lastAnnounceStartTime": 0,
							"lastAnnounceSucceeded": true,
							"lastAnnounceTime":      1585101903,
							"lastAnnounceTimedOut":  false,
							"lastScrapeResult":      "Connection failed",
							"lastScrapeStartTime":   0,
							"lastScrapeSucceeded":   false,
							"lastScrapeTime":        1585413431,
							"lastScrapeTimedOut":    0,
							"leecherCount":          22,
							"nextAnnounceTime":      0,
							"nextScrapeTime":        1585420640,
							"scrape":                "udp://tracker.coppersurfer.tk:6969",
							"scrapeState":           1,
							"seederCount":           67,
							"tier":                  3,
						},
						map[string]interface{}{
							"announce":              "udp://exodus.desync.com:6969",
							"announceState":         0,
							"downloadCount":         1801,
							"hasAnnounced":          true,
							"hasScraped":            true,
							"host":                  "udp://exodus.desync.com:6969",
							"id":                    4,
							"isBackup":              false,
							"lastAnnouncePeerCount": 0,
							"lastAnnounceResult":    "Success",
							"lastAnnounceStartTime": 0,
							"lastAnnounceSucceeded": true,
							"lastAnnounceTime":      1585100859,
							"lastAnnounceTimedOut":  false,
							"lastScrapeResult":      "Connection failed",
							"lastScrapeStartTime":   1585416160,
							"lastScrapeSucceeded":   true,
							"lastScrapeTime":        1585416160,
							"lastScrapeTimedOut":    0,
							"leecherCount":          6,
							"nextAnnounceTime":      0,
							"nextScrapeTime":        1585417960,
							"scrape":                "udp://exodus.desync.com:6969",
							"scrapeState":           1,
							"seederCount":           36,
							"tier":                  4,
						},
					},
					"trackers": []interface{}{
						map[string]interface{}{
							"announce": "udp://tracker.leechers-paradise.org:6969",
							"id":       0,
							"scrape":   "udp://tracker.leechers-paradise.org:6969",
							"tier":     0,
						},
						map[string]interface{}{
							"announce": "udp://tracker.openbittorrent.com:80",
							"id":       1,
							"scrape":   "udp://tracker.openbittorrent.com:80",
							"tier":     1,
						},
						map[string]interface{}{
							"announce": "udp://open.demonii.com:1337",
							"id":       2,
							"scrape":   "udp://open.demonii.com:1337",
							"tier":     2,
						},
						map[string]interface{}{
							"announce": "udp://tracker.coppersurfer.tk:6969",
							"id":       3,
							"scrape":   "udp://tracker.coppersurfer.tk:6969",
							"tier":     3,
						},
						map[string]interface{}{
							"announce": "udp://exodus.desync.com:6969",
							"id":       4,
							"scrape":   "udp://exodus.desync.com:6969",
							"tier":     4,
						},
					},
					"uploadLimit":   100,
					"uploadLimited": false,
					"uploadRatio":   0,
					"uploadedEver":  0,
					"wanted": []interface{}{
						1,
						1,
						1,
						1,
						1,
						1,
						1,
						1,
						1,
					},
					"webseeds":            []interface{}{},
					"webseedsSendingToUs": 0,
				},
			},
		},
		"result": "success",
	}

	for i := 0; i < b.N; i++ {
		var resp response
		_ = fillStruct(data, &resp)
	}
}

func TestFetch(t *testing.T) {
	t.Run("should return an error for invalid marshal", func(st *testing.T) {
		client := New()
		res, err := client.fetch(context.Background(), request{
			Arguments: make(chan bool),
		})
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("should return an error with context nil", func(st *testing.T) {
		client := New()
		res, err := client.fetch(nil, request{
			Arguments: map[string]interface{}{},
		})
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("should return an error trying to execute request without valid url", func(st *testing.T) {
		client := New()
		res, err := client.fetch(context.Background(), request{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("should return an error an invalid body response", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`<>`))
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(WithURL(s.URL), WithHttpClient(s.Client()))
		res, err := client.fetch(context.Background(), request{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("should not add authorization header if username is empty", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.Header.Get("authorization"), "")
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(
			WithURL(s.URL),
			WithHttpClient(s.Client()),
			WithBasicAuth("", "secret"),
		)
		_, _ = client.fetch(context.Background(), request{})
	})

	t.Run("should not add authorization header if password is empty", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.Header.Get("authorization"), "")
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(
			WithURL(s.URL),
			WithHttpClient(s.Client()),
			WithBasicAuth("username", ""),
		)
		_, _ = client.fetch(context.Background(), request{})
	})

	t.Run("should not add authorization header if username and password are empty", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.Header.Get("authorization"), "")
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(
			WithURL(s.URL),
			WithHttpClient(s.Client()),
		)
		_, _ = client.fetch(context.Background(), request{})
	})

	t.Run("should add authorization header if username and password are not empty", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			assert.NotEqual(t, r.Header.Get("authorization"), "")
			auth := base64.StdEncoding.EncodeToString([]byte("username:secret"))
			assert.Equal(t, r.Header.Get("authorization"), "Basic "+auth)
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(
			WithURL(s.URL),
			WithHttpClient(s.Client()),
			WithBasicAuth("username", "secret"),
		)
		_, _ = client.fetch(context.Background(), request{})
	})

	t.Run("should return an error with result property different to success", func(st *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result": "unknown"}`))
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		client := New(WithURL(s.URL), WithHttpClient(s.Client()))
		res, err := client.fetch(context.Background(), request{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("should retry the same quantity of maxRetries property when http code is 409", func(st *testing.T) {
		tests := []struct {
			expected int
			input    int
		}{
			{input: 2, expected: 2},
			{input: 0, expected: 2},
			{input: 1, expected: 1},
			{input: 20, expected: 20},
			{input: -1, expected: 2},
		}

		for _, test := range tests {
			retries := new(int)
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				*retries = *retries + 1
				w.WriteHeader(http.StatusConflict)
			}))
			client := New(WithURL(s.URL), WithHttpClient(s.Client()))
			res, err := client.fetch(context.Background(), request{MaxRetries: test.input})

			assert.Nil(t, res)
			assert.NotNil(t, err)
			assert.Error(t, err)
			assert.Equal(t, test.expected, *retries)
		}
	})

	t.Run("should return a valid response", func(st *testing.T) {})
}

func BenchmarkFetch(t *testing.B) {}

func TestClient_Ping(t *testing.T) {
	tests := []struct {
		name            string
		response        []byte
		statusCode      int
		isErrorExpected bool
	}{
		{
			name:            "should get an error with unknown result",
			statusCode:      http.StatusOK,
			response:        []byte(`{"result": "unknown error"}`),
			isErrorExpected: true,
		},
		{
			name:       "should get valid response with HTTP Code 409",
			statusCode: http.StatusConflict,
			response:   []byte(`{"result": "unknown error"}`),
		},
		{
			name:       "should get an error with `unknown result`",
			statusCode: http.StatusConflict,
		},
		{
			name:       "should get valid response with HTTP Code 200",
			statusCode: http.StatusOK,
			response:   []byte(`{"result": "success"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				_, _ = w.Write(test.response)
			}
			s := httptest.NewServer(http.HandlerFunc(handler))
			client := New(
				WithURL(s.URL),
				WithBasicAuth("username", "password"),
				WithHttpClient(s.Client()),
			)
			err := client.Ping(context.Background())
			if test.isErrorExpected {
				assert.NotNil(st, err)
				assert.Error(st, err)
			} else {
				assert.Nil(st, err)
				assert.NoError(st, err)
			}
		})
	}
}

func TestClient_TorrentStart(t *testing.T) {}

func TestClient_TorrentStartNow(t *testing.T) {}

func TestClient_TorrentStop(t *testing.T) {}

func TestClient_TorrentVerify(t *testing.T) {}

func TestClient_TorrentReannounce(t *testing.T) {}

func TestClient_TorrentGet(t *testing.T) {}

func TestClient_TorrentRename(t *testing.T) {}

func TestClient_TorrentSet(t *testing.T) {}

func TestClient_TorrentAdd(t *testing.T) {}

func TestClient_TorrentRemove(t *testing.T) {}

func TestClient_TorrentMove(t *testing.T) {}

func TestClient_SessionSet(t *testing.T) {}

func TestClient_SessionGet(t *testing.T) {}

func TestClient_SessionStats(t *testing.T) {}

func TestClient_SessionClose(t *testing.T) {}

func TestClient_QueueMoveTop(t *testing.T) {}

func TestClient_QueueMoveBottom(t *testing.T) {}

func TestClient_QueueMoveUp(t *testing.T) {}

func TestClient_QueueMoveDown(t *testing.T) {}

func TestClient_FreeSpace(t *testing.T) {}

func TestClient_PortCheck(t *testing.T) {}

func TestClient_BlockListUpdate(t *testing.T) {}
