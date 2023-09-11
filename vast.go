// Package vast parses, manipulates and builds Digital Video Ad Serving Templates (VAST).
// The currently supported version is VAST 4.2.
package vast

import (
	"encoding/xml"
	"errors"
	"io"
)

type NumericBool bool

func (n NumericBool) MarshalText() ([]byte, error) {
	if n {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

type CData struct {
	Value string `xml:",cdata"`
}

type Ad struct {
	InLine        *InLine     `xml:"InLine,omitempty"`
	Wrapper       *Wrapper    `xml:"Wrapper,omitempty"`
	ID            string      `xml:"id,attr,omitempty"`
	Sequence      int         `xml:"sequence,attr,omitempty"`
	ConditionalAd NumericBool `xml:"conditionalAd,attr,omitempty"`
	AdType        AdType      `xml:"adType,attr,omitempty"`
}

type Impression struct {
	Value string `xml:",cdata"`
	ID    string `xml:"id,attr,omitempty"`
}

type AdDefinitionBase struct {
	AdSystem           AdSystem            `xml:"AdSystem"`
	Error              []CData             `xml:"Error,omitempty"`
	Extensions         *Extensions         `xml:"Extensions,omitempty"`
	Impression         []Impression        `xml:"Impression"`
	Pricing            *Pricing            `xml:"Pricing,omitempty"`
	ViewableImpression *ViewableImpression `xml:"ViewableImpression,omitempty"`
}

type AdParameters struct {
	Value      string      `xml:",chardata"`
	XMLEncoded NumericBool `xml:"xmlEncoded,attr,omitempty"`
}

type AdSystem struct {
	Value   string `xml:",chardata"`
	Version string `xml:"version,attr,omitempty"`
}

type AdType string

const (
	VideoAdType  AdType = "video"
	AudioAdType  AdType = "audio"
	HybridAdType AdType = "hybrid"
)

type AdVerifications struct {
	Verification []Verification `xml:"Verification,omitempty"`
}

type InLineCreatives struct {
	Creative []InLineCreative `xml:"Creative"`
}

type BlockedAdCategories struct {
	Value     string `xml:",chardata"`
	Authority string `xml:"authority,attr,omitempty"`
}

type Category struct {
	Value     string `xml:",chardata"`
	Authority string `xml:"authority,attr"`
}

type ClosedCaptionFile struct {
	Value    string `xml:",cdata"`
	Type     string `xml:"type,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
}

type ClosedCaptionFiles struct {
	ClosedCaptionFile []ClosedCaptionFile `xml:"ClosedCaptionFile,omitempty"`
}

type CompanionAdsCollection struct {
	Companion []CompanionAd `xml:"Companion,omitempty"`
	Required  Required      `xml:"required,attr,omitempty"`
}

type CompanionAd struct {
	HTMLResource           []CData             `xml:"HTMLResource,omitempty"`
	IFrameResource         []CData             `xml:"IFrameResource,omitempty"`
	StaticResource         []StaticResource    `xml:"StaticResource,omitempty"`
	AdParameters           *AdParameters       `xml:"AdParameters,omitempty"`
	AltText                string              `xml:"AltText,omitempty"`
	CompanionClickThrough  *CData              `xml:"CompanionClickThrough,omitempty"`
	CompanionClickTracking []string            `xml:"CompanionClickTracking,omitempty"`
	CreativeExtensions     *CreativeExtensions `xml:"CreativeExtensions,omitempty"`
	TrackingEvents         *TrackingEvents     `xml:"TrackingEvents,omitempty"`
	ID                     string              `xml:"id,attr,omitempty"`
	Width                  int                 `xml:"width,attr"`
	Height                 int                 `xml:"height,attr"`
	AssetWidth             int                 `xml:"assetWidth,attr,omitempty"`
	AssetHeight            int                 `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth          int                 `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int                 `xml:"expandedHeight,attr,omitempty"`
	APIFramework           string              `xml:"apiFramework,attr,omitempty"`
	AdSlotID               string              `xml:"adSlotId,attr,omitempty"`
	PXRatio                float64             `xml:"pxratio,attr,omitempty"`
	RenderingMode          RenderingMode       `xml:"renderingMode,attr,omitempty"`
}

type CreativeBase struct {
	Sequence     int    `xml:"sequence,attr,omitempty"`
	APIFramework string `xml:"apiFramework,attr,omitempty"`
	AdID         string `xml:"adId,attr,omitempty"`
}

type CreativeExtension struct {
	Items []string `xml:",any"`
	Type  string   `xml:"type,attr,omitempty"`
}

type CreativeExtensions struct {
	CreativeExtension []CreativeExtension `xml:"CreativeExtension,omitempty"`
}

type InLineCreative struct {
	CreativeBase

	CompanionAds       *CompanionAdsCollection `xml:"CompanionAds,omitempty"`
	CreativeExtensions *CreativeExtensions     `xml:"CreativeExtensions,omitempty"`
	Linear             *LinearInLine           `xml:"Linear,omitempty"`
	NonLinearAds       *NonLinearAds           `xml:"NonLinearAds,omitempty"`
	UniversalAdID      []UniversalAdID         `xml:"UniversalAdId"`
	ID                 string                  `xml:"id,attr,omitempty"`
}

type CreativeResource struct {
	HTMLResource   []CData          `xml:"HTMLResource,omitempty"`
	IFrameResource []CData          `xml:"IFrameResource,omitempty"`
	StaticResource []StaticResource `xml:"StaticResource,omitempty"`
}

type WrapperCreative struct {
	CreativeBase

	CompanionAds *CompanionAdsCollection `xml:"CompanionAds,omitempty"`
	Linear       *LinearWrapper          `xml:"Linear,omitempty"`
	NonLinearAds *NonLinearAds           `xml:"NonLinearAds,omitempty"`
	ID           string                  `xml:"id,attr,omitempty"`
}

type Creatives struct {
	Creative []WrapperCreative `xml:"Creative"`
}

// Currency must match the pattern `[a-zA-Z]{3}`.
type Currency string

type Delivery string

const (
	StreamingDelivery   Delivery = "streaming"
	ProgressiveDelivery Delivery = "progressive"
)

type Event string

const (
	MuteEvent                Event = "mute"
	UnmuteEvent              Event = "unmute"
	PauseEvent               Event = "pause"
	ResumeEvent              Event = "resume"
	RewindEvent              Event = "rewind"
	SkipEvent                Event = "skip"
	PlayerExpandEvent        Event = "playerExpand"
	PlayerCollapseEvent      Event = "playerCollapse"
	LoadedEvent              Event = "loaded"
	StartEvent               Event = "start"
	FirstQuartileEvent       Event = "firstQuartile"
	MidpointEvent            Event = "midpoint"
	ThirdQuartileEvent       Event = "thirdQuartile"
	CompleteEvent            Event = "complete"
	ProgressEvent            Event = "progress"
	CloseLinearEvent         Event = "closeLinear"
	CreativeViewEvent        Event = "creativeView"
	AcceptInvitationEvent    Event = "acceptInvitation"
	AdExpandEvent            Event = "adExpand"
	AdCollapseEvent          Event = "adCollapse"
	MinimizeEvent            Event = "minimize"
	CloseEvent               Event = "close"
	OverlayViewDurationEvent Event = "overlayViewDuration"
	OtherAdInteraction       Event = "otherAdInteraction"
	InteractiveStart         Event = "interactiveStart"
)

type ExecutableResource struct {
	Value        string `xml:",cdata"`
	APIFramework string `xml:"apiFramework,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`
}

type Extension struct {
	Value string `xml:",innerxml"`
	Type  string `xml:"type,attr,omitempty"`
}

type Extensions struct {
	Extension []Extension `xml:"Extension,omitempty"`
}

type IconClickFallbackImage struct {
	AltText        string `xml:"AltText,omitempty"`
	StaticResource *CData `xml:"StaticResource,omitempty"`
	Height         int    `xml:"height,attr,omitempty"`
	Width          int    `xml:"width,attr,omitempty"`
}

type IconClickFallbackImages struct {
	IconClickFallbackImage []IconClickFallbackImage `xml:"IconClickFallbackImage"`
}

type IconClicks struct {
	IconClickFallbackImages *IconClickFallbackImages `xml:"IconClickFallbackImages,omitempty"`
	IconClickThrough        string                   `xml:"IconClickThrough,omitempty"`
	IconClickTracking       []string                 `xml:"IconClickTracking,omitempty"`
}

type Icons struct {
	Icon []Icon `xml:"Icon"`
}

// Duration must be expressed in the standard time format `hh:mm:ss`.
type Duration string

type Icon struct {
	CreativeResource

	HTMLResource     []CData          `xml:"HTMLResource,omitempty"`
	IFrameResource   []CData          `xml:"IFrameResource,omitempty"`
	StaticResource   []StaticResource `xml:"StaticResource,omitempty"`
	IconClicks       *IconClicks      `xml:"IconClicks,omitempty"`
	IconViewTracking []string         `xml:"IconViewTracking,omitempty"`
	Program          string           `xml:"program,attr,omitempty"`
	Width            int              `xml:"width,attr,omitempty"`
	Height           int              `xml:"height,attr,omitempty"`
	XPosition        XPosition        `xml:"xPosition,attr,omitempty"`
	YPosition        YPosition        `xml:"yPosition,attr,omitempty"`
	Duration         Duration         `xml:"duration,attr,omitempty"`
	Offset           Duration         `xml:"offset,attr,omitempty"`
	APIFramework     string           `xml:"apiFramework,attr,omitempty"`
	PXRatio          float64          `xml:"pxratio,attr,omitempty"`
}

type InLine struct {
	AdDefinitionBase

	AdServingID     string           `xml:"AdServingId"`
	AdTitle         string           `xml:"AdTitle"`
	AdVerifications *AdVerifications `xml:"AdVerifications,omitempty"`
	Advertiser      string           `xml:"Advertiser,omitempty"`
	Category        []Category       `xml:"Category,omitempty"`
	Creatives       InLineCreatives  `xml:"Creatives"`
	Description     *CData           `xml:"Description,omitempty"`
	Expires         int              `xml:"Expires,omitempty"`
	Survey          *Survey          `xml:"Survey,omitempty"`
}

type InteractiveCreativeFile struct {
	Value            string      `xml:",cdata"`
	Type             string      `xml:"type,attr,omitempty"`
	APIFramework     string      `xml:"apiFramework,attr,omitempty"`
	VariableDuration NumericBool `xml:"variableDuration,attr,omitempty"`
}

type JavaScriptResource struct {
	Value           string      `xml:",cdata"`
	APIFramework    string      `xml:"apiFramework,attr,omitempty"`
	BrowserOptional NumericBool `xml:"browserOptional,attr,omitempty"`
}

type LinearBase struct {
	Icons          *Icons          `xml:"Icons,omitempty"`
	TrackingEvents *TrackingEvents `xml:"TrackingEvents,omitempty"`
	SkipOffset     SkipOffset      `xml:"skipoffset,attr,omitempty"`
}

type LinearInLine struct {
	LinearBase

	AdParameters *AdParameters `xml:"AdParameters,omitempty"`
	Duration     Duration      `xml:"Duration"`
	MediaFiles   MediaFiles    `xml:"MediaFiles"`
	VideoClicks  *VideoClicks  `xml:"VideoClicks,omitempty"`
}

type LinearWrapper struct {
	LinearBase

	VideoClicks *VideoClicks `xml:"VideoClicks,omitempty"`
}

type MediaFile struct {
	Value               string      `xml:",cdata"`
	ID                  string      `xml:"id,attr,omitempty"`
	Delivery            Delivery    `xml:"delivery,attr"`
	Type                string      `xml:"type,attr"`
	Width               int         `xml:"width,attr"`
	Height              int         `xml:"height,attr"`
	Codec               string      `xml:"codec,attr,omitempty"`
	Bitrate             int         `xml:"bitrate,attr,omitempty"`
	MinBitrate          int         `xml:"minBitrate,attr,omitempty"`
	MaxBitrate          int         `xml:"maxBitrate,attr,omitempty"`
	Scalable            NumericBool `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio NumericBool `xml:"maintainAspectRatio,attr,omitempty"`
	FileSize            int         `xml:"fileSize,attr,omitempty"`
	MediaType           string      `xml:"mediaType,attr,omitempty"`
	APIFramework        string      `xml:"apiFramework,attr,omitempty"`
}

type MediaFiles struct {
	ClosedCaptionFiles      *ClosedCaptionFiles       `xml:"ClosedCaptionFiles,omitempty"`
	MediaFile               []MediaFile               `xml:"MediaFile"`
	Mezzanine               []Mezzanine               `xml:"Mezzanine,omitempty"`
	InteractiveCreativeFile []InteractiveCreativeFile `xml:"InteractiveCreativeFile,omitempty"`
}

type Mezzanine struct {
	Value     string   `xml:",cdata"`
	Delivery  Delivery `xml:"delivery,attr"`
	Type      string   `xml:"type,attr"`
	Width     int      `xml:"width,attr"`
	Height    int      `xml:"height,attr"`
	Codec     string   `xml:"codec,attr,omitempty"`
	FileSize  int      `xml:"fileSize,attr,omitempty"`
	MediaType string   `xml:"mediaType,attr,omitempty"`
}

type Model string

const (
	CPCModel Model = "CPC"
	CPMModel Model = "CPM"
	CPEModel Model = "CPE"
	CPVModel Model = "CPV"
)

type NonLinearAdInLine struct {
	HTMLResource           []CData          `xml:"HTMLResource,omitempty"`
	IFrameResource         []CData          `xml:"IFrameResource,omitempty"`
	StaticResource         []StaticResource `xml:"StaticResource,omitempty"`
	AdParameters           *AdParameters    `xml:"AdParameters,omitempty"`
	NonLinearClickThrough  *CData           `xml:"NonLinearClickThrough,omitempty"`
	NonLinearClickTracking []CData          `xml:"NonLinearClickTracking,omitempty"`
	Width                  int              `xml:"width,attr"`
	Height                 int              `xml:"height,attr"`
	ExpandedWidth          int              `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int              `xml:"expandedHeight,attr,omitempty"`
	Scalable               NumericBool      `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    NumericBool      `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   Duration         `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework           string           `xml:"apiFramework,attr,omitempty"`
}

type NonLinearAds struct {
	TrackingEvents *TrackingEvents     `xml:"TrackingEvents,omitempty"`
	NonLinear      []NonLinearAdInLine `xml:"NonLinear,omitempty"`
}

// Offset must match the pattern `(\d{2}:[0-5]\d:[0-5]\d(\.\d\d\d)?|1?\d?\d(\.?\d)*%)`.
type Offset string

type Pricing struct {
	Value    float64  `xml:",cdata"`
	Model    Model    `xml:"model,attr"`
	Currency Currency `xml:"currency,attr"`
}

type RenderingMode string

const (
	DefaultRenderingMode    RenderingMode = "default"
	EndCardRenderingMode    RenderingMode = "end-card"
	ConcurrentRenderingMode RenderingMode = "concurrent"
)

type Required string

const (
	AllRequired  Required = "all"
	AnyRequired  Required = "any"
	NoneRequired Required = "none"
)

// SkipOffset must match the pattern `(\d{2}:[0-5]\d:[0-5]\d(\.\d\d\d)?|1?\d?\d(\.?\d)*%)`.
type SkipOffset string

type StaticResource struct {
	Value        string `xml:",cdata"`
	CreativeType string `xml:"creativeType,attr"`
}

type Survey struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr,omitempty"`
}

type Tracking struct {
	Value  string `xml:",cdata"`
	Event  string `xml:"event,attr"`
	Offset Offset `xml:"offset,attr,omitempty"`
}

type TrackingEventsVerification struct {
	Tracking []Tracking `xml:"Tracking,omitempty"`
}

type TrackingEvents struct {
	Tracking []Tracking `xml:"Tracking,omitempty"`
}

type UniversalAdID struct {
	Value      string `xml:",chardata"`
	IDRegistry string `xml:"idRegistry,attr"`
}

type Namespace string

const VASTNamespace Namespace = "http://www.iab.com/VAST"

var (
	ErrReadVAST      = errors.New("cannot read VAST")
	ErrUnmarshalVAST = errors.New("cannot unmarshal VAST")
	ErrMarshalVAST   = errors.New("cannot marshal VAST")
)

type Version string

const (
	VAST42Version Version = "4.2"
)

type VAST struct {
	Ad      []Ad      `xml:"Ad,omitempty"`
	Error   []CData   `xml:"Error,omitempty"`
	Version Version   `xml:"version,attr"`
	XMLNS   Namespace `xml:"xmlns,attr,omitempty"`
}

// New creates a new instance of VAST, sets the version to VAST42Version and the XML namespace to VASTNamespace.
func New() *VAST {
	return &VAST{
		Version: VAST42Version,
		XMLNS:   VASTNamespace,
	}
}

// Read creates a new instance of VAST and reads the content from an io.ReadCloser.
func Read(reader io.ReadCloser) (*VAST, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(ErrReadVAST, err)
	}

	_ = reader.Close()

	vast := &VAST{}
	if err := xml.Unmarshal(body, vast); err != nil {
		return nil, errors.Join(ErrUnmarshalVAST, err)
	}

	return vast, nil
}

// Bytes marshals the VAST to an XML document with indentations.
func (m *VAST) Bytes() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, errors.Join(ErrMarshalVAST, err)
	}

	return append([]byte(xml.Header), xmlData...), nil
}

