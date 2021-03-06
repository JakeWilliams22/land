class GqlClient {
    static GQL_URL = "http://localhost:8080/graphql";
    static query = (query) => {
        return fetch(this.GQL_URL,
            {
                method: 'POST',
                body: JSON.stringify(query)
            }).then(response => {
                console.log(response);
                return response;
            });
    };
}

export default GqlClient;