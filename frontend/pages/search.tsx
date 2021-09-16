import { useQuery } from "@apollo/client";
import { useRouter } from "next/dist/client/router";
import PluginHorizontalCard from "../components/PluginHorizontalCard";
import TextField from "../components/TextField";
import { NevermorePluginPage, QuerySearchPluginsArgs } from "../graphql";
import { SEARCH_PLUGINS } from "../query";
import Head from 'next/head'
import styles from "../styles/sass/pages/search.module.scss"
import { GetServerSideProps } from "next";

export default function Search() {
    const router = useRouter()

    const query = router.query["q"]?.toString() || ""

    const setQuery = (q: string) => {
        router.replace({
            query: {
                q
            }
        })
    }

    const { data } = useQuery<{ searchPlugins: NevermorePluginPage }, QuerySearchPluginsArgs>(SEARCH_PLUGINS, {
        variables: {
            search: query
        }
    })

    return (
        <div className="container">
            <Head>
                <title>Search | Poesitory</title>
            </Head>
            <h1>Search for Plugins</h1>
            <TextField containerProps={{ style: { width: "95%" } }} autoFocus={true} placeholder="Plugin Search" value={query} onChange={(e) => setQuery(e.target.value)} />
            <div className={styles.results}>
            {(data?.searchPlugins != null) && (
                data.searchPlugins.plugins?.map(p => (
                    <PluginHorizontalCard key={p.id} plugin={p} />
                ))
            )}
            </div>
        </div>
    )
}