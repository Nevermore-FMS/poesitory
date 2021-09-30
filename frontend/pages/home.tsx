import { useQuery } from '@apollo/client'
import { GetServerSideProps } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import PluginHorizontalCard from '../components/PluginHorizontalCard'
import { User, UserOwnedPluginsArgs } from '../graphql'
import { addApolloState, initializeApollo } from '../lib/apolloClient'
import { GET_ME_PLUGINS } from '../query'
import styles from "../styles/sass/pages/home.module.scss"


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
            <Link href="/home/plugin/new"><a className="button-secondary"><span className="material-icons">add</span><span> New</span></a></Link>
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

    const result = await client.query<{ me?: User }, UserOwnedPluginsArgs>({
        query: GET_ME_PLUGINS,
        variables: {
            page: 1
        },
    })

    if (result.data.me == null) {
        return {
            redirect: {
                destination: "/login",
                permanent: false
            }
        }
    }

    return addApolloState(client, {
        props: {},
    })
}