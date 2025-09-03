package bitnamialt

type HelmChart struct {
	// URL is the OCI repository URL for the Helm Chart. If empty, the Registry
	// and ChartName fields should be non-empty.
	URL string `json:"url,omitempty"`
	// Registry is the Helm Chart registry URL for the Helm Chart. If empty,
	// URL should be non-empty and contain the OCI repository URL for the Helm
	// Chart.
	Registry string `json:"registry,omitempty"`
	// Name is the name for the Helm Chart. If empty, URL should be non-empty
	// and contain the OCI repository URL for the Helm Chart. If not empty,
	// Registry should also be not empty.
	Name string `json:"name,omitempty"`
}

type Artifacts struct {
	// Images is a collection of Docker/OCI Image URLs
	Images []string `json:"images"`
	// Charts is a collection of Helm Chart locations
	Charts []HelmChart `json:"charts,omitempty"`
}

type ProjectEntry struct {
	// Project is the common name for the Application, Kubernetes Add-on, or
	// OSS project.
	Project string `json:"project"`
	// Bitnami contains artifact information for Bitnami-provided
	// images/charts.
	Bitnami Artifacts `json:"bitnami"`
	// Alternatives contains artifact information for non-Bitnami-provided
	// alternative images/charts.
	Alternatives Artifacts `json:"alternatives"`
}

type Entries []ProjectEntry
