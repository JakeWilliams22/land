let landingPageByIdQuery = (pageId) => {
    return `
        {
          landingPage(id: "` + pageId  + `") {
            id
            title
            subTitle
            bodyText
            googleAnalyticsId
            logoName
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


let addEmailSubscriberQuery = (email, landingPageId) => {
    return `
        mutation {
          addEmailSubscriber(
            email: "` + email + `", 
            landingPageId: "` + landingPageId + `")
        }
    `
}
export {landingPageByIdQuery, addEmailSubscriberQuery};