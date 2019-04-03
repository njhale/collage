package image

// Image is an image
type Image interface {
	Layers()
	Pull()
	Push()
}

// Collage represents a
type Collage interface {
	Image

	// Add adds an image's layers to the Collage
	Add(addee Image)

	// Remove removes an image's layers from the Collage
	Remove(removee Image)
}

// Pattern describes a set of images to
type Pattern interface {
}

// Snipper snips layers from an image
type Snipper interface {
	Snip(image Image, pattern Pattern) (target, scraps Image)
}
