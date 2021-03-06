import { useMemo } from 'react'
import https from "https"
import { ApolloClient, HttpLink, InMemoryCache, NormalizedCacheObject } from '@apollo/client'
import merge from 'deepmerge'
import isEqual from 'lodash/isEqual'
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next'

export const APOLLO_STATE_PROP_NAME = '__APOLLO_STATE__'

let apolloClient: ApolloClient<NormalizedCacheObject>

function createApolloClient(cookie?: string): ApolloClient<NormalizedCacheObject> {
  const enchancedFetch: WindowOrWorkerGlobalScope['fetch'] = (url, init) => {
    return fetch(url, {
        ...init,
        headers: {
            ...init?.headers,
            "Cookie": cookie ?? ""
        },
    }).then(response => response)
}
  return new ApolloClient({
    ssrMode: typeof window === 'undefined',
    link: new HttpLink({
      uri: process.env.NEXT_PUBLIC_API_URL,
      credentials: "include",
      fetch: enchancedFetch,
      fetchOptions: {
        agent: new https.Agent({ rejectUnauthorized: process.env.NEXT_PUBLIC_DEV_INSECURE === "true" })
      }
    }),
    cache: new InMemoryCache(),
    defaultOptions: {
      mutate: {
        errorPolicy: 'all'
      }
    }
  })
}

export function initializeApollo(context?: GetServerSidePropsContext, initialState: NormalizedCacheObject | null = null) {
  const _apolloClient = apolloClient ?? createApolloClient(context?.req?.headers?.cookie)

  if (initialState) {
    const existingCache = _apolloClient.extract()

    const data = merge(initialState, existingCache, {
      arrayMerge: (destinationArray, sourceArray) => [
        ...sourceArray,
        ...destinationArray.filter((d) =>
          sourceArray.every((s) => !isEqual(d, s))
        ),
      ],
    })

    _apolloClient.cache.restore(data)
  }

  if (typeof window === 'undefined') return _apolloClient
  if (!apolloClient) apolloClient = _apolloClient

  return _apolloClient
}

export function addApolloState(client: ApolloClient<NormalizedCacheObject>, pageProps: GetServerSidePropsResult<{ [key: string]: any }>) {
  if ((pageProps as any)?.props) {
    (pageProps as any).props[APOLLO_STATE_PROP_NAME] = client.cache.extract()
  }

  return pageProps
}

export function useApollo(pageProps: { [key: string]: any }) {
  const state = pageProps[APOLLO_STATE_PROP_NAME]
  const store = useMemo(() => initializeApollo(undefined, state), [state])
  return store
}