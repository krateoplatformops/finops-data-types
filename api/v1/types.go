// +kubebuilder:object:generate=true
package config

import (
	prv1 "github.com/krateoplatformops/provider-runtime/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExporterScraperConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExporterScraperConfigSpec   `json:"spec,omitempty"`
	Status ExporterScraperConfigStatus `json:"status,omitempty"`
}

type ExporterScraperConfigSpec struct {
	// +optional
	ExporterConfig ExporterConfigSpec `yaml:"exporterConfig" json:"exporterConfig"`
	// +optional
	ScraperConfig ScraperConfigSpec `yaml:"scraperConfig" json:"scraperConfig"`
}

type ExporterScraperConfigStatus struct {
	prv1.ConditionedStatus `json:",inline"`

	ActiveExporter corev1.ObjectReference `json:"active,omitempty"`
	ConfigMap      corev1.ObjectReference `json:"configMaps,omitempty"`
	Service        corev1.ObjectReference `json:"services,omitempty"`
}

type ExporterConfigSpec struct {
	Provider ObjectRef `yaml:"provider" json:"provider"`
	Url      string    `yaml:"url" json:"url"`
	// +optional
	UrlParsed             string `yaml:"urlParsed" json:"urlParsed,omitempty"`
	RequireAuthentication bool   `yaml:"requireAuthentication" json:"requireAuthentication"`
	// +kubebuilder:validation:Pattern=`^(\bcert-file\b)|(\bbearer-token\b)$`
	AuthenticationMethod string `yaml:"authenticationMethod" json:"authenticationMethod"`
	// +optional
	BearerToken SecretKeySelector `yaml:"bearerToken" json:"bearerToken"`
	// +kubebuilder:validation:Pattern=`^(\b[Cc]ost\b)|(\b[Rr]esource\b)$`
	// +kubebuilder:default=cost
	// +optional
	MetricType           string            `yaml:"metricType" json:"metricType"`
	PollingIntervalHours int               `yaml:"pollingIntervalHours" json:"pollingIntervalHours"`
	AdditionalVariables  map[string]string `yaml:"additionalVariables" json:"additionalVariables"`
}

type ScraperConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScraperConfigSpec   `json:"spec,omitempty"`
	Status ScraperConfigStatus `json:"status,omitempty"`
}

type ScraperConfigSpec struct {
	TableName            string `yaml:"tableName" json:"tableName"`
	PollingIntervalHours int    `yaml:"pollingIntervalHours" json:"pollingIntervalHours"`
	// +optional
	Url string `yaml:"url" json:"url,omitempty"`
	// +kubebuilder:default=cost
	// +optional
	MetricType               string    `yaml:"metricType" json:"metricType"`
	ScraperDatabaseConfigRef ObjectRef `yaml:"scraperDatabaseConfigRef" json:"scraperDatabaseConfigRef"`
}

type ScraperConfigStatus struct {
	prv1.ConditionedStatus `json:",inline"`
	ActiveScraper          corev1.ObjectReference `json:"active,omitempty"`
	ConfigMap              corev1.ObjectReference `json:"configMaps,omitempty"`
}

type FocusConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FocusConfigSpec   `json:"spec,omitempty"`
	Status FocusConfigStatus `json:"status,omitempty"`
}

type FocusConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FocusConfig `json:"items"`
}

// FocusConfigSpec defines the desired state of FocusConfig
type FocusConfigSpec struct {
	FocusSpec FocusSpec `yaml:"focusSpec" json:"focusSpec"`
	// +optional
	ScraperConfig ScraperConfigSpec `yaml:"scraperConfig" json:"scraperConfig"`
}

// FocusConfigStatus defines the observed state of FocusConfig
type FocusConfigStatus struct {
	prv1.ConditionedStatus `json:",inline"`
	GroupKey               string `json:"groupKey,omitempty"`
}

