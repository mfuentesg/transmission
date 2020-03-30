package transmission

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Method string

const (
	MethodTorrentStart      Method = "torrent-start"
	MethodTorrentStartNow   Method = "torrent-start-now"
	MethodTorrentStop       Method = "torrent-stop"
	MethodTorrentVerify     Method = "torrent-verify"
	MethodTorrentReannounce Method = "torrent-reannounce"
	MethodTorrentSet        Method = "torrent-set"
	MethodTorrentGet        Method = "torrent-get"
	MethodTorrentAdd        Method = "torrent-add"
	MethodTorrentRemove     Method = "torrent-remove"
	MethodTorrentMove       Method = "torrent-set-location"
	MethodTorrentRename     Method = "torrent-rename-path"

	MethodSessionClose Method = "session-close"
	MethodSessionGet   Method = "session-get"
	MethodSessionSet   Method = "session-set"
	MethodSessionStats Method = "session-stats"

	MethodQueueMoveTop    Method = "queue-move-top"
	MethodQueueMoveUp     Method = "queue-move-up"
	MethodQueueMoveDown   Method = "queue-move-down"
	MethodQueueMoveBottom Method = "queue-move-bottom"

	MethodFreeSpace       Method = "free-space"
	MethodPortTest        Method = "port-test"
	MethodBlockListUpdate Method = "blocklist-update"

	ResponseResultSuccess = "success"
	SessionIDHeader       = "X-Transmission-Session-Id"

	DefaultMaxRetries = 2
	MaxRetries        = 10
)

var (
	ErrInvalidSessionID = errors.New("invalid session-id header")
)

type response struct {
	Result    string                 `json:"result,omitempty"`    // string whose value MUST be "success" on success, or an error string on failure
	Arguments map[string]interface{} `json:"arguments,omitempty"` // object of key/value pairs
	Tag       int64                  `json:"tag,omitempty"`       // number used by clients to track responses
}

type request struct {
	Method     Method      `json:"method,omitempty"`    // string telling the name of the method to invoke
	Arguments  interface{} `json:"arguments,omitempty"` // object of key/value pairs
	Tag        int64       `json:"tag,omitempty"`       // number used by clients to track responses (same request tag value)
	AvoidRetry bool        `json:"-"`
}

type Filter struct {
	Ids interface{} `json:"ids"`
}

type TorrentGet struct {
	Ids    interface{} `json:"ids,omitempty"`
	Fields []string    `json:"fields,omitempty"`
	// This fields supports only objects or table
	// ("table"  format always returns same as "objects" format)s
	// Format string   `json:"format,omitempty"`
}

type TorrentAdd struct {
	Cookies           string  `json:"cookies,omitempty"`
	DownloadDir       string  `json:"download-dir,omitempty"`
	Filename          string  `json:"filename,omitempty"`
	MetaInfo          string  `json:"metainfo,omitempty"`
	Paused            bool    `json:"paused,omitempty"`
	PeerLimit         int64   `json:"peer-limit"`
	BandwidthPriority int64   `json:"bandwidthPriority"`
	FilesWanted       []int64 `json:"files-wanted,omitempty"`
	FilesUnwanted     []int64 `json:"files-unwanted,omitempty"`
	PriorityHigh      []int64 `json:"priority-high,omitempty"`
	PriorityLow       []int64 `json:"priority-low,omitempty"`
	PriorityNormal    []int64 `json:"priority-normal,omitempty"`
}

type TorrentMove struct {
	Ids      interface{} `json:"ids,omitempty"`
	Location string      `json:"location,omitempty"`
	Move     bool        `json:"move,omitempty"`
}

type TorrentRemove struct {
	Ids             interface{} `json:"ids,omitempty"`
	DeleteLocalData bool        `json:"delete-local-data,omitempty"`
}

type TorrentRename struct {
	Ids  interface{} `json:"ids,omitempty"`
	Path string      `json:"path"` // represents current torrent name
	Name string      `json:"name"`
}

