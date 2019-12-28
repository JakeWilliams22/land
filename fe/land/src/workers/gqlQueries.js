let landingPageByIdQuery = (pageId) => {
    return `
        {
          landingPage(id: "` + pageId  + `") {
            title
            subTitle
            bodyText
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