type FocusSpec struct {
	// +optional
	AvailabilityZone string            `yaml:"availabilityZone" json:"availabilityZone"`
	BilledCost       resource.Quantity `yaml:"billedCost" json:"billedCost"`
	// +optional
	BillingAccountId string `yaml:"billingAccountId" json:"billingAccountId"`
	// +optional
	BillingAccountName string `yaml:"billingAccountName" json:"billingAccountName"`
	// +kubebuilder:validation:Pattern=`^[A-Z]{3}$`
	BillingCurrency    string      `yaml:"billingCurrency" json:"billingCurrency"`
	BillingPeriodEnd   metav1.Time `yaml:"billingPeriodEnd" json:"billingPeriodEnd"`
	BillingPeriodStart metav1.Time `yaml:"billingPeriodStart" json:"billingPeriodStart"`
	// +optional
	CapacityReservationId string `yaml:"capacityReservationId" json:"capacityReservationId"`
	// +kubebuilder:validation:Pattern=`(\b[Uu]used\b)|(\b[Uu]nused\b)`
	CapacityReservationStatus bool `yaml:"capacityReservationStatus" json:"capacityReservationStatus"`
	// +kubebuilder:validation:Pattern=`(\b[Aa]djustment\b)|(\b[Pp]urchase\b)|(\b[Tt]ax\b)|(\b[Uu]sage\b)`
	ChargeCategory string `yaml:"chargeCategory" json:"chargeCategory"`
	// +kubebuilder:validation:Pattern=`(\b[Cc]orrection\b)`
	// +optional
	ChargeClass       string `yaml:"chargeClass" json:"chargeClass"`
	ChargeDescription string `yaml:"chargeDescription" json:"chargeDescription"`
	// +kubebuilder:validation:Pattern=`(\b[Oo]ne-{0,1}[Tt]ime\b)|(\b[Rr]ecurring\b)|(\b[Uu]sage-{0,1}[Bb]ased\b)`
	// +optional
	ChargeFrequency   string      `yaml:"chargeFrequency" json:"chargeFrequency"`
	ChargePeriodEnd   metav1.Time `yaml:"chargePeriodEnd" json:"chargePeriodEnd"`
	ChargePeriodStart metav1.Time `yaml:"chargePeriodStart" json:"chargePeriodStart"`
	// +kubebuilder:validation:Pattern=`(\b[Ss]spend\b)|(\b[Uu]sage\b)`
	// +optional
	CommitmentDiscountCategory string `yaml:"commitmentDiscountCategory" json:"commitmentDiscountCategory"`
	// +optional
	CommitmentDiscountId string `yaml:"commitmentDiscoutId" json:"commitmentDiscoutId"`
	// +optional
	CommitmentDiscountType string `yaml:"commitmentDiscountType" json:"commitmentDiscountType"`
	// +kubebuilder:validation:Pattern=`(\b[Uu]sed\b)|(\b[Uu]nused\b)`
	// +optional
	CommitmentDiscountStatus string `yaml:"commitmentDiscountStatus" json:"commitmentDiscountStatus"`
	// +optional
	CommitmentDiscountName string `yaml:"commitmentDiscountName" json:"commitmentDiscountName"`
	// +optional
	CommitmentDiscountQuantity resource.Quantity `yaml:"commitmentDiscountQuantity" json:"commitmentDiscountQuantity"`
	// +optional
	CommitmentDiscountUnit string            `yaml:"commitmentDiscountUnit" json:"commitmentDiscountUnit"`
	ConsumedQuantity       resource.Quantity `yaml:"consumedQuantity" json:"consumedQuantity"`
	ConsumedUnit           string            `yaml:"consumedUnit" json:"consumedUnit"`
	ContractedCost         resource.Quantity `yaml:"contractedCost" json:"contractedCost"`
	// +optional
	ContractedUnitCost resource.Quantity `yaml:"contractedUnitCost" json:"contractedUnitCost"`
	// +optional
	EffectiveCost     resource.Quantity `yaml:"effectiveCost" json:"effectiveCost"`
	InvoiceIssuerName string            `yaml:"invoiceIssuerName" json:"invoiceIssuerName"`
	// +optional
	ListCost resource.Quantity `yaml:"listCost" json:"listCost"`
	// +optional
	ListUnitPrice resource.Quantity `yaml:"listUnitPrice" json:"listUnitPrice"`
	// +kubebuilder:validation:Pattern=`(\b[Oo]n-{0,1}[Dd]emand\b)|(\b[Dd]ynamic\b)|(\b[Cc]ommitment-{0,1}[Bb]ased\b)|(\b[Oo]ther\b)`
	// +optional
	PricingCategory string `yaml:"pricingCategory" json:"pricingCategory"`
	// +optional
	PricingQuantity resource.Quantity `yaml:"pricingQuantity" json:"pricingQuantity"`
	// +optional
	PricingUnit string `yaml:"pricingUnit" json:"pricingUnit"`
	// +optional
	ProviderName string `yaml:"providerName" json:"providerName"`
	// +optional
	PublisherName string `yaml:"publisherName" json:"publisherName"`
	// +optional
	RegionId string `yaml:"regionId" json:"regionId"`
	// +optional
	RegionName string `yaml:"regionName" json:"regionName"`
	// +optional
	ResourceId string `yaml:"resourceId" json:"resourceId"`
	// +optional
	ResourceName string `yaml:"resourceName" json:"resourceName"`
	// +optional
	ResourceType string `yaml:"resourceType" json:"resourceType"`
	// +kubebuilder:validation:Pattern=`(\bAI and Machine Learning\b)|(\bAnalytics\b)|(\bBusiness\b)|(\bCompute\b)|(\bDatabases\b)|(\bDeveloper Tools\b)|(\bMulticloud\b)|(\bIdentity\b)|(\bIntegration\b)|(\bInternet of Things\b)|(\bManagement and Governance\b)|(\bMedia\b)|(\bMigration\b)|(\bMobile\b)|(\bNetworking\b)|(\bSecurity\b)|(\bStorage\b)|(\bWeb\b)|(\bOther\b)`
	ServiceCategory string `yaml:"serviceCategory" json:"serviceCategory"`
	// +optional
	ServiceName        string `yaml:"serviceName" json:"serviceName"`
	ServiceSubcategory string `yaml:"serviceSubcategory" json:"serviceSubcategory"`
	// +optional
	SkuId string `yaml:"skuId" json:"skuId"`
	// +optional
	SkuMeter string `yaml:"skuMeter" json:"skuMeter"`
	// +optional
	SkuPriceDetails []TagsType `yaml:"skuPriceDetails" json:"skuPriceDetails"`
	// +optional
	SkuPriceId string `yaml:"skuPriceId" json:"skuPriceId"`
	// +optional
	SubAccountId string `yaml:"subAccountId" json:"subAccountId"`
	// +optional
	SubAccountName string `yaml:"subAccountName" json:"subAccountName"`
	// +optional
	Tags []TagsType `yaml:"tags" json:"tags"`
}

type TagsType struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

type ObjectRef struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

// A SecretKeySelector is a reference to a secret key in an arbitrary namespace.
type SecretKeySelector struct {
	// Name of the referenced object.
	Name string `json:"name"`

	// Namespace of the referenced object.
	Namespace string `json:"namespace"`

	// The key to select.
	Key string `json:"key"`
}
