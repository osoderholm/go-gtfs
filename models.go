package gtfs

// GTFS -
type GTFS struct {
	Path           string // The path to the containing directory
	Agencies       []Agency
	Routes         []Route
	Stops          []Stop
	StopsTimes     []StopTime
	Trips          []Trip
	Calendars      []Calendar
	CalendarDates  []CalendarDate
	Transfers      []Transfer
	Shapes         []Shape
	FareAttributes []FareAttribute
	FareRules      []FareRule
	Frequencies    []Frequency
	Pathways       []Pathway
	Levels         []Level
	FeedInfo       []FeedInfo
	Translations   []Translation
	Attributions   []Attribution
}

// Route -
type Route struct {
	ID        string `csv:"route_id"`
	AgencyID  string `csv:"agency_id"`
	ShortName string `csv:"route_short_name"`
	LongName  string `csv:"route_long_name"`
	Type      int    `csv:"route_type"`
	Desc      string `csv:"route_url"`
	URL       string `csv:"route_desc"`
	Color     string `csv:"route_color"`
	TextColor string `csv:"route_text_color"`
	SortOrder int    `csv:"route_sort_order"`
}

// Trip -
type Trip struct {
	ID                   string `csv:"trip_id"`
	ShortName            string `csv:"trip_short_name"`
	RouteID              string `csv:"route_id"`
	ServiceID            string `csv:"service_id"`
	ShapeID              string `csv:"shape_id"`
	DirectionID          string `csv:"direction_id"`
	Headsign             string `csv:"trip_headsign"`
	BlockID              string `csv:"block_id"`
	WheelchairAccessible int    `csv:"wheelchair_accessible"`
	BikesAllowed         int    `csv:"bikes_allowed"`
}

// Stop -
type Stop struct {
	ID                 string  `csv:"stop_id"`
	Code               string  `csv:"stop_code"`
	Name               string  `csv:"stop_name"`
	Description        string  `csv:"stop_desc"`
	Latitude           float64 `csv:"stop_lat"`
	Longitude          float64 `csv:"stop_lon"`
	ZoneID             string  `csv:"zone_id"`
	URL                string  `csv:"stop_url"`
	Timezone           string  `csv:"stop_timezone"`
	Type               string  `csv:"location_type"`
	Parent             string  `csv:"parent_station"`
	WheelchairBoarding int     `csv:"wheelchair_boarding"`
	LevelID            string  `csv:"level_id"`
	PlatformCode       string  `csv:"platform_code"`
}

// StopTime -
type StopTime struct {
	StopID            string  `csv:"stop_id"`
	StopSeq           int     `csv:"stop_sequence"`
	StopHeadSign      string  `csv:"stop_headsign"`
	TripID            string  `csv:"trip_id"`
	ShapeDistTraveled float64 `csv:"shape_dist_traveled"`
	DepartureTime     string  `csv:"departure_time"`
	ArrivalTime       string  `csv:"arrival_time"`
	PickupType        int     `csv:"pickup_type"`
	DropOffType       int     `csv:"drop_off_type"`
	TimePoint         int     `csv:"timepoint"`
}

// Calendar -
type Calendar struct {
	ServiceID string `csv:"service_id"`
	Monday    int    `csv:"monday"`
	Tuesday   int    `csv:"tuesday"`
	Wednesday int    `csv:"wednesday"`
	Thursday  int    `csv:"thursday"`
	Friday    int    `csv:"friday"`
	Saturday  int    `csv:"saturday"`
	Sunday    int    `csv:"sunday"`
	StartDate string `csv:"start_date"`
	EndDate   string `csv:"end_date"`
}

// CalendarDate -
type CalendarDate struct {
	ServiceID     string `csv:"service_id"`
	Date          string `csv:"date"`
	ExceptionType int    `csv:"exception_type"`
}

// Transfer -
type Transfer struct {
	FromStopID string `csv:"from_stop_id"`
	ToStopID   string `csv:"to_stop_id"`
	Type       int    `csv:"transfer_type"`
	MinTime    int    `csv:"min_transfer_time"`
}

