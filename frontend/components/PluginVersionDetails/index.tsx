import { NevermorePluginVersion } from "../../graphql";
import Link from 'next/link'
import styles from "./index.module.scss"
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import rehypeHighlight from 'rehype-highlight'

export default function PluginVersionDetails({ pluginVersion }: { pluginVersion: NevermorePluginVersion }) {
    return (
        <div>
            <h1>{pluginVersion.plugin?.name}</h1>
            <p>{pluginVersion.fullIdentifier}</p>
            <div className={styles.detailsHolder}>
                <div className={styles.detailsMain}>
                    <div className="card">
                        <p className={styles.readmeHeader}>README</p>
                        <ReactMarkdown remarkPlugins={[remarkGfm]} rehypePlugins={[rehypeHighlight]}>{pluginVersion.readme || ""}</ReactMarkdown>
                    </div>
                </div>
                <div className={styles.detailsSidebar}>
                    <div className="card">
                        <b>Created by</b>
                        <div>
                            <Link href={`/user/${pluginVersion.plugin?.owner?.id}/plugins`}>
                                <a>{pluginVersion.plugin?.owner?.username}</a>
                            </Link>
                        </div>
                    </div>
                    <div className="card">
                        <b>Type</b>
                        <div>{pluginVersion.plugin?.type}</div>
                    </div>
                    <div className="card">
                        <b>Versions</b>
                        {pluginVersion.plugin?.channels?.sort((a, b) => a.name === "STABLE" ? -1 : 0)?.map((c) => (
                            <div key={c.name}>
                                <i>{c.name}</i>
                                <div className={styles.versionLinks}>
                                    {c.versions?.map((v) => (
                                        <Link key={v.id} href={`/plugin/${v.fullIdentifier}`}>
                                            <a >{v.fullIdentifier}</a>
                                        </Link>
                                    ))}
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    )
}