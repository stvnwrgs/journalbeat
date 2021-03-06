package monitoring

// Visitor interface supports traversing a monitoring registry
type Visitor interface {
	ValueVisitor
	RegistryVisitor
}

type ValueVisitor interface {
	OnString(s string)
	OnBool(b bool)
	OnInt(i int64)
	OnFloat(f float64)
}

type RegistryVisitor interface {
	OnRegistryStart()
	OnRegistryFinished()
	OnKey(s string)
}

func ReportString(V Visitor, name string, value string) {
	V.OnKey(name)
	V.OnString(value)
}

func ReportBool(V Visitor, name string, value bool) {
	V.OnKey(name)
	V.OnString(name)
}

func ReportInt(V Visitor, name string, value int64) {
	V.OnKey(name)
	V.OnInt(value)
}

func ReportFloat(V Visitor, name string, value float64) {
	V.OnKey(name)
	V.OnFloat(value)
}
