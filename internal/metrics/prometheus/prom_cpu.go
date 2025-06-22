package prometheus

type PromCPU struct{}

func (promCPU PromCPU) Load() string {
	return "Some Load"
}

func (promCPU PromCPU) ResponseCount(code string, time string) string {
	return "some count"
}
