import { GetServerSideProps } from "next"
import { User } from "../graphql"
import Head from 'next/head'
import { addApolloState, initializeApollo } from "../lib/apolloClient"
import { constructLoginRedirect } from "../lib/redirect"
import { GET_ME_USERNAME } from "../query"
import styles from "../styles/pages/webauth.module.scss"
import { useEffect, useState } from "react"
import Button from "../styles/ohms-style/react/components/Button"

export default function WebAuth({ username, token }: { username: string, token: string }) {
    const [accepted, setAccepted] = useState<boolean | null>(null)
    const [lastAuthorized, setLastAuthorized] = useState("Never")

    useEffect(() => {
        const id = setInterval(() => {
            if (accepted === true) {
                fetch('http://localhost:25622/webauth', {
                    method: 'POST',
                    body: token
                }).then(r => {
                    if (r.status === 200) {
                        setLastAuthorized(new Date().toLocaleTimeString())
                    } else {
                        console.error(r)
                    }
                }).catch(() => {})
            }
        }, 1000)
        return () => {
            clearInterval(id)
        }
    })

    return (
        <div className="container">
            <Head>
                <title>WebAuth | Poesitory</title>
            </Head>
            <div className={styles.holder}>
                {accepted == null && (
                    <>
                        <h1>Authorize Poesitory CLI as {username}?</h1>
                        <div className={styles.buttons}>
                            <Button variant="primary" onClick={() => setAccepted(false)}>No</Button>
                            <Button variant="secondary" onClick={() => setAccepted(true)}>Yes</Button>
                        </div>
                    </>
                )}

                {accepted == false && (
                    <>
                        <h1>Authorization rejected</h1>
                        <p>You may now close this page or tab.</p>
                    </>
                )}

                {accepted == true && (
                    <>
                        <h1>Authorization accepted</h1>
                        <p>Keep this page open to auto-accept future authorization requests.</p>
                        <p>Last Authorized: {lastAuthorized}</p>
                    </>
                )}
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
        return constructLoginRedirect("/webauth")
    }

    return addApolloState(client, {
        props: {
            username: result.data.me.username,
            token: context.req.cookies.token
        },
    })
}