// Agency -
type Agency struct {
	ID       string `csv:"agency_id"`
	Name     string `csv:"agency_name"`
	URL      string `csv:"agency_url"`
	Timezone string `csv:"agency_timezone"`
	Language string `csv:"agency_lang"`
	Phone    string `csv:"agency_phone"`
	FareURL  string `csv:"agency_fare_url"`
	Email    string `csv:"agency_email"`
}

// Shape -
type Shape struct {
	ID               string  `csv:"shape_id"`
	Sequence         int     `csv:"shape_pt_sequence"`
	Latitude         float64 `csv:"shape_pt_lat"`
	Longitude        float64 `csv:"shape_pt_lon"`
	DistanceTraveled float64 `csv:"shape_dist_traveled"`
}

// FareAttribute -
type FareAttribute struct {
	ID               string  `csv:"fare_id"`
	Price            float64 `csv:"price"`
	CurrencyType     string  `csv:"currency_type"`
	PaymentMethod    int     `csv:"payment_method"`
	Transfers        int     `csv:"transfers"`
	AgencyID         string  `csv:"agency_id"`
	TransferDuration int     `csv:"transfer_duration"`
}

// FareRule -
type FareRule struct {
	FareID        string `csv:"fare_id"`
	RouteID       string `csv:"route_id"`
	OriginID      string `csv:"origin_id"`
	DestinationID string `csv:"destination_id"`
	ContainsID    string `csv:"contains_id"`
}

// Frequency -
type Frequency struct {
	TripID      string `csv:"trip_id"`
	StartTime   string `csv:"start_time"`
	EndTime     string `csv:"end_time"`
	HeadwaySecs int    `csv:"headway_secs"`
	ExactTimes  int    `csv:"exact_times"`
}

// Pathway -
type Pathway struct {
	ID                   string  `csv:"pathway_id"`
	FromStopID           string  `csv:"from_stop_id"`
	ToStopID             string  `csv:"to_stop_id"`
	PathwayMode          int     `csv:"pathway_mode"`
	IsBidirectional      int     `csv:"is_bidirectional"`
	Length               float64 `csv:"length"`
	TraversalTime        int     `csv:"traversal_time"`
	StairCount           int     `csv:"stair_count"`
	MaxSlope             float64 `csv:"max_slope"`
	MinWidth             float64 `csv:"min_width"`
	SignpostedAs         string  `csv:"signposted_as"`
	ReversedSignpostedAs string  `csv:"reversed_signposted_as"`
}

// Levels -
type Level struct {
	ID    string  `csv:"level_id"`
	Index float64 `csv:"level_index"`
	Name  string  `csv:"level_name"`
}

// FeedInfo -
type FeedInfo struct {
	PublisherName    string `csv:"feed_publisher_name"`
	PublisherURL     string `csv:"feed_publisher_url"`
	FeedLang         string `csv:"feed_lang"`
	DefaultLang      string `csv:"default_lang"`
	FeedStartDate    string `csv:"feed_start_date"`
	FeedEndDate      string `csv:"feed_end_date"`
	FeedVersion      string `csv:"feed_version"`
	FeedContactEmail string `csv:"feed_contact_email"`
	FeedContactURL   string `csv:"feed_contact_url"`
}

// Translation -
type Translation struct {
	TableName   string `csv:"table_name"`
	FieldName   string `csv:"field_name"`
	Language    string `csv:"language"`
	Translation string `csv:"translation"`
	RecordID    string `csv:"record_id"`
	RecordSubID string `csv:"record_sub_id"`
	FieldValue  string `csv:"field_value"`
}

// Attribution -
type Attribution struct {
	ID               string `csv:"attribution_id"`
	AgencyID         string `csv:"agency_id"`
	RouteID          string `csv:"route_id"`
	TripID           string `csv:"trip_id"`
	OrganizationName string `csv:"organization_name"`
	IsProducer       int    `csv:"is_producer"`
	IsOperator       int    `csv:"is_operator"`
	IsAuthority      int    `csv:"is_authority"`
	URL              string `csv:"attribution_url"`
	Email            string `csv:"attribution_email"`
	Phone            string `csv:"attribution_phone"`
}
