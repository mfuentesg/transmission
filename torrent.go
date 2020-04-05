package transmission

type File struct {
	BytesCompleted int64  `json:"bytesCompleted,omitempty"`
	Length         int64  `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
}

type FileStat struct {
	BytesCompleted int64 `json:"bytesCompleted,omitempty"`
	Priority       int64 `json:"priority,omitempty"`
	Wanted         bool  `json:"wanted"`
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	FlagStr            string  `json:"flagStr,omitempty"`
	Port               int64   `json:"port,omitempty"`
	Progress           float64 `json:"progress,omitempty"`
	RateToClient       int64   `json:"rateToClient,omitempty"`
	RateToPeer         int64   `json:"rateToPeer,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom"`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsIncoming         bool    `json:"isIncoming"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	IsUTP              bool    `json:"isUTP"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
}

type PeersFrom struct {
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
	ID       int64  `json:"id,omitempty"`
	Scrape   string `json:"scrape,omitempty"`
	Tier     int64  `json:"tier,omitempty"`
}

type TrackerStat struct {
	Announce              string  `json:"announce,omitempty"`
	AnnounceState         int64   `json:"announceState,omitempty"`
	DownloadCount         int64   `json:"downloadCount,omitempty"`
	Host                  string  `json:"host,omitempty"`
	ID                    int64   `json:"id,omitempty"`
	LastAnnouncePeerCount int64   `json:"lastAnnouncePeerCount,omitempty"`
	LastAnnounceResult    string  `json:"lastAnnounceResult,omitempty"`
	LastAnnounceStartTime int64   `json:"lastAnnounceStartTime,omitempty"`
	LastAnnounceTime      int64   `json:"lastAnnounceTime,omitempty"`
	LastScrapeResult      string  `json:"lastScrapeResult,omitempty"`
	LastScrapeStartTime   int64   `json:"lastScrapeStartTime,omitempty"`
	LastScrapeTime        int64   `json:"lastScrapeTime,omitempty"`
	LeecherCount          int64   `json:"leecherCount,omitempty"`
	NextAnnounceTime      int64   `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime        int64   `json:"nextScrapeTime,omitempty"`
	Scrape                string  `json:"scrape,omitempty"`
	ScrapeState           int64   `json:"scrapeState,omitempty"`
	SeederCount           int64   `json:"seederCount,omitempty"`
	Tier                  int64   `json:"tier,omitempty"`
	HasAnnounced          bool    `json:"hasAnnounced"`
	HasScraped            bool    `json:"hasScraped"`
	IsBackup              bool    `json:"isBackup"`
	LastAnnounceSucceeded bool    `json:"lastAnnounceSucceeded"`
	LastAnnounceTimedOut  bool    `json:"lastAnnounceTimedOut"`
	LastScrapeSucceeded   bool    `json:"lastScrapeSucceeded"`
	LastScrapeTimedOut    NumBool `json:"lastScrapeTimedOut"`
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
	ID                      int64         `json:"id,omitempty"`
	Labels                  []string      `json:"labels,omitempty"`
	LeftUntilDone           int64         `json:"leftUntilDone,omitempty"`
	MagnetLink              string        `json:"magnetLink,omitempty"`
	ManualAnnounceTime      int64         `json:"manualAnnounceTime,omitempty"`
	MaxConnectedPeers       int64         `json:"maxConnectedPeers,omitempty"`
	MetadataPercentComplete float64       `json:"metadataPercentComplete,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Path                    string        `json:"path,omitempty"` // This field is not part of the standard response
	PeerLimit               int64         `json:"peer-limit,omitempty"`
	Peers                   []Peer        `json:"peers,omitempty"`
	PeersConnected          int64         `json:"peersConnected,omitempty"`
	PeersFrom               PeersFrom     `json:"peersFrom,omitempty"`
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
	UploadRatio             float64       `json:"uploadRatio,omitempty"`
	Wanted                  []int64       `json:"wanted,omitempty"`
	WebSeeds                []string      `json:"webseeds,omitempty"`
	WebSeedsSendingToUs     int64         `json:"webseedsSendingToUs,omitempty"`
	DownloadLimited         bool          `json:"downloadLimited"`
	HonorsSessionLimits     bool          `json:"honorsSessionLimits"`
	IsFinished              bool          `json:"isFinished"`
	IsPrivate               bool          `json:"isPrivate"`
	IsStalled               bool          `json:"isStalled"`
	UploadLimited           bool          `json:"uploadLimited"`
}
