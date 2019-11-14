let landingPageByIdQuery = (pageId) => {
    return `
        {
          landingPage(id: "` + pageId  + `") {
            title
            subTitle
            bodyText
            questions {
              mcQuestions {
                question
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