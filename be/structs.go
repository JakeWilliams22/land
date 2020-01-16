package main

type LandingPage struct {
	Title             *string
	SubTitle          *string
	BodyText          *string
	JoinEmailList     *JoinEmailList
	Questions         []Question
	Id                *string //TODO: Force this to be alphanumeric
	GoogleAnalyticsId *string
}

type QuestionType int32

const (
	MULTIPLE_CHOICE QuestionType = iota
	OPEN_ENDED
)

type Question interface {
	QuestionText() string
	QuestionType() QuestionType
}

type MCQuestion struct {
	Question string
	Answers  []string
}

func (mcq MCQuestion) QuestionText() string {
	return mcq.Question
}

func (mcq MCQuestion) QuestionType() QuestionType {
	return MULTIPLE_CHOICE
}

type OpenEndedQuestion struct {
	Question string
}

func (oeq OpenEndedQuestion) QuestionText() string {
	return oeq.Question
}

func (oeq OpenEndedQuestion) QuestionType() QuestionType {
	return OPEN_ENDED
}

type JoinEmailList struct {
	JoinPrompt     string
	JoinButtonText string
}

type GqlRequest struct {
	GqlQuery string
}
