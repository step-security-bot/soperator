package v1

// ImageSpec defines the desired state of the image used for container
type ImageSpec struct {
	// Repository defines the path to the image
	Repository string `json:"repository"`

	// Tag defines the tag of the image
	Tag string `json:"tag"`
}
