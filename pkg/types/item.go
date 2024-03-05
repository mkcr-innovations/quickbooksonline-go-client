package types

type Item struct {
	BaseEntity
	ItemCategoryType     *ItemCategoryTypeEnum `json:"ItemCategoryType,omitempty"` // Required, minorVersion: 3
	Name                 string                `json:"Name"`
	InvStartDate         string                `json:"InvStartDate"`                   // Conditionally required
	Type                 *string               `json:"Type,omitempty"`                 // Conditionally required, minorVersion: specified
	QtyOnHand            *float64              `json:"QtyOnHand,omitempty"`            // Conditionally required
	AssetAccountRef      *Ref                  `json:"AssetAccountRef,omitempty"`      // Conditionally required
	Sku                  *string               `json:"Sku,omitempty"`                  // Optional, minorVersion: 4
	SalesTaxIncluded     *bool                 `json:"SalesTaxIncluded,omitempty"`     // Optional
	TrackQtyOnHand       *bool                 `json:"TrackQtyOnHand,omitempty"`       // Optional
	SalesTaxCodeRef      *Ref                  `json:"SalesTaxCodeRef,omitempty"`      // Optional
	ClassRef             *Ref                  `json:"ClassRef,omitempty"`             // Optional, minorVersion: 41
	Source               *SourceEnum           `json:"Source,omitempty"`               // Optional, minorVersion: 59
	PurchaseTaxIncluded  *bool                 `json:"PurchaseTaxIncluded,omitempty"`  // Optional
	Description          *string               `json:"Description,omitempty"`          // Optional
	AbatementRate        *float64              `json:"AbatementRate,omitempty"`        // Optional, minorVersion: 3
	SubItem              *bool                 `json:"SubItem,omitempty"`              // Optional
	Taxable              *bool                 `json:"Taxable,omitempty"`              // Optional
	UQCDisplayText       *string               `json:"UQCDisplayText,omitempty"`       // Optional, minorVersion: 33
	ReorderPoint         *float64              `json:"ReorderPoint,omitempty"`         // Optional
	PurchaseDesc         *string               `json:"PurchaseDesc,omitempty"`         // Optional
	MetaData             *ModificationMetaData `json:"MetaData,omitempty"`             // Optional
	PrefVendorRef        *Ref                  `json:"PrefVendorRef,omitempty"`        // Optional, minorVersion: 31
	Active               *bool                 `json:"Active,omitempty"`               // Optional
	UQCId                *string               `json:"UQCId,omitempty"`                // Optional, minorVersion: 33
	ReverseChargeRate    *float64              `json:"ReverseChargeRate,omitempty"`    // Optional, minorVersion: 3
	PurchaseTaxCodeRef   *Ref                  `json:"PurchaseTaxCodeRef,omitempty"`   // Optional
	ServiceType          *ServiceTypeEnum      `json:"ServiceType,omitempty"`          // Optional, minorVersion: 3
	PurchaseCost         *float64              `json:"PurchaseCost,omitempty"`         // Optional
	ParentRef            *Ref                  `json:"ParentRef,omitempty"`            // Optional
	UnitPrice            *float64              `json:"UnitPrice,omitempty"`            // Optional
	FullyQualifiedName   string                `json:"FullyQualifiedName"`             // Read-only, system defined
	ExpenseAccountRef    *Ref                  `json:"ExpenseAccountRef,omitempty"`    // Optional for France locales
	Level                int                   `json:"Level"`                          // Read-only, system defined
	IncomeAccountRef     *Ref                  `json:"IncomeAccountRef,omitempty"`     // Conditionally required
	TaxClassificationRef *Ref                  `json:"TaxClassificationRef,omitempty"` // Optional, minorVersion: 34
}

type ItemPaginatedResponse struct {
	BasePaginatedResponse
	Item []Item `json:"Item"`
}

type ItemResponse struct {
	BaseResponse
	Item Item `json:"Item"`
}

type ItemCategoryTypeEnum string

const (
	ProductItemCategoryType ItemCategoryTypeEnum = "Product"
	ServiceItemCategoryType ItemCategoryTypeEnum = "Service"
)

type ServiceTypeEnum string

