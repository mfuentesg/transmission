package transmission

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Method string

const (
	// actions
	MethodStart      Method = "torrent-start"
	MethodStartNow   Method = "torrent-start-now"
	MethodStop       Method = "torrent-stop"
	MethodVerify     Method = "torrent-verify"
	MethodReannounce Method = "torrent-reannounce"

	// mutators
	MethodSet Method = "torrent-set"

	// accessors
	MethodGet Method = "torrent-get"

	MethodAdd    Method = "torrent-add"
	MethodRemove Method = "torrent-remove"
	MethodMove   Method = "torrent-set-location"
	MethodRename Method = "torrent-rename-path"

	ResponseResultError = "error"

	SessionIdHeader = "X-Transmission-Session-Id"
)

type File struct {
	BytesCompleted int64  `json:"bytesCompleted,omitempty"`
	Length         int64  `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
}

type FileStat struct {
	BytesCompleted int64 `json:"bytesCompleted,omitempty"`
	Wanted         bool  `json:"wanted,omitempty"`
	Priority       int64 `json:"priority,omitempty"`
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
	ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
	FlagStr            string  `json:"flagStr,omitempty"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
	IsEncrypted        bool    `json:"isEncrypted,omitempty"`
	IsIncoming         bool    `json:"isIncoming,omitempty"`
	IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
	IsUTP              bool    `json:"isUTP,omitempty"`
	PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
	PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
	Port               int64   `json:"port,omitempty"`
	Progress           float64 `json:"progress,omitempty"`
	RateToClient       int64   `json:"rateToClient,omitempty"`
	RateToPeer         int64   `json:"rateToPeer,omitempty"`
}

type PeerFrom struct {
	FromCache    int64 `json:"fromCache,omitempty"`
	FromDht      int64 `json:"fromDht,omitempty"`
	FromIncoming int64 `json:"fromIncoming,omitempty"`
	FromLpd      int64 `json:"fromLpd,omitempty"`
	FromLtep     int64 `json:"fromLtep,omitempty"`
	FromPex      int64 `json:"fromPex,omitempty"`
	FromTracker  int64 `json:"fromTracker,omitempty"`
}

type Tracker struct {
	Announce string `json:"announce,omitempty"`
	Id       int64  `json:"id,omitempty"`
	Scrape   string `json:"scrape,omitempty"`
	Tier     int64  `json:"tier,omitempty"`
}

