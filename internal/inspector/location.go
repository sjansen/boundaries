package inspector

import "strconv"

const (
	GoRoot Location = iota + 1
	GoPath
	Project
	Vendor
)

type Location int

func (l Location) String() string {
	switch l {
	case GoRoot:
		return "GoRoot"
	case GoPath:
		return "GoPath"
	case Project:
		return "Project"
	case Vendor:
		return "Vendor"
	}
	return "(invalid)"
}

func (l *Location) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch s {
	case "GoRoot":
		*l = GoRoot
	case "GoPath":
		*l = GoPath
	case "Project":
		*l = Project
	case "Vendor":
		*l = Vendor
	}

	return nil
}
