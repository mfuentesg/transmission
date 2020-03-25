package transmission

type Units struct {
	SpeedUnits  [4]string // TODO: set default value "kB/s",	"MB/s",	"GB/s",	"TB/s"
	SpeedBytes  int64
	SizeUnits   [4]string // TODO: set default value kB, MB, GB, TB
	SizeBytes   int64
	MemoryUnits [4]string // TODO: set default value KiB, MiB, GiB, TiB
	MemoryBytes int64
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
	AltSpeedEnabled           bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin         int64   `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd           int64   `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int64   `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int64   `json:"alt-speed-up,omitempty"`
	BlockListUrl              string  `json:"blocklist-url,omitempty"`
	BlockListEnabled          bool    `json:"blocklist-enabled,omitempty"`
	BlockListSize             int64   `json:"blocklist-size,omitempty"`
	CacheSizeMb               int64   `json:"cache-size-mb,omitempty"`
	ConfigDir                 string  `json:"config-dir,omitempty"`
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
	RpcVersion                int64   `json:"rpc-version,omitempty"`
	RpcVersionMinimum         int64   `json:"rpc-version-minimum,omitempty"`
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
	Units                     Units   `json:"units,omitempty"`
	UtpEnabled                bool    `json:"utp-enabled,omitempty"`
	Version                   string  `json:"version,omitempty"`
}
