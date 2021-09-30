import { useMutation, useQuery } from "@apollo/client";
import { GetServerSideProps } from "next";
import Head from 'next/head'
import { useRouter } from "next/router";
import { MutatePluginPayload, MutationCreateUploadTokenArgs, MutationDeleteUploadTokenArgs, NevermorePlugin, QueryPluginArgs, User } from "../../../graphql";
import { addApolloState, initializeApollo } from "../../../lib/apolloClient";
import { CREATE_UPLOAD_TOKEN, DELETE_UPLOAD_TOKEN } from "../../../mutation";
import { GET_ME_USERNAME, GET_PLUGIN } from "../../../query";
import styles from "../../../styles/sass/pages/home.plugin.id.module.scss"


export default function MyPluginPage() {
    const router = useRouter()
    const { data, refetch } = useQuery<{ plugin?: NevermorePlugin }, QueryPluginArgs>(GET_PLUGIN, {
        variables: {
            id: router.query!.id! as string
        }
    })

    const [deleteToken] = useMutation<{ deleteUploadToken: MutatePluginPayload }, MutationDeleteUploadTokenArgs>(DELETE_UPLOAD_TOKEN, {
        refetchQueries: [GET_PLUGIN]
    })

    const [createToken, { data: createTokenData }] = useMutation<{ createUploadToken?: string } | null, MutationCreateUploadTokenArgs>(CREATE_UPLOAD_TOKEN, {
        variables: {
            pluginID: data?.plugin?.id ?? ""
        },
        refetchQueries: [GET_PLUGIN]
    })

    if (data?.plugin == null) {
        return (
            <div />
        )
    }

    const plugin = data.plugin

    return (
        <div className="container">
            <Head>
                <title>{plugin.name} | Poesitory</title>
            </Head>
            <h1>{plugin.name}</h1>
            <p>Type: {plugin.type}</p>
            <p>Latest Version: {plugin.latestFullIdentifier ?? "None"}</p>
            <h3>Upload Tokens</h3>
            {createTokenData?.createUploadToken != null && (
                <div className="warning">Your new upload token is <code>{createTokenData?.createUploadToken}</code>. Be sure to copy it now, as this is the only time you can see it.</div>
            )}
            <div>
                <button className="button-secondary" onClick={() => createToken()}><span className="material-icons">add</span><span> New</span></button>
            </div>
            {plugin.uploadTokens?.map(ut => (
                <div key={ut.id} className={styles.uploadToken}>
                    <div className={styles.uploadTokenDetails}>
                        <p>ID: {ut.id}</p>
                        <p>Created At: {new Date(ut.createdAt * 1000).toLocaleString()}</p>
                    </div>
                    <button className="button-secondary" onClick={() => deleteToken({ variables: { id: ut.id } })}><span className="material-icons">delete</span></button>
                </div>
            ))}
            <small>Upload tokens can be used in the <a href="https://github.com/Nevermore-FMS/poesitory/blob/main/cli/poesitory/README.md" target="_blank" rel="noreferrer">Poesitory CLI</a> to upload plugins. Useful for CI/CD</small>
        </div>
    )
}


export const getServerSideProps: GetServerSideProps = async (context) => {
    const client = initializeApollo(context)

    const meResult = await client.query<{ me?: User }>({
        query: GET_ME_USERNAME
    })

    if (meResult.data.me == null) {
        return {
            redirect: {
                destination: "/login",
                permanent: false
            }
        }
    }

    const pluginResult = await client.query<{ plugin?: NevermorePlugin }, QueryPluginArgs>({
        query: GET_PLUGIN,
        variables: {
            id: context.params!.id! as string
        }
    })

    if (pluginResult.data.plugin == null || pluginResult.data.plugin.owner?.id !== meResult.data.me.id) {
        return {
            notFound: true
        }
    }

    return addApolloState(client, {
        props: {},
    })
}