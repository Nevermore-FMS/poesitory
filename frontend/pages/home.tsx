import { useQuery } from '@apollo/client'
import { GetServerSideProps } from 'next'
import nookies from 'nookies'
import Head from 'next/head'
import Link from 'next/link'
import PluginHorizontalCard from '../components/PluginHorizontalCard'
import { User, UserOwnedPluginsArgs } from '../graphql'
import { addApolloState, initializeApollo } from '../lib/apolloClient'
import { GET_ME_PLUGINS } from '../query'
import LinkButton from '../styles/ohms-style/react/components/LinkButton'
import styles from "../styles/pages/home.module.scss"
import { constructLoginRedirect } from '../lib/redirect'


export default function HomePage() {
    const { data } = useQuery<{ me?: User }, UserOwnedPluginsArgs>(GET_ME_PLUGINS, {
        variables: {
            page: 1
        }
    })

    return (
        <div className="container">
            <Head>
                <title>My Plugins | Poesitory</title>
            </Head>
            <h1>My Plugins</h1>
            <Link passHref href="/home/plugin/new"><LinkButton variant="secondary"><span className="material-icons">add</span><span> New</span></LinkButton></Link>
            <div className={styles.main}>
                {(data?.me?.ownedPlugins?.plugins != null) && (
                    data.me.ownedPlugins.plugins.map(p => (
                        <PluginHorizontalCard key={p.id} plugin={p} href={`/home/plugin/${p.id}`} />
                    ))
                )}
            </div>
        </div>
    )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
    const client = initializeApollo(context)

    const cookies = nookies.get(context)
    if (cookies["redirect-to"] != null) {
        const rto = decodeURIComponent(cookies["redirect-to"] as string)
        nookies.destroy(context, "redirect-to")
        if (rto.length > 0 && !isUrlAbsolute(rto)) {
            return {
                redirect: {
                    destination: rto,
                    permanent: false
                }
            }
        }
    }

    const result = await client.query<{ me?: User }, UserOwnedPluginsArgs>({
        query: GET_ME_PLUGINS,
        variables: {
            page: 1
        },
    })

    if (result.data.me == null) {
        return constructLoginRedirect("/home")
    }

    return addApolloState(client, {
        props: {},
    })
}

function isUrlAbsolute(url: string) { 
    return (url.indexOf('://') > 0 || url.indexOf('//') === 0);
}