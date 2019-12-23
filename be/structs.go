package main

type LandingPage struct {
        Title         *string
        SubTitle      *string
        BodyText      *string
        JoinEmailList *JoinEmailList
        Questions     *Questions
        Id            *string
}

type Questions struct {
        QuestionsPrompt string
        //TODO Figure out how to make this an interface.
        McQuestions        []MCQuestion
        OpenEndedQuestions []OpenEndedQuestion
}

func (q *Questions) isEmpty() bool {
  return len(q.McQuestions) + len(q.OpenEndedQuestions) == 0
}

type MCQuestion struct {
        Question string
        Answers  []string
}

type OpenEndedQuestion struct {
        Question string
}

type JoinEmailList struct {
        JoinPrompt     string
        JoinButtonText string
}

type GqlRequest struct {
        GqlQuery string
}