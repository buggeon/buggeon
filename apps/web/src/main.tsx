import { createRoot } from 'react-dom/client'
import './index.scss'
import App from './App.tsx'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'
import { persistor, store } from './store/slices/index.ts'
import { PersistGate } from 'redux-persist/integration/react'
import { ApolloProvider } from '@apollo/client/react'
import { client } from './graphql/client.ts'

createRoot(document.getElementById('root')!).render(
    <BrowserRouter>
        <Provider store={store}>
            <PersistGate persistor={persistor}>
                <ApolloProvider client={client}>
                    <App/>
                </ApolloProvider>
            </PersistGate>
        </Provider>
    </BrowserRouter>
)