const (
	ADVTServiceType                           ServiceTypeEnum = "ADVT"
	AIRPORTSERVICESServiceType                ServiceTypeEnum = "AIRPORTSERVICES"
	AIRTRANSPORTServiceType                   ServiceTypeEnum = "AIRTRANSPORT"
	AIRTRVLAGNTServiceType                    ServiceTypeEnum = "AIRTRVLAGNT"
	ARCHITECTASSTMGMTServiceType              ServiceTypeEnum = "ARCHITECT"
	ASSTMGMTServiceType                       ServiceTypeEnum = "ASSTMGMT"
	ATMMAINTENANCEServiceType                 ServiceTypeEnum = "ATMMAINTENANCE"
	AUCTIONSERVServiceType                    ServiceTypeEnum = "AUCTIONSERV"
	AUTHSERSTServiceType                      ServiceTypeEnum = "AUTHSERST"
	BANKANDFINServiceType                     ServiceTypeEnum = "BANKANDFIN"
	BEAUTYPARLORServiceType                   ServiceTypeEnum = "BEAUTYPARLOR"
	BROADCASTServiceType                      ServiceTypeEnum = "BROADCAST"
	BUSINESSAUXServiceType                    ServiceTypeEnum = "BUSINESSAUX"
	BUSINESSEXHIBITIONServiceType             ServiceTypeEnum = "BUSINESSEXHIBITION"
	BUSINESSSUPPORTSERVServiceType            ServiceTypeEnum = "BUSINESSSUPPORTSERV"
	CAServiceType                             ServiceTypeEnum = "CA"
	CABLEOPTRServiceType                      ServiceTypeEnum = "CABLEOPTR"
	CARGOHANDServiceType                      ServiceTypeEnum = "CARGOHAND"
	CLEANINGSERVServiceType                   ServiceTypeEnum = "CLEANINGSERV"
	CLEARANDFORWServiceType                   ServiceTypeEnum = "CLEARANDFORW"
	CLUBSANDASSSERVICEServiceType             ServiceTypeEnum = "CLUBSANDASSSERVICE"
	COMMCOACHORTRAININGServiceType            ServiceTypeEnum = "COMMCOACHORTRAINING"
	CONSENGServiceType                        ServiceTypeEnum = "CONSENG"
	CONSTRCOMMERCIALCOMPLEXServiceType        ServiceTypeEnum = "CONSTRCOMMERCIALCOMPLEX"
	CONTAINERRAILTRANSServiceType             ServiceTypeEnum = "CONTAINERRAILTRANS"
	CONVSERVServiceType                       ServiceTypeEnum = "CONVSERV"
	COSTACCServiceType                        ServiceTypeEnum = "COSTACC"
	COURIERServiceType                        ServiceTypeEnum = "COURIER"
	CREDITCARDServiceType                     ServiceTypeEnum = "CREDITCARD"
	CREDITRATAGNCYServiceType                 ServiceTypeEnum = "CREDITRATAGNCY"
	CRUISESHIPTOURServiceType                 ServiceTypeEnum = "CRUISESHIPTOUR"
	CSServiceType                             ServiceTypeEnum = "CS"
	CUSHOUSEAGServiceType                     ServiceTypeEnum = "CUSHOUSEAG"
	DESIGNSERVServiceType                     ServiceTypeEnum = "DESIGNSERV"
	DEVELOPSUPPLYCONTENTServiceType           ServiceTypeEnum = "DEVELOPSUPPLYCONTENT"
	DREDGINGServiceType                       ServiceTypeEnum = "DREDGING"
	DRYCLEANINGServiceType                    ServiceTypeEnum = "DRYCLEANING"
	ERECTIONCOMMORINSTALLServiceType          ServiceTypeEnum = "ERECTIONCOMMORINSTALL"
	EVENTMGMTServiceType                      ServiceTypeEnum = "EVENTMGMT"
	FASHIONDESServiceType                     ServiceTypeEnum = "FASHIONDES"
	FOREXBROKINGServiceType                   ServiceTypeEnum = "FOREXBROKING"
	FORWARDCONTRACTServiceType                ServiceTypeEnum = "FORWARDCONTRACT"
	FRANCHISESERVServiceType                  ServiceTypeEnum = "FRANCHISESERV"
	GENERALINSURANCEServiceType               ServiceTypeEnum = "GENERALINSURANCE"
	GOODSTRANSPORTServiceType                 ServiceTypeEnum = "GOODSTRANSPORT"
	HEALTHCLUBANDFITNESSServiceType           ServiceTypeEnum = "HEALTHCLUBANDFITNESS"
	INFORMATIONSERVServiceType                ServiceTypeEnum = "INFORMATIONSERV"
	INSURAUXServiceType                       ServiceTypeEnum = "INSURAUX"
	INTDECServiceType                         ServiceTypeEnum = "INTDEC"
	INTELLECTUALPROPERTYServiceType           ServiceTypeEnum = "INTELLECTUALPROPERTY"
	INTERNATIONALAIRTRAVELServiceType         ServiceTypeEnum = "INTERNATIONALAIRTRAVEL"
	INTERNETCAFEServiceType                   ServiceTypeEnum = "INTERNETCAFE"
	INTERNETTELEPHONYServiceType              ServiceTypeEnum = "INTERNETTELEPHONY"
	LIFEINSServiceType                        ServiceTypeEnum = "LIFEINS"
	MAILLISTCOMPILEServiceType                ServiceTypeEnum = "MAILLISTCOMPILE"
	MANDAPKEEPERServiceType                   ServiceTypeEnum = "MANDAPKEEPER"
	MANPWRRECRUITServiceType                  ServiceTypeEnum = "MANPWRRECRUIT"
	MGMTCONSULServiceType                     ServiceTypeEnum = "MGMTCONSUL"
	MGMTMAINTREPAIRServiceType                ServiceTypeEnum = "MGMTMAINTREPAIR"
	MININGOILServiceType                      ServiceTypeEnum = "MININGOIL"
	MKTRESAGNCYServiceType                    ServiceTypeEnum = "MKTRESAGNCY"
	ONLINEINFORMRETRIEVALServiceType          ServiceTypeEnum = "ONLINEINFORMRETRIEVAL"
	OPINIONPOLLServiceType                    ServiceTypeEnum = "OPINIONPOLL"
	OUTDOORCATERINGServiceType                ServiceTypeEnum = "OUTDOORCATERING"
	PACKAGINGSERVServiceType                  ServiceTypeEnum = "PACKAGINGSERV"
	PANDALSHAMIANAServiceType                 ServiceTypeEnum = "PANDALSHAMIANA"
	PHOTOGRAPHYServiceType                    ServiceTypeEnum = "PHOTOGRAPHY"
	PORTServiceType                           ServiceTypeEnum = "PORT"
	PORTSERServiceType                        ServiceTypeEnum = "PORTSER"
	PROCESSCLEARHOUSEServiceType              ServiceTypeEnum = "PROCESSCLEARHOUSE"
	PUBLICRELATIONMGMTServiceType             ServiceTypeEnum = "PUBLICRELATIONMGMT"
	RAILTRAVELAGNTServiceType                 ServiceTypeEnum = "RAILTRAVELAGNT"
	REALESTAGTServiceType                     ServiceTypeEnum = "REALESTAGT"
	RECOVERYAGENTSServiceType                 ServiceTypeEnum = "RECOVERYAGENTS"
	REGISTRARSERVServiceType                  ServiceTypeEnum = "REGISTRARSERV"
	RENTACABServiceType                       ServiceTypeEnum = "RENTACAB"
	RENTINGIMMOVABLEPROPServiceType           ServiceTypeEnum = "RENTINGIMMOVABLEPROP"
	RESIDENTIALCOMPLEXCONSTServiceType        ServiceTypeEnum = "RESIDENTIALCOMPLEXCONST"
	SALEOFSPACEFORADVTServiceType             ServiceTypeEnum = "SALEOFSPACEFORADVT"
	SCANDTECHCONSULServiceType                ServiceTypeEnum = "SCANDTECHCONSUL"
	SECAGServiceType                          ServiceTypeEnum = "SECAG"
	SERVICESPROVIDEDFORTRANSACTIONServiceType ServiceTypeEnum = "SERVICESPROVIDEDFORTRANSACTION"
	SHARETRANSFERSERVServiceType              ServiceTypeEnum = "SHARETRANSFERSERV"
	SHIPMGMTServiceType                       ServiceTypeEnum = "SHIPMGMT"
	SITEPREPServiceType                       ServiceTypeEnum = "SITEPREP"
	SOUNDRECORDServiceType                    ServiceTypeEnum = "SOUNDRECORD"
	SPONSORSHIPServiceType                    ServiceTypeEnum = "SPONSORSHIP"
	STAGServiceType                           ServiceTypeEnum = "STAG"
	STOCKBROKINGServiceType                   ServiceTypeEnum = "STOCKBROKING"
	STOCKEXCHGSERVServiceType                 ServiceTypeEnum = "STOCKEXCHGSERV"
	STORANDWAREHOUSINGServiceType             ServiceTypeEnum = "STORANDWAREHOUSING"
	SUPPLYTANGIBLEGOODSServiceType            ServiceTypeEnum = "SUPPLYTANGIBLEGOODS"
	SURVEYANDMAPMAKINGServiceType             ServiceTypeEnum = "SURVEYANDMAPMAKING"
	SURVEYMINERALSServiceType                 ServiceTypeEnum = "SURVEYMINERALS"
	TECHINSPECTIONServiceType                 ServiceTypeEnum = "TECHINSPECTION"
	TECHTESTINGServiceType                    ServiceTypeEnum = "TECHTESTING"
	TELECOMMUNICATIONSERVServiceType          ServiceTypeEnum = "TELECOMMUNICATIONSERV"
	TELEVISIONANDRADIOServiceType             ServiceTypeEnum = "TELEVISIONANDRADIO"
	TOUROPServiceType                         ServiceTypeEnum = "TOUROP"
	TRANSPORTPIPELINEServiceType              ServiceTypeEnum = "TRANSPORTPIPELINE"
	TRAVELAGENTServiceType                    ServiceTypeEnum = "TRAVELAGENT"
	ULIPMANAGEMENTServiceType                 ServiceTypeEnum = "ULIPMANAGEMENT"
	UNDERWRITERServiceType                    ServiceTypeEnum = "UNDERWRITER"
	VIDEOTAPEPRODServiceType                  ServiceTypeEnum = "VIDEOTAPEPROD"
	WORKSCONTRACTServiceType                  ServiceTypeEnum = "WORKSCONTRACT"
)
