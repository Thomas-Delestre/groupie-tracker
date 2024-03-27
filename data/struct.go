package data

// API Artist
type Artist_data struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Image         string   `json:"image"`
	Members       []string `json:"members"`
	Creationdate  int      `json:"creationDate"`
	Firstalbum    string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	Concert_dates string   `json:"concertDates"`
	Relations     string   `json:"relations"`
}

// API détails des Locations de chaque Artists
type Locations_data struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// API détails des Dates de concert de chaque Artists
type Concert_dates_data struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// API détails des Relation dates/lieux de chaque Artists
type Relations_data struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Global struct
type All_Data struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	Members        []string            `json:"members"`
	Creationdate   int                 `json:"creationDate"`
	Firstalbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	Dates          []string            `json:"dates"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Filters value stockage
type Filters_values struct {
	CD_start int
	CD_end   int
	FA_start int
	FA_end   int
	Checkbox [7]int
	Location string
}
