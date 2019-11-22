package message

type Replyer interface {
	Reply(corpID, userId string) string
}
