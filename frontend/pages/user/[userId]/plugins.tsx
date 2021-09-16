import { GetServerSideProps } from 'next';
import Head from 'next/head'
import PluginHorizontalCard from "../../../components/PluginHorizontalCard";
import { NevermorePlugin, User } from "../../../graphql";
import styles from "../../../styles/sass/pages/search.module.scss"

export default function UserPlugins({ user, plugins }: { user: User, plugins: NevermorePlugin[] }) {

    return (
        <div className="container">
            <Head>
                <title>{user.username}&apos;s plugins | Poesitory</title>
            </Head>
            <h1>Plugins by {user.username}</h1>
            <div className={styles.results}>
                {plugins?.map(p => (
                    <PluginHorizontalCard key={p.id} plugin={p} />
                ))}
            </div>
        </div>
    )
} 

export const getServerSideProps: GetServerSideProps = async (context) => {
    //const client = initializeApollo()
    //TODO This page

    return {
        notFound: true
    }
}