type TorrentSet struct {
	BandwidthPriority   int64              `json:"bandwidthPriority"`
	DownloadLimit       int64              `json:"downloadLimit"`
	FilesWanted         []int64            `json:"files-wanted,omitempty"`
	FilesUnwanted       []int64            `json:"files-unwanted,omitempty"`
	Ids                 interface{}        `json:"ids,omitempty"`
	Labels              []string           `json:"labels,omitempty"`
	Location            string             `json:"location,omitempty"`
	PeerLimit           int64              `json:"peer-limit"`
	PriorityHigh        []int64            `json:"priority-high,omitempty"`
	PriorityLow         []int64            `json:"priority-low,omitempty"`
	PriorityNormal      []int64            `json:"priority-normal,omitempty"`
	QueuePosition       int64              `json:"queuePosition"`
	SeedIdleLimit       int64              `json:"seedIdleLimit"`
	SeedIdleMode        int64              `json:"seedIdleMode"`
	SeedRatioLimit      float64            `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int64              `json:"seedRatioMode"`
	TrackerAdd          []string           `json:"trackerAdd,omitempty"`
	TrackerRemove       []int64            `json:"trackerRemove,omitempty"`
	TrackerReplace      []map[int64]string `json:"trackerReplace,omitempty"`
	UploadLimit         int64              `json:"uploadLimit"`
	DownloadLimited     bool               `json:"downloadLimited,omitempty"`
	HonorsSessionLimits bool               `json:"honorsSessionLimits,omitempty"`
	UploadLimited       bool               `json:"uploadLimited,omitempty"`
}

type SessionGet struct {
	Fields []string `json:"fields,omitempty"`
}

type SessionSet struct {
	AltSpeedDown              int64   `json:"alt-speed-down,omitempty"`
	AltSpeedTimeBegin         int64   `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnd           int64   `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int64   `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int64   `json:"alt-speed-up,omitempty"`
	BlockListURL              string  `json:"blocklist-url,omitempty"`
	CacheSizeMb               int64   `json:"cache-size-mb,omitempty"`
	DownloadDir               string  `json:"download-dir,omitempty"`
	DownloadQueueSize         int64   `json:"download-queue-size,omitempty"`
	Encryption                string  `json:"encryption,omitempty"`
	IdleSeedingLimit          int64   `json:"idle-seeding-limit,omitempty"`
	IncompleteDir             string  `json:"incomplete-dir,omitempty"`
	PeerLimitGlobal           int64   `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int64   `json:"peer-limit-per-torrent,omitempty"`
	PeerPort                  int64   `json:"peer-port,omitempty"`
	QueueStalledMinutes       int64   `json:"queue-stalled-minutes,omitempty"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename,omitempty"`
	SeedRatioLimit            float64 `json:"seedRatioLimit,omitempty"`
	SeedQueueSize             int64   `json:"seed-queue-size,omitempty"`
	SpeedLimitDown            int64   `json:"speed-limit-down,omitempty"`
	SpeedLimitUp              int64   `json:"speed-limit-up,omitempty"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled,omitempty"`
	BlockListEnabled          bool    `json:"blocklist-enabled,omitempty"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled,omitempty"`
	DhtEnabled                bool    `json:"dht-enabled,omitempty"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled,omitempty"`
	LpdEnabled                bool    `json:"lpd-enabled,omitempty"`
	PexEnabled                bool    `json:"pex-enabled,omitempty"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start,omitempty"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled,omitempty"`
	RenamePartialFiles        bool    `json:"rename-partial-files,omitempty"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled,omitempty"`
	SeedRatioLimited          bool    `json:"seedRatioLimited,omitempty"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled,omitempty"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled,omitempty"`
	StartAddedTorrents        bool    `json:"start-added-torrents,omitempty"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files,omitempty"`
	UtpEnabled                bool    `json:"utp-enabled,omitempty"`
}

type FreeSpace struct {
	Path      string `json:"path"`
	SizeBytes int64  `json:"size-bytes"`
}

type PortCheck struct {
	PortIsOpen bool `json:"port-is-open"`
}

type BlockList struct {
	BlockListSize int64 `json:"blocklist-size"`
}

type Option func(*Client)

type Client struct {
	Username   string
	Password   string
	URL        string
	SessionID  string
	HTTPClient *http.Client
	MaxRetries int
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.URL = url
	}
}

func WithBasicAuth(user, password string) Option {
	return func(c *Client) {
		c.Username = user
		c.Password = password
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		if client != nil {
			c.HTTPClient = client
		} else {
			c.HTTPClient = &http.Client{}
		}
	}
}

func WithMaxRetries(maxRetries int) Option {
	return func(c *Client) {
		switch {
		case maxRetries <= 0:
			c.MaxRetries = DefaultMaxRetries
		case maxRetries > MaxRetries:
			// avoid infinite retries
			c.MaxRetries = MaxRetries
		default:
			c.MaxRetries = maxRetries
		}
	}
}

func New(opts ...Option) *Client {
	client := Client{HTTPClient: &http.Client{}, MaxRetries: DefaultMaxRetries}

	for _, o := range opts {
		o(&client)
	}

	return &client
}

func fillStruct(base interface{}, target interface{}) error {
	buf, err := json.Marshal(base)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, target)
}

func (c *Client) doRequest(ctx context.Context, body []byte, maxRetries int) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with context: %+v", err)
	}

	if c.Password != "" && c.Username != "" {
		request.SetBasicAuth(c.Username, c.Password)
	}

	request.Header.Set("User-Agent", "transmission")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(SessionIDHeader, c.SessionID)
	resp, err := c.HTTPClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("unexpected error sending http request: %+v", err)
	}

	if resp.StatusCode == http.StatusConflict {
		c.SessionID = resp.Header.Get(SessionIDHeader)

		if maxRetries-1 <= 0 {
			return nil, ErrInvalidSessionID
		}

		return c.doRequest(ctx, body, maxRetries-1)
	}

	return resp, nil
}

