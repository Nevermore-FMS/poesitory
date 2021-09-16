import { ApolloProvider } from '@apollo/client'
import type { AppProps } from 'next/app'
import Header from '../components/Header'
import { useApollo } from '../lib/apolloClient'
import "../styles/sass/main.scss"
import "../styles/sass/pages/nprogress.scss"
import "../styles/normalize.css"
import "highlight.js/styles/default.css"
import Footer from '../components/Footer'
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import Router from 'next/router';

Router.events.on('routeChangeStart', () => NProgress.start());
Router.events.on('routeChangeComplete', () => NProgress.done());
Router.events.on('routeChangeError', () => NProgress.done())

export default function App({ Component, pageProps }: AppProps) {
  const apolloClient = useApollo(pageProps)

  return (
    <ApolloProvider client={apolloClient}>
      <Header />
      <Component {...pageProps} />
      <Footer />
    </ApolloProvider>
  )
}