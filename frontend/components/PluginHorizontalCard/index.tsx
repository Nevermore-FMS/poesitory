import { NevermorePlugin } from "../../graphql";
import Link from 'next/link'

import styles from "./index.module.scss"

export default function PluginHorizontalCard({ plugin }: { plugin: NevermorePlugin }) {
    return (
        <Link href={`/plugin/${plugin.name}`}>
            <a className={["card", styles.pluginCard].join(' ')}>
                <div>
                    <p>{plugin.name}</p>
                    <div>Type: <b>{plugin.type}</b></div>
                    {plugin.owner && (
                        <div>Created by: <b>{plugin.owner.username}</b></div>
                    )}
                </div>
                <div className={styles.pluginIdentifier}>
                    <p>{plugin.latestFullIdentifier}</p>
                </div>
            </a>
        </Link>
    )
}