package api

type ProgramDefinitions struct {
	Status      int                 `json:"Status"`
	Message     string              `json:"Message"`
	Definitions []ProgramDefinition `json:"Definitions"`
}

type ProgramDefinition struct {
	ProgramID    int    `json:"ProgramId"`
	DetectionXML string `json:"DetectionXML"`
}

type ScanRequest struct {
	Programs      []Program             `json:"Programs"`
	ResVersion    string                `json:"ResVersion"`
	Parameters    ScanRequestParameters `json:"Parameters"`
	Time          float32               `json:"Time"`
	ClientVersion string                `json:"ClientVersion"`
	System        System                `json:"System"`
	ScanType      string                `json:"ScanType"`
}

type ScanRequestParameters struct {
	HideBeta     bool `json:"HideBeta"`
	ShowDetailed bool `json:"ShowDetailed"`
	ShowAll      bool `json:"ShowAll"`
}

type System struct {
	VersionString string   `json:"VersionString"`
	VersionMajor  int      `json:"VersionMajor"`
	VersionMinor  int      `json:"VersionMinor"`
	Revision      int      `json:"Revision"`
	ServicePack   string   `json:"ServicePack"`
	Bit           int      `json:"Bit"`
	IsServer      bool     `json:"IsServer"`
	GraphicsCards []string `json:"GraphicsCards"`
}

type Program struct {
	Registry interface{} `json:"Registry"`
	ID       int         `json:"Id"`
	Files    []Files     `json:"Files"`
}

type Files struct {
	File   string `json:"File"`
	VerPv  string `json:"VerPv"`
	VerPvr string `json:"VerPvr"`
	VerFv  string `json:"VerFv"`
	VerFvr string `json:"VerFvr"`
	Len    int    `json:"Len"`
	MD5    string `json:"Md5"`
}

type ScanResults struct {
	ID        string      `json:"Id"`
	Status    int         `json:"Status"`
	Message   interface{} `json:"Message"`
	Timestamp string      `json:"Timestamp"`
	Updates   []Update    `json:"Updates"`
}

type Update struct {
	HasUpdate                  bool   `json:"HasUpdate"`
	Size                       int    `json:"Size"`
	InstalledDatePublished     string `json:"InstalledDatePublished"`
	Title                      string `json:"Title"`
	ProgramID                  int    `json:"ProgramID"`
	ApplicationManagerIconData string `json:"ApplicationManagerIconData"`
	DatePublished              string `json:"DatePublished"`
	IconURL                    string `json:"IconURL"`
	InstalledVersion           string `json:"InstalledVersion"`
	LatestVersion              bool   `json:"LatestVersion"`
	InstallationPath           bool   `json:"InstallationPath"`
	FileID                     int    `json:"FileID"`
	IsBeta                     bool   `json:"IsBeta"`
	DownloadURL                string `json:"DownloadUrl"`
}

type InstallerProgramsRequest struct {
	Programs []InstallerProgram `json:"Programs"`
}

// TODO: This probably isn't complete..
type InstallerProgramsResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

type InstallerProgram struct {
	Publisher        string `json:"Publisher"`
	Name             string `json:"Name"`
	Is64Bit          bool   `json:"Is64Bit"`
	Version          string `json:"Version"`
	HiveLocalMachine bool   `json:"HiveLocalMachine"`
	HiveCurrentUser  bool   `json:"HiveCurrentUser"`
	Key              string `json:"Key"`
}
