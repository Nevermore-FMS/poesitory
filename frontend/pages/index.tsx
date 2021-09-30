import { useRouter } from 'next/dist/client/router'
import Head from 'next/head'
import TextField from "../styles/ohms-style/react/components/TextField"
import styles from "../styles/pages/index.module.scss"

export default function Home() {
    const router = useRouter()

    return (
        <div>
            <Head>
                <title>Poesitory - Plugin Repository for Nevermore FMS</title>
            </Head>
            <div className={styles.homeSpread}>
                <picture>
                    <source srcSet="/img/eao_bird_circle.webp" type="image/webp" />
                    <source srcSet="/img/eao_bird_circle.png" type="image/png" />
                    <img src="/img/eao_bird_circle.png" alt="" className={styles.homeLogo} />
                </picture>
                <h1>Poesitory</h1>
                <h2>Plugin Repository for Nevermore FMS</h2>
                <TextField light placeholder="Search for a plugin" onChange={(e) => router.push({
                    pathname: "/search",
                    query: {
                        q: e.target.value
                    }
                })} />
            </div>
        </div>
    )
}
