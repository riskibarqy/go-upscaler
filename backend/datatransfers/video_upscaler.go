package datatransfers

type VideoUpscalerRequest struct {
	InputPlainFileName string // filename without extension
	InputFullFileName  string
	InputFileExt       string // .mp4, .mkv etc
	InputFileSize      int64
	TempFilePath       string
	TempDir            string
	Model              string
	VideoFPS           int // if its not filled, it will automatically use default video fps
	AudioFileName      string
	ScaleMultiplier    int // realersgan params : scale multiplier 2, 3, 4 default : 4
	TileSize           int // Real-ESRGAN parameter: Default = 0 (auto). Higher values improve detail but increase GPU memory usage.
	SavePath           string
	IsHaveAudio        bool
	LoadingProgress    int
	TotalBatches       int
	CurrentBatch       int
}

type InputFileRequest struct {
	FileCode   string
	FileBase64 string
	FileName   string
	Model      string
	Scale      int
}

type FFProbeStreamsMetadataResponse struct {
	TotalFrames int
	FPS         int
	Height      int
	Width       int
}
