let landingPageByIdQuery = (pageId) => {
    return `
        {
          landingPage(id: "` + pageId  + `") {
            id
            title
            subTitle
            bodyText
            googleAnalyticsId
            questions {
                questionPrompt
                questionType
                ... on MCQuestion {
                    answers
                }
            }
            joinEmailList {
                joinPrompt
                joinButtonText
            }
          }
        }`;
};

export {landingPageByIdQuery};