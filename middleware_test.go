package slack

import "testing"

func TestRecovery(t *testing.T) {
	s:=SlackClient{}
	Recovery(&s)

}
