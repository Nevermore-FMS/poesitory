import { useApolloClient } from '@apollo/client'
import { GetServerSideProps } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { useState } from 'react'
import Select from '../../../components/Select'
import TextField from '../../../components/TextField'
import { MutatePluginPayload, MutationCreatePluginArgs, NevermorePluginPage, NevermorePluginType, User } from '../../../graphql'
import { addApolloState, initializeApollo } from '../../../lib/apolloClient'
import { CREATE_PLUGIN } from '../../../mutation'
import { GET_ME_USERNAME } from '../../../query'
import styles from "../../../styles/sass/pages/home.plugin.new.module.scss"

export default function NewPluginPage() {
    const client = useApolloClient()
    const router = useRouter()
    const [saving, setSaving] = useState(false)
    const [name, setName] = useState("")
    const [type, setType] = useState<string>(NevermorePluginType.Generic)
    const [error, setError] = useState<string | null>(null)

    const createPlugin = async () => {
        setSaving(true)
        try {
            const result = await client.mutate<{ createPlugin: MutatePluginPayload }, MutationCreatePluginArgs>({
                mutation: CREATE_PLUGIN,
                variables: {
                    name,
                    type: type as NevermorePluginType
                }
            })
            if ((result.errors?.length ?? 0) > 0) {
                setError(result.errors!!.map(e => e.message).join(' '))
                return
            }
            router.push(`/home/plugin/${result.data?.createPlugin.plugin?.id}`)
        } finally {
            setSaving(false)
        }
    }

    return (
        <div className="container">
            <Head>
                <title>New Plugin | Poesitory</title>
            </Head>
            <h1>New Plugin</h1>
            {error != null && (<div className="error">{error}</div>)}
            <div className={styles.form}>
                <TextField placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
                <Select placeholder="Type" value={type} onChange={(e) => setType(e.target.value)}>
                    <option value="GENERIC">Generic</option>
                    <option value="GAME">Game</option>
                    <option value="NETWORK_CONFIGURATOR">Network Configurator</option>
                </Select>
                <button disabled={saving} className={saving ? "button-disabled" : "button-secondary"} onClick={createPlugin}>Create Plugin</button>
            </div>
        </div>
    )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
    const client = initializeApollo(context)

    const result = await client.query<{ me?: User }>({
        query: GET_ME_USERNAME
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