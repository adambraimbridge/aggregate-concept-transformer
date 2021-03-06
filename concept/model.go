package concept

import (
	"github.com/Financial-Times/aggregate-concept-transformer/s3"
)

type MembershipRole struct {
	RoleUUID        string `json:"membershipRoleUUID,omitempty"`
	InceptionDate   string `json:"inceptionDate,omitempty"`
	TerminationDate string `json:"terminationDate,omitempty"`
}

type ConcordedConcept struct {
	// Required fields
	PrefUUID  string `json:"prefUUID,omitempty"`
	PrefLabel string `json:"prefLabel,omitempty"`
	Type      string `json:"type,omitempty"`
	// Additional fields
	Aliases           []string `json:"aliases,omitempty"`
	ParentUUIDs       []string `json:"parentUUIDs,omitempty"`
	BroaderUUIDs      []string `json:"broaderUUIDs,omitempty"`
	RelatedUUIDs      []string `json:"relatedUUIDs,omitempty"`
	SupersededByUUIDs []string `json:"supersededByUUIDs,omitempty"`
	DescriptionXML    string   `json:"descriptionXML,omitempty"`
	ImageURL          string   `json:"_imageUrl,omitempty"`
	EmailAddress      string   `json:"emailAddress,omitempty"`
	FacebookPage      string   `json:"facebookPage,omitempty"`
	TwitterHandle     string   `json:"twitterHandle,omitempty"`
	ScopeNote         string   `json:"scopeNote,omitempty"`
	ShortLabel        string   `json:"shortLabel,omitempty"`
	// Brand
	Strapline string `json:"strapline,omitempty"`
	// Person
	Salutation string `json:"salutation,omitempty"`
	BirthYear  int    `json:"birthYear,omitempty"`
	// Financial Instrument
	FigiCode string `json:"figiCode,omitempty"`
	IssuedBy string `json:"issuedBy,omitempty"`
	// Membership
	InceptionDate    string           `json:"inceptionDate,omitempty"`
	MembershipRoles  []MembershipRole `json:"membershipRoles,omitempty"`
	OrganisationUUID string           `json:"organisationUUID,omitempty"`
	PersonUUID       string           `json:"personUUID,omitempty"`
	TerminationDate  string           `json:"terminationDate,omitempty"`
	// Organisation
	CountryCode            string   `json:"countryCode,omitempty"`
	CountryOfRisk          string   `json:"countryOfRisk,omitempty"`
	CountryOfIncorporation string   `json:"countryOfIncorporation,omitempty"`
	CountryOfOperations    string   `json:"countryOfOperations,omitempty"`
	FormerNames            []string `json:"formerNames,omitempty"`
	TradeNames             []string `json:"tradeNames,omitempty"`
	LeiCode                string   `json:"leiCode,omitempty"`
	PostalCode             string   `json:"postalCode,omitempty"`
	ProperName             string   `json:"properName,omitempty"`
	ShortName              string   `json:"shortName,omitempty"`
	YearFounded            int      `json:"yearFounded,omitempty"`
	IsDeprecated           bool     `json:"isDeprecated,omitempty"`
	// Location
	ISO31661 string `json:"iso31661,omitempty"`
	// Source representations
	SourceRepresentations []s3.Concept `json:"sourceRepresentations,omitempty"`
}
