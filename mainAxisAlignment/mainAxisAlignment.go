package mainAxisAlignment

// MainAxisAlignment how the children should be placed along the main axis in a flex layout
type MainAxisAlignment string

const (
	// Start place the children as close to the start of the main axis as possible
	Start MainAxisAlignment = "start"
	// End place the children as close to the end of the main axis as possible
	End MainAxisAlignment = "end"
	// Center place the children as close to the middle of the main axis as possible
	Center MainAxisAlignment = "center"
	// SpaceBetween place the free space evenly between the children
	SpaceBetween MainAxisAlignment = "spaceBetween"
	// SpaceAround place the free space evenly between the children as well as half of that
	// space before and after the first and last child.
	SpaceAround MainAxisAlignment = "spaceAround"
	// SpaceEvenly place the free space evenly between the children as well as before and
	// after the first and last child.
	SpaceEvenly MainAxisAlignment = "spaceEvenly"
)
