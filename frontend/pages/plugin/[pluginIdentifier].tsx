import { GetServerSideProps } from "next";
import { NevermorePluginVersion, QueryPluginVersionArgs } from "../../graphql";
import { addApolloState, initializeApollo } from "../../lib/apolloClient";
import Head from 'next/head'
import { GET_PLUGIN_VERSION } from "../../query";
import PluginVersionDetails from "../../components/PluginVersionDetails";


export default function PluginVersionPage({ pluginVersion }: { pluginVersion: NevermorePluginVersion }) {
    return (
        <div className="container">
            <Head>
                <title>{pluginVersion.plugin?.name} | Poesitory</title>
            </Head>
            <PluginVersionDetails pluginVersion={pluginVersion} />
        </div>
    )
}


export const getServerSideProps: GetServerSideProps = async (context) => {
    const client = initializeApollo()

    const result = await client.query<{ pluginVersion?: NevermorePluginVersion }, QueryPluginVersionArgs>({
        query: GET_PLUGIN_VERSION,
        variables: {
            versionIdentifier: context.params!.pluginIdentifier! as string
        },
        errorPolicy: 'all'
    })

    if (result.data.pluginVersion == null) {
        return {
            notFound: true
        }
    }

    return addApolloState(client, {
        props: {
            pluginVersion: result.data.pluginVersion
        },
    })
}