type Verification struct {
	ExecutableResource     []ExecutableResource        `xml:"ExecutableResource,omitempty"`
	JavaScriptResource     []JavaScriptResource        `xml:"JavaScriptResource,omitempty"`
	TrackingEvents         *TrackingEventsVerification `xml:"TrackingEvents,omitempty"`
	VerificationParameters string                      `xml:"VerificationParameters,omitempty"`
	Vendor                 string                      `xml:"vendor,attr,omitempty"`
}

type ClickThrough struct {
	Value string `xml:",cdata"`
	ID    string `xml:"id,attr,omitempty"`
}

type VideoClicks struct {
	ClickTracking []CData      `xml:"ClickTracking,omitempty"`
	ClickThrough  ClickThrough `xml:"ClickThrough,omitempty"`
	CustomClick   []string     `xml:"CustomClick,omitempty"`
}

type ViewableImpression struct {
	Viewable         []CData `xml:"Viewable,omitempty"`
	NotViewable      []CData `xml:"NotViewable,omitempty"`
	ViewUndetermined []CData `xml:"ViewUndetermined,omitempty"`
}

type Wrapper struct {
	AdDefinitionBase

	AdVerifications          *AdVerifications      `xml:"AdVerifications,omitempty"`
	BlockedAdCategories      []BlockedAdCategories `xml:"BlockedAdCategories,omitempty"`
	Creatives                *Creatives            `xml:"Creatives,omitempty"`
	VASTAdTagURI             CData                 `xml:"VASTAdTagURI"`
	FollowAdditionalWrappers NumericBool           `xml:"followAdditionalWrappers,attr,omitempty"`
	AllowMultipleAds         NumericBool           `xml:"allowMultipleAds,attr,omitempty"`
	FallbackOnNoAd           NumericBool           `xml:"fallbackOnNoAd,attr,omitempty"`
}

// XPosition must match the pattern `([0-9]*|left|right)`.
type XPosition string

// YPosition must match the pattern `([0-9]*|top|bottom)`.
type YPosition string
