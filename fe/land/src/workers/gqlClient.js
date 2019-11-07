class GqlClient {
    GQL_URL = "http://localhost:8080/graphql";

    query = (query) => {
        return fetch(this.GQL_URL,
            {
                method: 'post',
                body: JSON.stringify(query)
            })
    }
}

export default GqlClient;