func (c *Client) fetch(ctx context.Context, request request) (*response, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	// overwrite client maxRetries value
	maxRetries := c.MaxRetries
	if request.AvoidRetry {
		maxRetries = 1
	}

	resp, err := c.doRequest(ctx, body, maxRetries)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res response
	if err := json.Unmarshal(buf, &res); err != nil {
		return nil, err
	}

	if res.Result != ResponseResultSuccess {
		return nil, fmt.Errorf("unexpected result: %s", res.Result)
	}

	return &res, nil
}

func (c *Client) Ping(ctx context.Context) error {
	// this is just a hack to retrieve a valid session id token
	_, err := c.fetch(ctx, request{Method: "ping", AvoidRetry: true})
	if errors.Is(err, ErrInvalidSessionID) {
		return nil
	}

	return err
}

func (c *Client) TorrentStart(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStart, Arguments: args})
	return err
}

func (c *Client) TorrentStartNow(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStartNow, Arguments: args})
	return err
}

func (c *Client) TorrentStop(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStop, Arguments: args})
	return err
}

func (c *Client) TorrentVerify(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentVerify, Arguments: args})
	return err
}

func (c *Client) TorrentReannounce(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentReannounce, Arguments: args})
	return err
}

func (c *Client) TorrentGet(ctx context.Context, args TorrentGet) ([]Torrent, error) {
	var torrents []Torrent

	response, err := c.fetch(ctx, request{Method: MethodTorrentGet, Arguments: args})

	if err != nil {
		return torrents, err
	}

	list, ok := response.Arguments["torrents"]
	if !ok {
		return torrents, nil
	}

	err = fillStruct(list, &torrents)

	return torrents, err
}

func (c *Client) TorrentRename(ctx context.Context, args TorrentRename) (Torrent, error) {
	var torrent Torrent

	resp, err := c.fetch(ctx, request{Method: MethodTorrentRename, Arguments: args})

	if err != nil {
		return torrent, err
	}

	err = fillStruct(resp.Arguments, &torrent)

	return torrent, err
}

func (c *Client) TorrentSet(ctx context.Context, args TorrentSet) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentSet, Arguments: args})
	return err
}

func (c *Client) TorrentAdd(ctx context.Context, args TorrentAdd) (Torrent, error) {
	var torrent Torrent

	resp, err := c.fetch(ctx, request{Method: MethodTorrentAdd, Arguments: args})
	if err != nil {
		return torrent, err
	}

	added, ok := resp.Arguments["torrent-added"]
	if !ok {
		return torrent, nil
	}

	err = fillStruct(added, &torrent)

	return torrent, err
}

func (c *Client) TorrentRemove(ctx context.Context, args TorrentRemove) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentRemove, Arguments: args})
	return err
}

func (c *Client) TorrentMove(ctx context.Context, args TorrentMove) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentMove, Arguments: args})
	return err
}

func (c *Client) SessionSet(ctx context.Context, args SessionSet) error {
	_, err := c.fetch(ctx, request{Method: MethodSessionSet, Arguments: args})
	return err
}

func (c *Client) SessionGet(ctx context.Context, args SessionGet) (Session, error) {
	var session Session

	resp, err := c.fetch(ctx, request{Method: MethodSessionGet, Arguments: args})
	if err != nil {
		return session, err
	}

	err = fillStruct(resp.Arguments, &session)

	return session, err
}

func (c *Client) SessionStats(ctx context.Context) (SessionStats, error) {
	var stats SessionStats

	resp, err := c.fetch(ctx, request{Method: MethodSessionStats})
	if err != nil {
		return stats, err
	}

	err = fillStruct(resp.Arguments, &stats)

	return stats, err
}

func (c *Client) SessionClose(ctx context.Context) error {
	_, err := c.fetch(ctx, request{Method: MethodSessionClose})
	return err
}

func (c *Client) QueueMoveTop(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueMoveTop, Arguments: args})
	return err
}

func (c *Client) QueueMoveBottom(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueMoveBottom, Arguments: args})
	return err
}

func (c *Client) QueueMoveUp(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueMoveUp, Arguments: args})
	return err
}

func (c *Client) QueueMoveDown(ctx context.Context, args Filter) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueMoveDown, Arguments: args})
	return err
}

func (c *Client) FreeSpace(ctx context.Context, args FreeSpace) (FreeSpace, error) {
	var free FreeSpace

	resp, err := c.fetch(ctx, request{Method: MethodFreeSpace, Arguments: args})
	if err != nil {
		return free, err
	}

	err = fillStruct(resp.Arguments, &free)

	return free, err
}

func (c *Client) PortCheck(ctx context.Context) (PortCheck, error) {
	var port PortCheck

	resp, err := c.fetch(ctx, request{Method: MethodPortTest})
	if err != nil {
		return port, err
	}

	err = fillStruct(resp.Arguments, &port)

	return port, err
}

func (c *Client) BlockListUpdate(ctx context.Context) (BlockList, error) {
	var blockList BlockList

	resp, err := c.fetch(ctx, request{Method: MethodBlockListUpdate})
	if err != nil {
		return blockList, err
	}

	err = fillStruct(resp.Arguments, &blockList)

	return blockList, err
}