type TrackerStat struct {
	Announce              string `json:"announce,omitempty"`
	AnnounceState         int64  `json:"announceState,omitempty"`
	DownloadCount         int64  `json:"downloadCount,omitempty"`
	HasAnnounced          bool   `json:"hasAnnounced,omitempty"`
	HasScraped            bool   `json:"hasScraped,omitempty"`
	Host                  string `json:"host,omitempty"`
	Id                    int64  `json:"id,omitempty"`
	IsBackup              bool   `json:"isBackup,omitempty"`
	LastAnnouncePeerCount int64  `json:"lastAnnouncePeerCount,omitempty"`
	LastAnnounceResult    string `json:"lastAnnounceResult,omitempty"`
	LastAnnounceStartTime int64  `json:"lastAnnounceStartTime,omitempty"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded,omitempty"`
	LastAnnounceTime      int64  `json:"lastAnnounceTime,omitempty"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut,omitempty"`
	LastScrapeResult      string `json:"lastScrapeResult,omitempty"`
	LastScrapeStartTime   int64  `json:"lastScrapeStartTime,omitempty"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded,omitempty"`
	LastScrapeTime        int64  `json:"lastScrapeTime,omitempty"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut,omitempty"`
	LeecherCount          int64  `json:"leecherCount,omitempty"`
	NextAnnounceTime      int64  `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime        int64  `json:"nextScrapeTime,omitempty"`
	Scrape                string `json:"scrape,omitempty"`
	ScrapeState           int64  `json:"scrapeState,omitempty"`
	SeederCount           int64  `json:"seederCount,omitempty"`
	Tier                  int64  `json:"tier,omitempty"`
}

type Torrent struct {
	ActivityDate            int64         `json:"activityDate,omitempty"`
	AddedDate               int64         `json:"addedDate,omitempty"`
	BandwidthPriority       int64         `json:"bandwidthPriority,omitempty"`
	Comment                 string        `json:"comment,omitempty"`
	CorruptEver             int64         `json:"corruptEver,omitempty"`
	Creator                 string        `json:"creator,omitempty"`
	DateCreated             int64         `json:"dateCreated,omitempty"`
	DesiredAvailable        int64         `json:"desiredAvailable,omitempty"`
	DoneDate                int64         `json:"doneDate,omitempty"`
	DownloadDir             string        `json:"downloadDir,omitempty"`
	DownloadedEver          int64         `json:"downloadedEver,omitempty"`
	DownloadLimit           int64         `json:"downloadLimit,omitempty"`
	DownloadLimited         bool          `json:"downloadLimited,omitempty"`
	EditDate                int64         `json:"editDate,omitempty"`
	Error                   int64         `json:"error,omitempty"`
	ErrorString             string        `json:"errorString,omitempty"`
	Eta                     int64         `json:"eta,omitempty"`
	EtaIdle                 int64         `json:"etaIdle,omitempty"`
	Files                   []File        `json:"files,omitempty"`
	FileStats               []FileStat    `json:"fileStats,omitempty"`
	HashString              string        `json:"hashString,omitempty"`
	HaveUnchecked           int64         `json:"haveUnchecked,omitempty"`
	HaveValid               int64         `json:"haveValid,omitempty"`
	HonorsSessionLimits     bool          `json:"honorsSessionLimits,omitempty"`
	Id                      int64         `json:"id,omitempty"`
	IsFinished              bool          `json:"isFinished,omitempty"`
	IsPrivate               bool          `json:"isPrivate,omitempty"`
	IsStalled               bool          `json:"isStalled,omitempty"`
	Labels                  []string      `json:"labels,omitempty"`
	LeftUntilDone           int64         `json:"leftUntilDone,omitempty"`
	MagnetLink              string        `json:"magnetLink,omitempty"`
	ManualAnnounceTime      int64         `json:"manualAnnounceTime,omitempty"`
	MaxConnectedPeers       int64         `json:"maxConnectedPeers,omitempty"`
	MetadataPercentComplete float64       `json:"metadataPercentComplete,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Path                    string        `json:"path,omitempty"` // Used
	PeerLimit               int64         `json:"peer-limit,omitempty"`
	Peers                   []Peer        `json:"peers,omitempty"`
	PeersConnected          int64         `json:"peersConnected,omitempty"`
	PeersFrom               []PeerFrom    `json:"peersFrom,omitempty"`
	PeersGettingFromUs      int64         `json:"peersGettingFromUs,omitempty"`
	PeersSendingToUs        int64         `json:"peersSendingToUs,omitempty"`
	PercentDone             float64       `json:"percentDone,omitempty"`
	Pieces                  string        `json:"pieces,omitempty"`
	PieceCount              int64         `json:"pieceCount,omitempty"`
	PieceSize               int64         `json:"pieceSize,omitempty"`
	Priorities              []int64       `json:"priorities,omitempty"`
	QueuePosition           int64         `json:"queuePosition,omitempty"`
	RateDownload            int64         `json:"rateDownload,omitempty"`
	RateUpload              int64         `json:"rateUpload,omitempty"`
	RecheckProgress         float64       `json:"recheckProgress,omitempty"`
	SecondsDownloading      int64         `json:"secondsDownloading,omitempty"`
	SecondsSeeding          int64         `json:"secondsSeeding,omitempty"`
	SeedIdleLimit           int64         `json:"seedIdleLimit,omitempty"`
	SeedIdleMode            int64         `json:"seedIdleMode,omitempty"`
	SeedRatioLimit          float64       `json:"seedRatioLimit,omitempty"`
	SeedRatioMode           int64         `json:"seedRatioMode,omitempty"`
	SizeWhenDone            int64         `json:"sizeWhenDone,omitempty"`
	StartDate               int64         `json:"startDate,omitempty"`
	Status                  int64         `json:"status,omitempty"`
	Trackers                []Tracker     `json:"trackers,omitempty"`
	TrackerStats            []TrackerStat `json:"trackerStats,omitempty"`
	TotalSize               int64         `json:"totalSize,omitempty"`
	TorrentFile             string        `json:"torrentFile,omitempty"`
	UploadedEver            int64         `json:"uploadedEver,omitempty"`
	UploadLimit             int64         `json:"uploadLimit,omitempty"`
	UploadLimited           bool          `json:"uploadLimited,omitempty"`
	UploadRatio             float64       `json:"uploadRatio,omitempty"`
	Wanted                  []int64       `json:"wanted,omitempty"`
	WebSeeds                []string      `json:"webseeds,omitempty"`
	WebSeedsSendingToUs     int64         `json:"webseedsSendingToUs,omitempty"`
}

