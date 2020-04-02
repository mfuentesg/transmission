package transmission

type Units struct {
	SpeedUnits  []string `json:"speed-units,omitempty"`
	SpeedBytes  int64    `json:"speed-bytes,omitempty"`
	SizeUnits   []string `json:"size-units,omitempty"`
	SizeBytes   int64    `json:"size-bytes,omitempty"`
	MemoryUnits []string `json:"memory-units,omitempty"`
	MemoryBytes int64    `json:"memory-bytes,omitempty"`
}

type CumulativeStats struct {
	UploadedBytes   int64 `json:"uploadedBytes"`
	DownloadedBytes int64 `json:"downloadedBytes"`
	FilesAdded      int64 `json:"filesAdded"`
	SessionCount    int64 `json:"sessionCount"`
	SecondsActive   int64 `json:"secondsActive"`
}

type CurrentStats struct {
	UploadedBytes   int64 `json:"uploadedBytes"`
	DownloadedBytes int64 `json:"downloadedBytes"`
	FilesAdded      int64 `json:"filesAdded"`
	SessionCount    int64 `json:"sessionCount"`
	SecondsActive   int64 `json:"secondsActive"`
}

type SessionStats struct {
	ActiveTorrentCount int64           `json:"activeTorrentCount"`
	DownloadSpeed      int64           `json:"downloadSpeed"`
	PausedTorrentCount int64           `json:"pausedTorrentCount"`
	TorrentCount       int64           `json:"torrentCount"`
	UploadSpeed        int64           `json:"uploadSpeed"`
	CumulativeStats    CumulativeStats `json:"cumulative-stats"`
	CurrentStats       CurrentStats    `json:"current-stats"`
}

type Session struct {
	AltSpeedDown              int64   `json:"alt-speed-down,omitempty"`
	AltSpeedTimeBegin         int64   `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnd           int64   `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int64   `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int64   `json:"alt-speed-up,omitempty"`
	BlockListURL              string  `json:"blocklist-url,omitempty"`
	BlockListSize             int64   `json:"blocklist-size,omitempty"`
	CacheSizeMb               int64   `json:"cache-size-mb,omitempty"`
	ConfigDir                 string  `json:"config-dir,omitempty"`
	DownloadDir               string  `json:"download-dir,omitempty"`
	Encryption                string  `json:"encryption,omitempty"`
	DownloadQueueSize         int64   `json:"download-queue-size,omitempty"`
	IdleSeedingLimit          int64   `json:"idle-seeding-limit,omitempty"`
	IncompleteDir             string  `json:"incomplete-dir,omitempty"`
	PeerLimitGlobal           int64   `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int64   `json:"peer-limit-per-torrent,omitempty"`
	PeerPort                  int64   `json:"peer-port,omitempty"`
	QueueStalledMinutes       int64   `json:"queue-stalled-minutes,omitempty"`
	RPCVersion                int64   `json:"rpc-version,omitempty"`
	RPCVersionMinimum         int64   `json:"rpc-version-minimum,omitempty"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename,omitempty"`
	SeedRatioLimit            float64 `json:"seedRatioLimit,omitempty"`
	SeedQueueSize             int64   `json:"seed-queue-size,omitempty"`
	SpeedLimitDown            int64   `json:"speed-limit-down,omitempty"`
	SpeedLimitUp              int64   `json:"speed-limit-up,omitempty"`
	Units                     Units   `json:"units,omitempty"`
	Version                   string  `json:"version,omitempty"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled"`
	BlockListEnabled          bool    `json:"blocklist-enabled"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled"`
	DhtEnabled                bool    `json:"dht-enabled"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled"`
	LpdEnabled                bool    `json:"lpd-enabled"`
	PexEnabled                bool    `json:"pex-enabled"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled"`
	RenamePartialFiles        bool    `json:"rename-partial-files"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled"`
	SeedRatioLimited          bool    `json:"seedRatioLimited"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled"`
	StartAddedTorrents        bool    `json:"start-added-torrents"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files"`
	UtpEnabled                bool    `json:"utp-enabled"`
}
