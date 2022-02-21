package ydl

type Format struct {
	Ext         string      `json:"ext"`
	FormatNote  string      `json:"format_note"`
	Acodec      string      `json:"acodec"`
	Abr         float64     `json:"abr,omitempty"`
	Container   string      `json:"container,omitempty"`
	FormatID    string      `json:"format_id"`
	URL         string      `json:"url"`
	ManifestURL string      `json:"manifest_url,omitempty"`
	Width       int         `json:"width,omitempty"`
	Height      int         `json:"height,omitempty"`
	Tbr         float64     `json:"tbr,omitempty"`
	Asr         int         `json:"asr,omitempty"`
	Fps         float64     `json:"fps,omitempty"`
	Language    interface{} `json:"language,omitempty"`
	Filesize    int         `json:"filesize,omitempty"`
	Vcodec      string      `json:"vcodec"`
	Format      string      `json:"format"`
	Protocol    string      `json:"protocol"`
	HTTPHeaders struct {
		UserAgent      string `json:"User-Agent"`
		AcceptCharset  string `json:"Accept-Charset"`
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
	} `json:"http_headers"`
	PlayerURL  string `json:"player_url,omitempty"`
	Resolution string `json:"resolution,omitempty"`
}
type Formats []Format

type RequestedFormat struct {
	Ext         string      `json:"ext"`
	Height      int         `json:"height,omitempty"`
	FormatNote  string      `json:"format_note"`
	Vcodec      string      `json:"vcodec"`
	FormatID    string      `json:"format_id"`
	URL         string      `json:"url"`
	ManifestURL string      `json:"manifest_url,omitempty"`
	Width       int         `json:"width,omitempty"`
	Tbr         float64     `json:"tbr"`
	Asr         int         `json:"asr,omitempty"`
	Fps         int         `json:"fps,omitempty"`
	Language    interface{} `json:"language,omitempty"`
	Filesize    int         `json:"filesize"`
	Acodec      string      `json:"acodec"`
	Format      string      `json:"format"`
	Protocol    string      `json:"protocol"`
	HTTPHeaders struct {
		UserAgent      string `json:"User-Agent"`
		AcceptCharset  string `json:"Accept-Charset"`
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
	} `json:"http_headers"`
	PlayerURL string  `json:"player_url,omitempty"`
	Abr       float64 `json:"abr,omitempty"`
}
type RequestedFormats []RequestedFormat

type Thumbnail struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}
type Thumbnails []Thumbnail

// Info is an object representing the JSON returned from a -J download (dump-single-json)
type Info struct {
	ID          string   `json:"id"`
	Uploader    string   `json:"uploader"`
	UploaderID  string   `json:"uploader_id"`
	UploaderURL string   `json:"uploader_url"`
	UploadDate  string   `json:"upload_date"`
	License     string   `json:"license"`
	Creator     string   `json:"creator"`
	Title       string   `json:"title"`
	AltTitle    string   `json:"alt_title"`
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`
	Subtitles   struct {
	} `json:"subtitles"`
	AutomaticCaptions struct {
	} `json:"automatic_captions"`
	Duration           float64          `json:"duration"`
	AgeLimit           int              `json:"age_limit"`
	Annotations        interface{}      `json:"annotations"`
	Chapters           interface{}      `json:"chapters"`
	WebpageURL         string           `json:"webpage_url"`
	ViewCount          int              `json:"view_count"`
	LikeCount          int              `json:"like_count"`
	DislikeCount       int              `json:"dislike_count"`
	AverageRating      float64          `json:"average_rating"`
	Formats            Formats          `json:"formats"`
	IsLive             interface{}      `json:"is_live"`
	StartTime          interface{}      `json:"start_time"`
	EndTime            interface{}      `json:"end_time"`
	Series             interface{}      `json:"series"`
	SeasonNumber       interface{}      `json:"season_number"`
	EpisodeNumber      interface{}      `json:"episode_number"`
	Extractor          string           `json:"extractor"`
	WebpageURLBasename string           `json:"webpage_url_basename"`
	ExtractorKey       string           `json:"extractor_key"`
	Playlist           interface{}      `json:"playlist"`
	PlaylistIndex      interface{}      `json:"playlist_index"`
	Thumbnails         Thumbnails       `json:"thumbnails"`
	DisplayID          string           `json:"display_id"`
	RequestedSubtitles interface{}      `json:"requested_subtitles"`
	RequestedFormats   RequestedFormats `json:"requested_formats"`
	Format             string           `json:"format"`
	FormatID           string           `json:"format_id"`
	Width              int              `json:"width"`
	Height             int              `json:"height"`
	Resolution         string           `json:"resolution"`
	Fps                int              `json:"fps"`
	Vcodec             string           `json:"vcodec"`
	Vbr                float64          `json:"vbr"`
	StretchedRatio     interface{}      `json:"stretched_ratio"`
	Acodec             string           `json:"acodec"`
	Abr                float64          `json:"abr"`
	Ext                string           `json:"ext"`
	Fulltitle          string           `json:"fulltitle"`
	Filename           string           `json:"_filename"`
}
