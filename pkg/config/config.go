package config

import (
	corev1 "k8s.io/api/core/v1"
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
	ExporterConfig ExporterConfig `yaml:"exporterConfig" json:"exporterConfig"`
	// +optional
	ScraperConfig ScraperConfig `yaml:"scraperConfig" json:"scraperConfig"`
}

type ExporterScraperConfigStatus struct {
	ActiveExporter corev1.ObjectReference `json:"active,omitempty"`
	ConfigMap      corev1.ObjectReference `json:"configMaps,omitempty"`
	Service        corev1.ObjectReference `json:"services,omitempty"`
}

type ExporterConfig struct {
	Provider ObjectRef `yaml:"provider" json:"provider"`
	Url      string    `yaml:"url" json:"url"`
	// +optional
	UrlParsed             string `yaml:"urlParsed" json:"urlParsed,omitempty"`
	RequireAuthentication bool   `yaml:"requireAuthentication" json:"requireAuthentication"`
	// +kubebuilder:validation:Pattern=`^(\bcert-file\b)|(\bbearer-token\b)$`
	AuthenticationMethod string `yaml:"authenticationMethod" json:"authenticationMethod"`
	// +optional
	BearerToken ObjectRef `yaml:"bearerToken" json:"bearerToken"`
	// +kubebuilder:validation:Pattern=`^(\b[Cc]ost\b)|(\b[Rr]esource\b)$`
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
	Url                      string    `yaml:"url" json:"url,omitempty"`
	ScraperDatabaseConfigRef ObjectRef `yaml:"scraperDatabaseConfigRef" json:"scraperDatabaseConfigRef"`
}

type ScraperConfigStatus struct {
	ActiveScraper corev1.ObjectReference `json:"active,omitempty"`
	ConfigMap     corev1.ObjectReference `json:"configMaps,omitempty"`
}

type ObjectRef struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
}
