export function constructLoginRedirect(page?: string): any {
    const p = new URLSearchParams()
    if (page != null) p.set("redirect-to", page)
    const params = p.toString()
    return {
        redirect: {
            permanent: false,
            destination: `/login${params.length > 0 ? "?" : ""}${params}`,
        }
    }
}
