package parser

// Decode configuration
func Decode(labels map[string]string, element interface{}, rootName string, filters ...string) error {
	node, err := DecodeToNode(labels, rootName, filters...)
	if err != nil {
		return err
	}

	err = AddMetadata(element, node)
	if err != nil {
		return err
	}

	err = Fill(element, node)
	if err != nil {
		return err
	}

	return nil
}

// Encode converts an element to labels.
// element -> node (value) -> label (node)
func Encode(element interface{}, rootName string) (map[string]string, error) {
	node, err := EncodeToNode(element, rootName, true)
	if err != nil {
		return nil, err
	}

	return EncodeNode(node), nil
}