type RequestArgs struct {
	Ids    []int64  `json:"ids,omitempty"`
	Fields []string `json:"fields,omitempty"`
	Format string   `json:"format,omitempty"`
}

type AddPayload struct {
	Cookies           string  `json:"cookies,omitempty"`
	DownloadDir       string  `json:"download-dir,omitempty"`
	Filename          string  `json:"filename,omitempty"`
	MetaInfo          string  `json:"metainfo,omitempty"`
	Paused            bool    `json:"paused,omitempty"`
	PeerLimit         int64   `json:"peer-limit,omitempty"`
	BandwidthPriority int64   `json:"bandwidthPriority,omitempty"`
	FilesWanted       []int64 `json:"files-wanted,omitempty"`
	FilesUnwanted     []int64 `json:"files-unwanted,omitempty"`
	PriorityHigh      []int64 `json:"priority-high,omitempty"`
	PriorityLow       []int64 `json:"priority-low,omitempty"`
	PriorityNormal    []int64 `json:"priority-normal,omitempty"`
}

type MovePayload struct {
	Ids      []string `json:"ids,omitempty"`
	Location string   `json:"location,omitempty"`
	Move     bool     `json:"move,omitempty"`
}

type RemovePayload struct {
	Ids             []string `json:"ids,omitempty"`
	DeleteLocalData bool     `json:"delete-local-data,omitempty"`
}

type RenamePayload struct {
	Ids  string `json:"ids,omitempty"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type ResponseArgs struct {
	Torrents     []Torrent `json:"torrents,omitempty"`
	TorrentAdded Torrent   `json:"torrent-added,omitempty"`
	Id           int64     `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Path         string    `json:"path,omitempty"`
}

type Response struct {
	Result string       `json:"result,omitempty"`
	Args   ResponseArgs `json:"arguments,omitempty"`
	Tag    int64        `json:"tag,omitempty"`
}

type RequestAuth struct {
	Username string
	Password string
}

type Request struct {
	Method    Method      `json:"method,omitempty"`
	Arguments interface{} `json:"arguments,omitempty"`
	Tag       int64       `json:"tag,omitempty"`
}

type Option func(*options)

type Client struct {
	options
}

type options struct {
	Username  string
	Password  string
	URL       string
	sessionId string
}

func WithURL(url string) Option {
	return func(o *options) {
		o.URL = url
	}
}

func WithBasicAuth(user, password string) Option {
	return func(o *options) {
		o.Username = user
		o.Password = password
	}
}

func New(opts ...Option) *Client {
	defaults := options{}
	for _, o := range opts {
		o(&defaults)
	}
	return &Client{options: defaults}
}

