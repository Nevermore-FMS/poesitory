import Head from 'next/head'
import styles from "../styles/sass/pages/404.module.scss"

export default function Page404() {
    return (
        <div className="container">
            <Head>
                <title>404 Not Found | Poesitory</title>
            </Head>
            <div className={styles.notFound}>
                <h1>Oh No!</h1>
                <img alt="Robot yeeting onto a platform" src="/gif/yeet.gif" />
                <h3>It looks like you yeeted over to the wrong page :(</h3>
                <span>Use the navigation tabs at the top to browse the site</span>
            </div>

        </div>
    )
}