let landingPageByIdQuery = (pageId) => {
    return `
        {
          landingPage(id: "` + pageId  + `") {
            title
            subTitle
            bodyText
            questions {
                questionPrompt
            }
            joinEmailList {
                joinPrompt
                joinButtonText
            }
          }
        }`;
};

export {landingPageByIdQuery};