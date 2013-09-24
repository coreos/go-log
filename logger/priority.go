package logger

type Priority int

const (
	PriEmerg Priority = iota
	PriAlert
	PriCrit
	PriErr
	PriWarning
	PriNotice
	PriInfo
	PriDebug
)

func (priority Priority) String() string {
	switch priority {
	case PriEmerg:
		return "EMERGENCY"
	case PriAlert:
		return "ALERT"
	case PriCrit:
		return "CRITICAL"
	case PriErr:
		return "ERROR"
	case PriWarning:
		return "WARNING"
	case PriNotice:
		return "NOTICE"
	case PriInfo:
		return "INFO"
	case PriDebug:
		return "DEBUG"

	default:
		return "UNKNOWN"
	}
}
