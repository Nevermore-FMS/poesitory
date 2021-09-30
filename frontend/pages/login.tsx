import Head from 'next/head'
import { useRouter } from 'next/router'
import { GithubIcon } from '../components/Icons/githubIcon'
import { GitlabIcon } from '../components/Icons/gitlabIcon'
import styles from "../styles/sass/pages/login.module.scss"

export default function LoginPage() {
    const router = useRouter()

    return (
        <div className="container">
            <Head>
                <title>Login | Poesitory</title>
            </Head>
            {router.query["error"] != null && (
                <div className="warning">{router.query["error"]}</div>
            )}
            <div className={styles.buttonHolder}>
                <a className="button-secondary" href="/api/github/login"><GithubIcon style={{ width: "2.3em", marginRight: "1em" }} /> Login with Github</a>
                <a className="button-secondary" href="/api/gitlab/login"><GitlabIcon style={{ width: "2.3em", marginRight: "1em" }} /> Login with Gitlab</a>
            </div>
        </div>
    )
}