package Jobs

type Job interface {
	Publish()
	Consume()
}
