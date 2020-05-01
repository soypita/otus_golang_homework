package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipelineIn := make(In)
	pipelineOut := make(Out)

	if len(stages) == 0 {
		emptyCh := make(chan I)
		close(emptyCh)
		return emptyCh
	}

	for i := 0; i < len(stages); i++ {
		if i == 0 {
			pipelineOut = stages[i](in)
		} else {
			pipelineOut = stages[i](pipelineIn)
		}
		pipelineIn = pipelineOut
	}
	proxyOutCh := make(Bi)
	go proxyOutListener(pipelineOut, proxyOutCh, done)
	return proxyOutCh
}

func proxyOutListener(pipelineOut Out, proxyOutCh Bi, done In) {
	defer close(proxyOutCh)
	for {
		select {
		case value, ok := <-pipelineOut:
			if !ok {
				return
			}
			proxyOutCh <- value
		case <-done:
			return
		}
	}
}