func (c *Client) fetch(request Request) (Response, error) {
	var empty Response
	body, err := json.Marshal(&request)
	if err != nil {
		return empty, err
	}

	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return empty, err
	}

	if c.Password != "" && c.Username != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}

	req.Header.Set("User-Agent", "transmission")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(SessionIdHeader, c.sessionId)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return empty, err
	}

	if resp.StatusCode == http.StatusConflict {
		c.sessionId = resp.Header.Get(SessionIdHeader)
		return c.fetch(request)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	defer resp.Body.Close()

	var response Response
	if err := json.Unmarshal(buf, &response); err != nil {
		return empty, err
	}
	return response, nil
}

func (c *Client) getAll(ids []int64, fields []string) ([]Torrent, error) {
	var empty []Torrent
	if len(fields) == 0 {
		return empty, errors.New("request must includes at least one field")
	}

	response, err := c.fetch(Request{
		Method:    MethodGet,
		Arguments: RequestArgs{Ids: ids, Fields: fields},
	})

	if err != nil {
		return empty, err
	}

	if response.Result == ResponseResultError {
		return empty, errors.New("unable to get results from transmission service")
	}

	return response.Args.Torrents, nil
}

func (c *Client) Ping() error {
	_, err := c.fetch(Request{})
	return err
}

func (c *Client) Get(id int64, fields []string) (Torrent, error) {
	var empty Torrent
	if len(fields) == 0 {
		return empty, errors.New("request must includes at least one field")
	}

	torrents, err := c.getAll([]int64{id}, fields)
	if err != nil {
		return empty, err
	}

	if len(torrents) == 1 {
		return torrents[0], nil
	}

	return empty, errors.New("torrent not found")
}

func (c *Client) GetAll(ids []int64, fields []string) ([]Torrent, error) {
	return c.getAll(ids, fields)
}

// Torrent actions
func (c *Client) performAction(ids []int64, method Method) error {
	resp, err := c.fetch(Request{
		Method:    method,
		Arguments: RequestArgs{Ids: ids},
	})

	if err != nil {
		return err
	}

	if resp.Result == ResponseResultError {
		return errors.New("could not perform action")
	}
	return nil
}

func (c *Client) Start(id int64) error {
	return c.performAction([]int64{id}, MethodStart)
}

func (c *Client) StartAll(ids []int64) error {
	return c.performAction(ids, MethodStart)
}

func (c *Client) StartNow(id int64) error {
	return c.performAction([]int64{id}, MethodStartNow)
}

func (c *Client) StartAllNow(ids []int64) error {
	return c.performAction(ids, MethodStartNow)
}

func (c *Client) Stop(id int64) error {
	return c.performAction([]int64{id}, MethodStop)
}

func (c *Client) StopAll(ids []int64) error {
	return c.performAction(ids, MethodStop)
}

func (c *Client) Verify(id int64) error {
	return c.performAction([]int64{id}, MethodVerify)
}

func (c *Client) VerifyAll(ids []int64) error {
	return c.performAction(ids, MethodVerify)
}

func (c *Client) Reannounce(id int64) error {
	return c.performAction([]int64{id}, MethodReannounce)
}

func (c *Client) ReannounceAll(ids []int64) error {
	return c.performAction(ids, MethodReannounce)
}

// Torrent mutators

// Add torrent
func (c *Client) Add(args AddPayload) (Torrent, error) {
	var empty Torrent
	resp, err := c.fetch(Request{
		Method:    MethodAdd,
		Arguments: args,
	})
	if err != nil {
		return empty, err
	}
	return resp.Args.TorrentAdded, nil
}

// Remove torrent
func (c *Client) Remove(args RemovePayload) error {
	_, err := c.fetch(Request{
		Method:    MethodRemove,
		Arguments: args,
	})
	return err
}

func (c *Client) Rename(args RenamePayload) (Torrent, error) {
	var torrent Torrent
	if len(args.Ids) > 1 {
		return torrent, errors.New("could not edit multiple torrents at the same time")
	}

	resp, err := c.fetch(Request{
		Method:    MethodRename,
		Arguments: args,
	})

	if err != nil {
		return torrent, err
	}

	torrent.Id = resp.Args.Id
	torrent.Name = resp.Args.Name
	torrent.Path = resp.Args.Path

	return torrent, nil
}
