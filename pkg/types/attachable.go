package types

type Attachable struct {
	BaseEntity
	FileName                 *string               `json:"FileName,omitempty"`       // Conditionally required, max 1000 chars
	Note                     *string               `json:"Note,omitempty"`           // Conditionally required, max 2000 chars
	Category                 *string               `json:"Category,omitempty"`       // Optional, max 100 chars
	ContentType              *string               `json:"ContentType,omitempty"`    // Optional, max 100 chars
	PlaceName                *string               `json:"PlaceName,omitempty"`      // Optional, max 2000 chars
	AttachableRef            []AttachableRef       `json:"AttachableRef,omitempty"`  // Optional
	Long                     *string               `json:"Long,omitempty"`           // Optional, max 100 chars
	Tag                      *string               `json:"Tag,omitempty"`            // Optional, max 2000 chars
	Lat                      *string               `json:"Lat,omitempty"`            // Optional, max 100 chars
	MetaData                 *ModificationMetaData `json:"MetaData,omitempty"`       // Optional
	FileAccessUri            string                `json:"FileAccessUri"`            // Read only, system defined
	Size                     float64               `json:"Size"`                     // Read only, system defined
	ThumbnailFileAccessUri   string                `json:"ThumbnailFileAccessUri"`   // Read only, system defined
	TempDownloadUri          string                `json:"TempDownloadUri"`          // Read only, system defined
	ThumbnailTempDownloadUri string                `json:"ThumbnailTempDownloadUri"` // Read only, system defined
}

type AttachableRef struct {
	IncludeOnSend *bool         `json:"IncludeOnSend,omitempty"` // Optional
	LineInfo      *string       `json:"LineInfo,omitempty"`      // Optional
	NoRefOnly     *bool         `json:"NoRefOnly,omitempty"`     // Optional
	CustomField   []CustomField `json:"CustomField,omitempty"`   // Optional
	Inactive      *bool         `json:"Inactive,omitempty"`      // Optional
	EntityRef     *Ref          `json:"EntityRef,omitempty"`     // Optional
}

type AttachablePaginatedResponse struct {
	BasePaginatedResponse
	Attachable []Attachable `json:"Attachable"`
}

type AttachableResponse struct {
	BaseResponse
	Attachable Attachable `json:"Attachable"`
}

type AttachableCategory string

const (
	AttachableCategoryContactPhoto AttachableCategory = "Contact Photo"
	AttachableCategoryDocument     AttachableCategory = "Document"
	AttachableCategoryImage        AttachableCategory = "Image"
	AttachableCategoryReceipt      AttachableCategory = "Receipt"
	AttachableCategorySignature    AttachableCategory = "Signature"
	AttachableCategorySound        AttachableCategory = "Sound"
	AttachableCategoryOther        AttachableCategory = "Other"
)
