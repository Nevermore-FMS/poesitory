import Head from 'next/head'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { GithubIcon } from '../components/Icons/githubIcon'
import { GitlabIcon } from '../components/Icons/gitlabIcon'
import LinkButton from '../styles/ohms-style/react/components/LinkButton'
import styles from "../styles/pages/login.module.scss"

export default function LoginPage() {
    const router = useRouter()

    useEffect(() => {
        const rto = (router.query["redirect-to"] ?? "") as string
        document.cookie = `redirect-to=${encodeURIComponent(rto)}; path=/`
    }, [router.query])

    return (
        <div className="container">
            <Head>
                <title>Login | Poesitory</title>
            </Head>
            {router.query["error"] != null && (
                <div className="warning">{router.query["error"]}</div>
            )}
            <div className={styles.buttonHolder}>
                <LinkButton variant="secondary" href="/api/github/login"><GithubIcon style={{ width: "2.3em", marginRight: "1em" }} /> Login with Github</LinkButton>
                <LinkButton variant="secondary" href="/api/gitlab/login"><GitlabIcon style={{ width: "2.3em", marginRight: "1em" }} /> Login with Gitlab</LinkButton>
            </div>
        </div>
    )
}