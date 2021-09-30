import { useRouter } from "next/dist/client/router"
import React, { useRef, useState } from "react"
import Head from 'next/head'
import Link from 'next/link'
import styles from "./index.module.scss"
import { useQuery } from "@apollo/client"
import { GET_ME_USERNAME } from "../../query"
import { User } from "../../graphql"

interface HeaderLinks {
    name: string
    page: string
}

export default function Header() {
    const [expanded, setExpanded] = useState(false)
    const growableRef = useRef<HTMLDivElement>(null)
    const router = useRouter()

    const { data: meData } = useQuery<{ me: User }>(GET_ME_USERNAME)

    const links: HeaderLinks[] = [
        {
            name: "Home",
            page: "/"
        },
        {
            name: "Search",
            page: "/search"
        },
    ]

    const externalLinks: HeaderLinks[] = [
        {
            name: "CLI",
            page: "https://github.com/Nevermore-FMS/poesitory/blob/main/cli/poesitory/README.md"
        },
    ]

    const renderLinks = [...(
        links.map((l) => (
            <Link href={l.page} key={l.name}>
                <a className={[styles.headerLink, router.pathname === l.page ? styles.active : null].join(" ")}>
                    {l.name}
                </a>
            </Link>
        ))
    ),
    ...(
        externalLinks.map((l) => (
            <a href={l.page} key={l.name} className={styles.headerLink} target="_blank" rel="noreferrer">
                {l.name} <span className="material-icons" style={{ fontSize: "0.7em" }}>open_in_new</span>
            </a>
        ))
    )]

    return (
        <header className={styles.header}>
            <Head>
                <link rel="icon" type="image/png" href="/img/eao_bird_circle.png" />
                <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />

            </Head>
            <div className={styles.desktopHeader}>
                <Link href="/">
                    <a className={styles.desktopHomeLink}>
                        <picture>
                            <source srcSet="/img/eao_bird_circle.webp" type="image/webp" />
                            <source srcSet="/img/eao_bird_circle.png" type="image/png" />
                            <img className={styles.headerLogo} src="/img/eao_bird_circle.png" alt="Edgar Allan Ohms Logo" />
                        </picture>
                        <span className={styles.headerText}>Poesitory</span>
                    </a>
                </Link>
                <div className={styles.desktopHeaderLinks}>
                    {renderLinks}
                    {meData?.me != null ? (
                        <>
                            <Link href="/home">
                                <a className={[styles.headerLink].join(" ")} style={{ marginLeft: "auto" }}>
                                    {meData.me.username}
                                </a>
                            </Link>
                            <span>|</span>
                            <a href="/api/logout" className={[styles.headerLink].join(" ")}>
                                Logout
                            </a>
                        </>
                    ) : (
                        <Link href="/login">
                            <a className={[styles.headerLink].join(" ")} style={{ marginLeft: "auto" }}>
                                Login
                            </a>
                        </Link>
                    )}

                </div>
            </div >
            <div className={styles.mobileHeader}>
                <div className={styles.mobileHeaderAction} onClick={() => setExpanded(!expanded)}>
                    <picture>
                        <source srcSet="/img/eao_bird_circle.webp" type="image/webp" />
                        <source srcSet="/img/eao_bird_circle.png" type="image/png" />
                        <img className={styles.headerLogo} src="/img/eao_bird_circle.png" alt="Edgar Allan Ohms Logo" />
                    </picture>
                    <span className={styles.headerText}>Poesitory</span>
                    <span className="material-icons">expand_more</span>
                </div>
                <div className={styles.mobileGrowable} style={{ maxHeight: expanded ? growableRef?.current?.clientHeight + "px" : "0px" }}>
                    <div className={styles.mobileHeaderLinks} ref={growableRef}>
                        {renderLinks}
                        {meData?.me != null ? (
                            <>
                                <Link href="/home">
                                    <a className={[styles.headerLink].join(" ")} style={{ marginTop: "30px" }}>
                                        {meData.me.username}
                                    </a>
                                </Link>
                                <a href="/api/logout" className={[styles.headerLink].join(" ")}>
                                    Logout
                                </a>
                            </>
                        ) : (
                            <Link href="/login">
                                <a className={[styles.headerLink].join(" ")} style={{ marginTop: "30px" }}>
                                    Login
                                </a>
                            </Link>
                        )}
                    </div>
                </div>
            </div>
        </header >
    )
}