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

	MethodQueueTop    Method = "queue-move-top"
	MethodQueueUp     Method = "queue-move-up"
	MethodQueueDown   Method = "queue-move-down"
	MethodQueueBottom Method = "queue-move-bottom"

	MethodFreeSpace       Method = "free-space"
	MethodPortTest        Method = "port-test"
	MethodBlockListUpdate Method = "blocklist-update"

	ResponseResultSuccess = "success"
	SessionIdHeader       = "X-Transmission-Session-Id"

	MaxRetries = 2
)

var (
	ErrInvalidSessionId = errors.New("invalid session-id header")
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
	MaxRetries int         `json:"-"`                   // used internally to avoid retries when token is invalid/expired
}

type TorrentAction struct {
	Ids interface{} `json:"ids"`
}

type QueueMovement struct {
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
	DownloadLimited     bool               `json:"downloadLimited,omitempty"`
	FilesWanted         []int64            `json:"files-wanted,omitempty"`
	FilesUnwanted       []int64            `json:"files-unwanted,omitempty"`
	HonorsSessionLimits bool               `json:"honorsSessionLimits,omitempty"`
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
	UploadLimited       bool               `json:"uploadLimited,omitempty"`
}

type SessionGet struct {
	Fields []string `json:"fields,omitempty"`
}

type SessionSet struct {
	AltSpeedDown              int64   `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin         int64   `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd           int64   `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int64   `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int64   `json:"alt-speed-up,omitempty"`
	BlockListUrl              string  `json:"blocklist-url,omitempty"`
	BlockListEnabled          bool    `json:"blocklist-enabled,omitempty"`
	CacheSizeMb               int64   `json:"cache-size-mb,omitempty"`
	DownloadDir               string  `json:"download-dir,omitempty"`
	DownloadQueueSize         int64   `json:"download-queue-size,omitempty"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled,omitempty"`
	DhtEnabled                bool    `json:"dht-enabled,omitempty"`
	Encryption                string  `json:"encryption,omitempty"`
	IdleSeedingLimit          int64   `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir             string  `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled,omitempty"`
	LpdEnabled                bool    `json:"lpd-enabled,omitempty"`
	PeerLimitGlobal           int64   `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int64   `json:"peer-limit-per-torrent,omitempty"`
	PexEnabled                bool    `json:"pex-enabled,omitempty"`
	PeerPort                  int64   `json:"peer-port,omitempty"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start,omitempty"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes       int64   `json:"queue-stalled-minutes,omitempty"`
	RenamePartialFiles        bool    `json:"rename-partial-files,omitempty"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled,omitempty"`
	SeedRatioLimit            float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited          bool    `json:"seedRatioLimited,omitempty"`
	SeedQueueSize             int64   `json:"seed-queue-size,omitempty"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled,omitempty"`
	SpeedLimitDown            int64   `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp              int64   `json:"speed-limit-up,omitempty"`
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
	SessionId  string
	HttpClient *http.Client
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

func WithHttpClient(client *http.Client) Option {
	return func(c *Client) {
		if client != nil {
			c.HttpClient = client
		} else {
			c.HttpClient = &http.Client{}
		}
	}
}

func New(opts ...Option) *Client {
	client := Client{HttpClient: &http.Client{}}
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

func (c *Client) doRequest(request *http.Request, maxRetries int) (*http.Response, error) {
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unexpected error sending http request: %+v", err)
	}
	if resp.StatusCode == http.StatusConflict {
		if maxRetries-1 <= 0 {
			return nil, ErrInvalidSessionId
		}
		return c.doRequest(request, maxRetries-1)
	}
	return resp, nil
}

func (c *Client) fetch(ctx context.Context, request request) (*response, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with context: %+v", err)
	}

	if c.Password != "" && c.Username != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	req.Header.Set("User-Agent", "transmission")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(SessionIdHeader, c.SessionId)

	maxRetries := request.MaxRetries
	if request.MaxRetries <= 0 {
		maxRetries = MaxRetries
	}
	resp, err := c.doRequest(req, maxRetries)
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
	_, err := c.fetch(ctx, request{Method: "ping", MaxRetries: 1})
	if errors.Is(err, ErrInvalidSessionId) {
		return nil
	}
	return err
}

func (c *Client) TorrentStart(ctx context.Context, args TorrentAction) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStart, Arguments: args})
	return err
}

func (c *Client) TorrentStartNow(ctx context.Context, args TorrentAction) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStartNow, Arguments: args})
	return err
}

func (c *Client) TorrentStop(ctx context.Context, args TorrentAction) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentStop, Arguments: args})
	return err
}

func (c *Client) TorrentVerify(ctx context.Context, args TorrentAction) error {
	_, err := c.fetch(ctx, request{Method: MethodTorrentVerify, Arguments: args})
	return err
}

func (c *Client) TorrentReannounce(ctx context.Context, args TorrentAction) error {
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
	return torrents, nil
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
	return torrent, nil
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

func (c *Client) QueueMoveTop(ctx context.Context, args QueueMovement) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueTop, Arguments: args})
	return err
}

func (c *Client) QueueMoveBottom(ctx context.Context, args QueueMovement) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueBottom, Arguments: args})
	return err
}

func (c *Client) QueueMoveUp(ctx context.Context, args QueueMovement) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueUp, Arguments: args})
	return err
}

func (c *Client) QueueMoveDown(ctx context.Context, args QueueMovement) error {
	_, err := c.fetch(ctx, request{Method: MethodQueueDown, Arguments: args})
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
