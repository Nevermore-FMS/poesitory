import styles from "./index.module.scss"

export default function Footer() {
    return (
        <footer className={styles.footer}>
            <span className={styles.footerText}><i>Poesitory and the Nevermore FMS is a project by the <a href="https://edgarallanohms.com/" target="_blank" rel="noreferrer">Edgar Allan Ohms</a>.</i></span>
        </footer>

    )
}