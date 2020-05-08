package capv

// Events returns the channel of progress messages
func (m MgmtCluster) Events() chan string {
	return m.EventStream
}
