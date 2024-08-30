package helpers

type TestCase struct {
	TestName    string
	Request     Request
	RequestBody interface{}
	Expected    ExpectedResponse
}

type PathParam struct {
	Name  string
	Value string
}

type Request struct {
	Method    string
	URL       string
	PathParam *PathParam
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
	BodyParts  []string
}
