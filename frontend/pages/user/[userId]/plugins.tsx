import { GetServerSideProps } from 'next';
import Head from 'next/head'
import PluginHorizontalCard from "../../../components/PluginHorizontalCard";
import { QueryUserArgs, User } from "../../../graphql";
import { addApolloState, initializeApollo } from '../../../lib/apolloClient';
import { GET_USER_PLUGINS } from '../../../query';
import styles from "../../../styles/pages/search.module.scss"

export default function UserPlugins({ user }: { user: User }) {

    return (
        <div className="container">
            <Head>
                <title>{user.username}&apos;s plugins | Poesitory</title>
            </Head>
            <h1>Plugins by {user.username}</h1>
            <div className={styles.results}>
                {user.ownedPlugins?.plugins?.map(p => (
                    <PluginHorizontalCard key={p.id} plugin={p} href={`/plugin/${p.name}`} />
                ))}
            </div>
        </div>
    )
} 

export const getServerSideProps: GetServerSideProps = async (context) => {
    const client = initializeApollo(context)

    const result = await client.query<{ user?: User }, QueryUserArgs>({
        query: GET_USER_PLUGINS,
        variables: {
            id: context.params!.userId! as string
        }
    })

    if (result.data.user == null) {
        return {
            notFound: true
        }
    }

    return addApolloState(client, {
        props: {
            user: result.data.user
        }
    })
    
}