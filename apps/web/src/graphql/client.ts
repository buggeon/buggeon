import { ApolloClient, InMemoryCache, ApolloLink, Observable } from '@apollo/client';
import { HttpLink } from '@apollo/client/link/http';
import { ErrorLink } from '@apollo/client/link/error';
import { ServerError } from '@apollo/client/errors';
import { refreshAccessToken } from '../utils/resreshToken';
import API_URL from '../../config';

const errorLink = new ErrorLink(({ error, operation, forward }) => {
    if (ServerError.is(error) && error.statusCode === 401) {
        return new Observable((observer) => {
            refreshAccessToken()
                .then((newToken) => {
                    operation.setContext(({ headers = {} }) => ({
                        headers: {
                            ...headers,
                            Authorization: `Bearer ${newToken}`,
                        },
                    }));
                    return forward(operation).subscribe(observer);
                })
                .catch((err) => {
                    observer.error(err);
                });
        });
    }
});

const authLink = new ApolloLink((operation, forward) => {
    const token = localStorage.getItem('accessToken');
    operation.setContext(({ headers = {} }) => ({
        headers: {
            ...headers,
            Authorization: token ? `Bearer ${token}` : '',
        },
    }));
    return forward(operation);
});

const httpLink = new HttpLink({
    uri:  `${API_URL}/query`,
});

export const client = new ApolloClient({
    link: ApolloLink.from([authLink, errorLink, httpLink]),
    cache: new InMemoryCache(),
    devtools: { enabled: true },
});